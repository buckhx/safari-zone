package main

import (
	"log"

	"github.com/buckhx/safari-zone"
)

const (
	regAddr = "192.168.99.100:30051"
	//pdxAddr = "localhost:50052"
	wrdAddr = "192.168.99.100:30053"
)

func main() {
	opts := safari.Opts{
		RegistryAddress: regAddr,
		WardenAddress:   wrdAddr,
	}
	c := safari.NewGUI(opts)
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
