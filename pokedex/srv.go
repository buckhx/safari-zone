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
	opts Opts
	reg  pbf.RegistryClient
}

func NewService(opts Opts) (s srv.Service, err error) {
	pdx, err := opts.LoadData()
	if err != nil {
		return
	}
	reg, err := opts.RegistryClient()
	if err != nil {
		return
	}
	s = &Service{
		Pokedex: pdx,
		reg:     reg,
		opts:    opts,
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
		if !claims.HasRole(registry.ProfScope, registry.ServiceScope) {
			// check if user has pokemon number in PC
			u, err := s.reg.GetTrainer(ctx, &pbf.Trainer{Uid: claims.Subject})
			if err != nil {
				return nil, grpc.Errorf(codes.NotFound, "Trainer not registered %s", err)
			}
			if !(pokelist{u.Pc}).HasNumber(num) {
				p = unknown(p.Number)
			}
		}
		pokes.Append(p)
	} else {
		return nil, grpc.Errorf(codes.NotFound, "Pokemon not recognized")
	}
	return pokes.Pokemon_Collection, nil
}

func (s *Service) Listen() error {
	tcp, err := net.Listen("tcp", s.opts.Address)
	if err != nil {
		return err
	}
	rpc, err := srv.ConfigureGRPC(s.opts)
	if err != nil {
		return err
	}
	pbf.RegisterPokedexServer(rpc, s)
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
