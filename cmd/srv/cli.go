package main

import (
	"fmt"
	"os"

	"github.com/buckhx/safari-zone/pokedex"
	"github.com/buckhx/safari-zone/registry"
	"github.com/buckhx/safari-zone/srv"
	"github.com/urfave/cli"
)

var Version string

func main() {
	app := cli.NewApp()
	app.Name = "Safari Zone Services"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name: "pokedex",
			Action: func(c *cli.Context) error {
				addr := fmt.Sprint(":", c.String("port"))
				reg := c.String("registry")
				data := c.String("data")
				pdx, err := pokedex.NewService(pokedex.Opts{
					Opts:     srv.Opts{Address: addr},
					Registry: reg,
					Data:     data,
				})
				if err != nil {
					return err
				}
				done := make(chan error)
				go func() {
					done <- pdx.Listen()
				}()
				go func() {
					gw := fmt.Sprint(":", c.String("gateway"))
					done <- srv.NewGateway(gw, pdx).Serve()
				}()
				return <-done
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "p, port",
					Value:  "50051",
					EnvVar: "POKEDEX_PORT",
				},
				cli.StringFlag{
					Name:   "gw, gateway",
					Value:  "8080",
					EnvVar: "POKEDEX_GATEWAY",
				},
				cli.StringFlag{
					Name:   "r, registry",
					Value:  "localhost:50052",
					EnvVar: "POKEDEX_REGISTRY",
				},
				cli.StringFlag{
					Name:   "d, data",
					Value:  "pokedex.csv",
					EnvVar: "POKEDEX_DATA",
				},
			},
		},
		{
			Name: "registry",
			Action: func(c *cli.Context) error {
				pem := c.String("key")
				addr := fmt.Sprint(":", c.String("port"))
				reg, err := registry.NewService(registry.Opts{
					Opts:    srv.Opts{Address: addr},
					KeyPath: pem,
				})
				if err != nil {
					return err
				}
				done := make(chan error)
				go func() {
					done <- reg.Listen()
				}()
				go func() {
					gw := fmt.Sprint(":", c.String("gateway"))
					done <- srv.NewGateway(gw, reg).Serve()
				}()
				return <-done
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "p, port",
					Value:  "50051",
					EnvVar: "REGISTRY_PORT",
				},
				cli.StringFlag{
					Name:   "gw, gateway",
					Value:  "8080",
					EnvVar: "REGISTRY_GATEWAY",
				},
				cli.StringFlag{
					Name:   "k, key",
					Value:  "reg.pem",
					EnvVar: "REGISTRY_KEY",
					Usage:  "Path to the private key .pem for token signing",
				},
			},
		},
		{
			Name: "safari",
			Action: func(c *cli.Context) error {
				addr := fmt.Sprint("", c.String("port"))
				reg := c.String("registry")
				pdx := c.String("pokedex")
				fmt.Printf("addr: %s, reg: %s, data: %s", addr, reg, pdx)
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "p, port",
					Value:  "50051",
					EnvVar: "SAFARI_PORT",
				},
				cli.StringFlag{
					Name:   "r, registry",
					Value:  "localhost:50052",
					EnvVar: "SAFARI_REGISTRY",
				},
				cli.StringFlag{
					Name:   "pokedex",
					Value:  "localhost:50053",
					EnvVar: "SAFARI_POKEDEX",
				},
			},
		},
	}
	app.Run(os.Args)
}

/*
const (
	pdxAddr = ":50051"
	regAddr = ":50052"
	sfrAddr = ":50053"
	gwAddr  = ":8080"
	pemfile = "dev/reg.pem"
)
func main() {
	reg, err := registry.NewService(pemfile, regAddr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err := reg.Listen()
		log.Println(err)
	}()
	pdx, err := pokedex.NewService(pdxAddr)
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
		err := sfr.Listen()
		log.Println(err)
	}()
	gw := srv.NewGateway(gwAddr, pdx, reg, sfr)
	err = gw.Serve()
	log.Fatal(err)
}

func runPdx() {

}
*/
