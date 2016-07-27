package main

import (
	"log"

	"github.com/buckhx/safari-zone"
)

const (
	pdxAddr = "localhost:50051"
	regAddr = "localhost:50052"
	safAddr = "localhost:50053"
)

func main() {
	opts := safari.Opts{
		RegistryAddress: regAddr,
		SafariAddress:   safAddr,
	}
	c := safari.NewGUI(opts)
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
