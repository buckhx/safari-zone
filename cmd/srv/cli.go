package main

import (
	"fmt"
	"os"

	"github.com/buckhx/safari-zone/registry"
	"github.com/urfave/cli"
)

const (
	pdxAddr = ":50051"
	regAddr = ":50052"
	sfrAddr = ":50053"
	gwAddr  = ":8080"
	pemfile = "dev/reg.pem"
)

func main() {
	app := cli.NewApp()
	app.Name = "Safari Zone Services"
	app.Commands = []cli.Command{
		{
			Name: "pokedex",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p, port",
					Value: "50051",
				},
			},
		},
		{
			Name: "registry",
			Action: func(c *cli.Context) error {
				pem := c.String("key")
				addr := fmt.Sprint(":", c.String("port"))
				reg, err := registry.NewService(pem, addr)
				if err != nil {
					return err
				}
				return reg.Listen()
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p, port",
					Value: "50052",
				},
				cli.StringFlag{
					Name:  "k, key",
					Value: "reg.pem",
					Usage: "Path to the private key .pem for token signing",
				},
			},
		},
		{
			Name: "safari",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p, port",
					Value: "50053",
				},
			},
		},
	}
	app.Run(os.Args)
}

/*
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
