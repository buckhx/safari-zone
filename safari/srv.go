package safari

import (
	"io"
	"log"
	"net"
	"net/http"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/registry"
	"github.com/buckhx/safari-zone/srv"
	"github.com/buckhx/safari-zone/srv/auth"
	"github.com/gengo/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const registryAddr = "localhost:50052" //TODO make this part of the opts

// Server API for Safari service
type Service struct {
	*game
	*registry.SrvClient
	opts srv.Opts
}

func NewService(addr string) (srv.Service, error) {
	return &Service{
		game:      newGame(),
		SrvClient: registry.NewSrvClient(registryAddr),
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
	claims, _ := auth.ClaimsFromContext(ctx)
	if req.Trainer.Uid != claims.Subject {
		return nil, grpc.Errorf(codes.PermissionDenied, "Claims not scoped to requested trainer")
	}
	tkt, err := sf.issueTicket(req.Trainer, req.Zone, 10)
	if err != nil {
		err = grpc.Errorf(codes.AlreadyExists, err.Error())
	}
	return tkt, err
}

// Encounter will attempt to catch the pokemon
//
// If caught, this pokemon will be deposited into the Trainer's PC
func (sf *Service) Encounter(stream pbf.Safari_EncounterServer) error {
	for {
		in, err := stream.Recv()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		}
		var msg string
		switch act := in.Move.(type) {
		case *pbf.Action_Attack:
			msg = "attacked with " + act.String()
		case *pbf.Action_Item:
			msg = "There's a time and place for everything!"
		case *pbf.Action_Switch:
			msg = "Now's not the time for that!"
		case *pbf.Action_Run:
			msg := "Got away safely!"
			return stream.Send(&pbf.BattleMessage{Msg: msg})
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
	err = s.BindRegistry()
	if err != nil {
		return nil
	}
	log.Printf("Service %T listening at %s", s, s.opts.Addr)
	return rpc.Serve(tcp)
}

func (s *Service) Mux() (http.Handler, error) {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pbf.RegisterSafariHandlerFromEndpoint(ctx, mux, s.opts.Addr, opts)
	if err != nil {
		mux = nil
	}
	return http.Handler(mux), err
}
