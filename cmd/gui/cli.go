package main

import (
	"log"

	"github.com/buckhx/safari-zone"
)

const (
	regAddr = "localhost:50051"
	//pdxAddr = "localhost:50052"
	wrdAddr = "localhost:50053"
)

func main() {
	opts := safari.Opts{
		RegistryAddress: regAddr,
		SafariAddress:   wrdAddr,
	}
	c := safari.NewGUI(opts)
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
