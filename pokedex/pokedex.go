package pokedex

import (
	"log"
	"net"
	"net/http"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/gengo/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Pokedex struct {
	addr string
}

func New(addr string) *Pokedex {
	return &Pokedex{
		addr: addr,
	}
}

func (s *Pokedex) GetPokemon(ctx context.Context, req *pbf.Pokemon) (p *pbf.Pokemon_Collection, err error) {
	p = &pbf.Pokemon_Collection{[]*pbf.Pokemon{
		&pbf.Pokemon{Number: req.Number, Name: "derp"}},
	}
	return
}

func (s *Pokedex) Listen() error {
	tcp, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Println(err)
		return err
	}
	rpc := grpc.NewServer()
	pbf.RegisterPokedexServer(rpc, s)
	log.Printf("%T listening at %s", s, s.addr)
	return rpc.Serve(tcp)
}

func (s *Pokedex) Mux() (http.Handler, error) {
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
