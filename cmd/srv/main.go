package main

import (
	"log"

	"github.com/buckhx/pokedex/srv"
)

const (
	port = ":50051"
)

func main() {
	//pokeapi.BaseUrl = "http://localhost:8888"
	//go pokeapi.MockServer(":8888")
	s := srv.New(port)
	err := s.Run()
	log.Fatal(err)
}
