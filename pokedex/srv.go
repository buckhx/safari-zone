package pokedex

import (
	"log"
	"net"
	"net/http"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/srv"
	"github.com/buckhx/safari-zone/srv/auth"
	"github.com/gengo/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Service struct {
	*Pokedex
	addr string
	opts srv.Opts
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

func (s *Service) GetPokemon(ctx context.Context, req *pbf.Pokemon) (pc *pbf.Pokemon_Collection, err error) {
	pokes := newPokelist()
	if p := s.ByNumber(int(req.Number)); p != nil {
		claims, _ := auth.ClaimsFromContext(ctx)
		if !claims.HasScope("PROFESSOR") { // or in users collection
			p = unknown(p.Number)
		}
		pokes.Append(p)
	}
	if pokes.Empty() {
		return nil, grpc.Errorf(codes.NotFound, "Pokemon not recognized")
	}
	pc = pokes.Pokemon_Collection
	return
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
