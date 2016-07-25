package safari

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/registry"
	"github.com/buckhx/safari-zone/srv"
	"github.com/buckhx/safari-zone/srv/auth"
	"github.com/gengo/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	registryAddr = "localhost:50052" //TODO make this part of the opts
	pdxAddr      = "localhost:50051" //TODO make this part of the opts
)

// Server API for Safari service
type Service struct {
	*warden
	reg  *registry.Client
	opts srv.Opts
}

func NewService(addr string) (srv.Service, error) {
	reg, err := registry.Dial(registryAddr)
	if err != nil {
		return nil, err
	}
	return &Service{
		warden: newGame(),
		reg:    reg,
		opts: srv.Opts{
			Addr: addr,
			Auth: auth.Opts{
				CertURI: "http://localhost:8080/registry/v0/cert",
			},
		},
	}, nil
}

func (s Service) Name() string {
	return "safari"
}

func (s Service) Version() string {
	return "v0"
}

// Enter might add a pokemon to the event
//
// A pokemon will be added to the event and timestmap set if one is encountered
func (sf *Service) Enter(ctx context.Context, req *pbf.Ticket) (*pbf.Ticket, error) {
	if req.Zone.Region != pbf.KANTO {
		return nil, grpc.Errorf(codes.FailedPrecondition, "That zone is under construction")
	}
	claims, _ := auth.ClaimsFromContext(ctx)
	if req.Trainer.Uid != claims.Subject {
		return nil, grpc.Errorf(codes.PermissionDenied, "Claims not scoped to requested trainer")
	}
	if tkt, ok := sf.tix.Get(claims.Subject).(*pbf.Ticket); ok { //TODO move this to the warden
		return tkt, nil
	}
	exp := &pbf.Ticket_Expiry{Time: time.Now().Add(10 * time.Minute).Unix(), Encounters: 5}
	tkt, err := sf.issueTicket(req.Trainer, req.Zone, exp)
	if err != nil {
		err = grpc.Errorf(codes.AlreadyExists, err.Error())
	}
	return tkt, err
}

// Encounter will attempt to catch the pokemon
//
// If caught, this pokemon will be deposited into the Trainer's PC
func (sf *Service) Encounter(stream pbf.Safari_EncounterServer) error {
	pok, err := sf.spawn(stream.Context())
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("A wild %s was encountered!", pok.Name)
	if err := stream.Send(&pbf.BattleMessage{Msg: msg}); err != nil {
		return err
	}
	spd := float64(pok.Speed)
	cth := float64(pok.CatchRate)
	for {
		//fmt.Println("SPEED", spd/255.0)
		//fmt.Println("CATCH", cth/255.0)
		in, err := stream.Recv()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		}
		if rand.Float64() <= spd/255.0 {
			msg := fmt.Sprintf("%s fled!", pok.Name)
			return stream.Send(&pbf.BattleMessage{Msg: msg, Status: pbf.DONE})
		}
		var msg string
		switch act := in.Move.(type) {
		case *pbf.Action_Attack:
			switch act.Attack {
			case "safari-ball":
				if rand.Float64() <= cth/255.0 {
					msg = fmt.Sprintf("%s was caught!", pok.Name)
					claims, ok := auth.ClaimsFromContext(stream.Context())
					if !ok {
						fmt.Println(stream.Context())
						return fmt.Errorf("No auth claims")
					}
					trn, err := sf.reg.GetTrainer(stream.Context(), &pbf.Trainer{Uid: claims.Subject})
					if err != nil {
						return err
					}
					trn.Pc.Pokemon = append(trn.Pc.Pokemon, pok)
					if _, err = sf.reg.UpdateTrainer(stream.Context(), trn); err != nil {
						return err
					}
					return stream.Send(&pbf.BattleMessage{Msg: msg, Status: pbf.DONE})
				} else {
					msg = fmt.Sprintf("%s broke free!", pok.Name)
				}
			case "throw-rock":
				spd -= 10
				cth -= 5
				msg = fmt.Sprintf("%s is angry!", pok.Name)
			case "offer-bait":
				cth += 10
				spd += 5
				msg = fmt.Sprintf("%s is eating...", pok.Name)
			default:
				msg = fmt.Sprintf("%s is watching carefully", pok.Name)
			}
		case *pbf.Action_Item:
			msg = "There's a time and place for everything!"
		case *pbf.Action_Switch:
			msg = "Now's not the time for that!"
		case *pbf.Action_Run:
			msg := "Got away safely!"
			return stream.Send(&pbf.BattleMessage{Msg: msg, Status: pbf.DONE})
		default:
			msg = "waiting..."
		}
		if err = stream.Send(&pbf.BattleMessage{Msg: msg}); err != nil {
			return err
		}
	}
}

func (s *Service) Listen() error {
	tcp, err := net.Listen("tcp", s.opts.Addr)
	if err != nil {
		return err
	}
	rpc, err := srv.ConfigureGRPC(s.opts)
	if err != nil {
		return err
	}
	pbf.RegisterSafariServer(rpc, s)
	log.Printf("Service %T listening at %s", s, s.opts.Addr)
	return rpc.Serve(tcp)
}

func (s *Service) Mux() (http.Handler, error) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pbf.RegisterSafariHandlerFromEndpoint(s.ctx, mux, s.opts.Addr, opts)
	if err != nil {
		mux = nil
	}
	return http.Handler(mux), err
}
