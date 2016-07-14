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

func (s *Service) GetPokemon(ctx context.Context, req *pbf.Pokemon) (pc *pbf.Pokemon_Collection, err error) {
	pokes := newPokelist()
	if p := s.ByNumber(int(req.Number)); p != nil {
		claims, _ := auth.ClaimsFromContext(ctx)
		if !claims.HasScope("PROFESSOR") { // or in users collection
			f := unknown(p.Number)
			u, err := s.reg.Get(ctx, &pbf.Trainer{Uid: claims.Subject})
			if err != nil {
				return nil, grpc.Errorf(codes.NotFound, "Trainer not registered"+err.Error())
			}
			for _, poke := range u.Pc.Pokemon {
				if poke.Number == req.Number {
					f = p
					break
				}
			}
			p = f
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
