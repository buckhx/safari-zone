package main

import (
	"log"

	"github.com/buckhx/safari-zone/pokedex"
	"github.com/buckhx/safari-zone/registry"
	"github.com/buckhx/safari-zone/safari"
	"github.com/buckhx/safari-zone/srv"
)

const (
	pdxAddr = ":50051"
	regAddr = ":50052"
	sfrAddr = ":50053"
	gwAddr  = ":8080"
	pemfile = "dev/reg.pem"
)

func main() {
	pdx, err := pokedex.NewService(pdxAddr)
	if err != nil {
		log.Fatal(err)
	}
	reg, err := registry.NewService(pemfile, regAddr)
	if err != nil {
		log.Fatal(err)
	}
	sfr, err := safari.NewService(sfrAddr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err := pdx.Listen()
		log.Println(err)
	}()
	go func() {
		err := reg.Listen()
		log.Println(err)
	}()
	go func() {
		err := sfr.Listen()
		log.Println(err)
	}()
	gw := srv.NewGateway(gwAddr, pdx, reg, sfr)
	err = gw.Serve()
	log.Fatal(err)
}

func runPdx() {

}
