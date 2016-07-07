package main

import (
	"log"

	"github.com/buckhx/safari-zone/pokedex"
	"github.com/buckhx/safari-zone/srv"
)

const (
	rpcAddr = ":50051"
	gwAddr  = ":8080"
)

func main() {
	s, err := pokedex.NewService(rpcAddr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err := s.Listen()
		log.Println(err)
	}()
	gw := srv.NewGateway(gwAddr, s)
	err = gw.Serve()
	log.Fatal(err)
}
