package srv

import (
	"net"

	"github.com/buckhx/pokedex/pokeapi"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Server struct {
	addr string
	api  *pokeapi.Client
	c    *Cache
}

func New(addr string) *Server {
	return &Server{
		addr: addr,
		api:  pokeapi.NewClient(),
		c:    newCache(),
	}
}

func (s *Server) GetPokemon(ctx context.Context, req *PokemonQuery) (*Pokemon, error) {
	p, err := s.api.FetchPokemon(int(req.ID))
	if err != nil {
		return nil, err
	}
	return &Pokemon{ID: int32(p.ID), Name: p.Name}, nil
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
