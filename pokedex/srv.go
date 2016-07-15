package pokedex

import (
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

type Service struct {
	*Pokedex
	addr string
	opts srv.Opts
	reg  pbf.RegistryClient
}

func NewService(addr string) (s srv.Service, err error) {
	pdx, err := FromCsv("pokedex/pokedex.csv")
	if err != nil {
		return
	}
	s = &Service{
		Pokedex: pdx,
		addr:    addr,
		opts: srv.Opts{
			Auth: auth.Opts{
				CertURI: "http://localhost:8080/registry/v0/cert",
			},
		},
	}
	return
}

func (s *Service) Name() string {
	return "pokedex"
}

func (s *Service) Version() string {
	return "v0"
}

// GetPokemon fetchs the pokemon by number if the trainer has them in their PC
// Professors know all about pokemon and can see details about every pokemon
func (s *Service) GetPokemon(ctx context.Context, req *pbf.Pokemon) (*pbf.Pokemon_Collection, error) {
	pokes := newPokelist()
	num := req.Number
	if p := s.ByNumber(int(num)); p != nil {
		claims, _ := auth.ClaimsFromContext(ctx)
		if !claims.HasScope(registry.ProfScope) { // check if user has pokemon number in PC
			u, err := s.reg.GetTrainer(ctx, &pbf.Trainer{Uid: uid})
			if err != nil {
				return nil, grpc.Errorf(codes.NotFound, "Trainer not registered "+err.Error())
			}
			if !(pokelist{u.Pc}).HasNumber(num) {
				p = unknown(p.Num)
			}
		}
		pokes.Append(p)
	} else {
		return nil, grpc.Errorf(codes.NotFound, "Pokemon not recognized")
	}
	return poke.Pokemon_Collection, nil
}

func (s *Service) Listen() error {
	tcp, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Println(err)
		return err
	}
	rpc, err := srv.ConfigureGRPC(s.opts)
	if err != nil {
		log.Println(err)
		return err
	}
	pbf.RegisterPokedexServer(rpc, s)
	err = s.bindRegistry(registryAddr)
	if err != nil {
		return nil
	}
	log.Printf("Service %T listening at %s", s, s.addr)
	return rpc.Serve(tcp)
}

func (s *Service) Mux() (http.Handler, error) {
	ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pbf.RegisterPokedexHandlerFromEndpoint(ctx, mux, s.addr, opts)
	if err != nil {
		mux = nil
	}
	return http.Handler(mux), err
}

func (s *Service) bindRegistry(addr string) error {
	if conn, err := grpc.Dial(addr, grpc.WithInsecure()); err == nil {
		s.reg = pbf.NewRegistryClient(conn)
	} else {
		return err
	}
	return nil
}
