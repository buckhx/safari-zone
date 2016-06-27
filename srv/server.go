package srv

import (
	"net"

	"github.com/buckhx/pokedex"
	"github.com/buckhx/pokedex/pokeapi"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Server struct {
	addr string
	api  *pokeapi.Client
	c    *pokedex.Cache
}

func New(addr string) *Server {
	return &Server{
		addr: addr,
		api:  pokeapi.NewClient(),
		c:    pokedex.NewCache(),
	}
}

func (s *Server) GetPokemon(ctx context.Context, req *PokemonQuery) (p *Pokemon, err error) {
	if v := s.c.GetProto(req); v != nil { //|| err != nil {
		p = &Pokemon{}
		err = proto.Unmarshal(v, p)
		return
	}
	r, err := s.api.FetchPokemon(int(req.ID))
	if err != nil {
		return
	}
	p = &Pokemon{ID: int32(r.ID), Name: r.Name}
	go s.c.SetProto(req, p)
	return
}

func (s *Server) Run() error {
	tcp, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	rpc := grpc.NewServer()
	RegisterPokedexServer(rpc, s)
	return rpc.Serve(tcp)
}
