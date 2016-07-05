package srv

import (
	"net"
	"net/http"

	"github.com/buckhx/pokedex/lib"
	"github.com/buckhx/pokedex/pbf"
	"github.com/buckhx/pokedex/pokeapi"
	"github.com/gengo/grpc-gateway/runtime"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Service struct {
	addr string
	api  *pokeapi.Client
	c    *pokedex.Cache
}

func New(addr string) *Service {
	return &Service{
		addr: addr,
		api:  pokeapi.NewClient(),
		c:    pokedex.NewCache(),
	}
}

func (s *Service) GetPokemon(ctx context.Context, req *pbf.Pokemon_Query) (p *pbf.Pokemon, err error) {
	if v := s.c.GetProto(req); v != nil { //|| err != nil {
		p = &pbf.Pokemon{}
		err = proto.Unmarshal(v, p)
		return
	}
	r, err := s.api.FetchPokemon(int(req.ID))
	if err != nil {
		return
	}
	p = &pbf.Pokemon{ID: int32(r.ID), Name: r.Name}
	go s.c.SetProto(req, p)
	return
}

func (s *Service) Listen() error {
	tcp, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	rpc := grpc.NewServer()
	pbf.RegisterPokedexServer(rpc, s)
	return rpc.Serve(tcp)
}

func (s *Service) Gateway(addr string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pbf.RegisterPokedexHandlerFromEndpoint(ctx, mux, s.addr, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8080", mux)
}
