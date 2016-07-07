package main

import (
	"log"

	"github.com/buckhx/pokedex"
	"github.com/buckhx/pokedex/srv"
)

const (
	rpcAddr = ":50051"
	gwAddr  = ":8080"
)

func main() {
	s := pokedex.New(rpcAddr)
	go func() {
		err := s.Listen()
		log.Println(err)
	}()
	gw := srv.NewGateway(gwAddr, s)
	err := gw.Serve()
	log.Fatal(err)
}
