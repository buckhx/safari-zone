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
	opts := safaribot.Opts{
		RegistryAddress: regAddr,
		SafariAddress:   safAddr,
	}
	c := safaribot.NewGUI(opts)
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
