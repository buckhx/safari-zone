package main

import (
	"log"

	"github.com/buckhx/pokedex/srv"
)

const (
	port = ":50051"
	gw   = ":8080"
)

func main() {
	//pokeapi.BaseUrl = "http://localhost:8888"
	//go pokeapi.MockServer(":8888")
	s := srv.New(port)
	go s.Listen()
	err := s.Gateway(gw)
	log.Fatal(err)
}
