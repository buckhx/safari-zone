package pokedex

import (
	"log"
	"net"
	"net/http"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/srv"
	"github.com/gengo/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type PokedexSrv struct {
	*Pokedex
	addr string
}

func NewService(addr string) (s srv.Service, err error) {
	pdx, err := FromCsv("pokedex/pokedex.csv")
	if err != nil {
		return
	}
	s = &PokedexSrv{
		Pokedex: pdx,
		addr:    addr,
	}
	return
}

func (s *PokedexSrv) GetPokemon(ctx context.Context, req *pbf.Pokemon) (pc *pbf.Pokemon_Collection, err error) {
	if p := s.ByNumber(int(req.Number)); p != nil {
		pc = &pbf.Pokemon_Collection{
			Pokemon: []*pbf.Pokemon{p},
		}
	}
	return
}

func (s *PokedexSrv) Listen() error {
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

func (s *PokedexSrv) Mux() (http.Handler, error) {
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
