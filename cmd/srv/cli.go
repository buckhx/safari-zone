package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/buckhx/safari-zone/pokedex"
	"github.com/buckhx/safari-zone/registry"
	"github.com/buckhx/safari-zone/srv"
	"github.com/buckhx/safari-zone/warden"
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
				return serve(c, pdx)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p, port",
					Value: "50051",
					//EnvVar: "POKEDEX_PORT",
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
				return serve(c, reg)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p, port",
					Value: "50051",
					//EnvVar: "REGISTRY_PORT",
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
			Name: "warden",
			Action: func(c *cli.Context) error {
				addr := fmt.Sprint(":", c.String("port"))
				reg := c.String("registry")
				pdx := c.String("pokedex")
				wrdn, err := warden.NewService(warden.Opts{
					Opts:     srv.Opts{Address: addr},
					Registry: reg,
					Pokedex:  pdx,
				})
				if err != nil {
					return err
				}
				return serve(c, wrdn)
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p, port",
					Value: "50051",
					//EnvVar: "SAFARI_PORT",
				},
				cli.StringFlag{
					Name:  "r, registry",
					Value: "localhost:50052",
					//EnvVar: "SAFARI_REGISTRY",
				},
				cli.StringFlag{
					Name:   "pokedex",
					Value:  "localhost:50053",
					EnvVar: "SAFARI_POKEDEX",
				},
			},
		},
		{
			Name: "gateway",
			Action: func(c *cli.Context) error {
				addr := fmt.Sprint(":", c.String("port"))
				reg, err := registry.Mux(c.String("registry"))
				if err != nil {
					return err
				}
				pdx, err := pokedex.Mux(c.String("pokedex"))
				if err != nil {
					return err
				}
				wrd, err := warden.Mux(c.String("warden"))
				if err != nil {
					return err
				}
				gw := srv.Gateway{
					Address: addr,
					Routes: []srv.Route{
						{
							Path:    "/registry",
							Handler: reg,
						},
						{
							Path:    "/pokedex",
							Handler: pdx,
						},
						{
							Path:    "/warden",
							Handler: wrd,
						},
						{
							Path:    "/ping",
							Handler: pinger{},
						},
					},
				}
				return gw.Serve()
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p, port",
					Value: "8080",
					//EnvVar: "SAFARI_PORT",
				},
				cli.StringFlag{
					Name:  "registry",
					Value: "localhost:50051",
					//EnvVar: "SAFARI_POKEDEX",
				},
				cli.StringFlag{
					Name:  "pokedex",
					Value: "localhost:50052",
					//EnvVar: "SAFARI_REGISTRY",
				},
				cli.StringFlag{
					Name:  "warden",
					Value: "localhost:50053",
					//EnvVar: "SAFARI_POKEDEX",
				},
			},
		},
	}
	app.Run(os.Args)
}

func serve(c *cli.Context, svc srv.Service) error {
	done := make(chan error)
	go func() {
		done <- svc.Listen()
	}()
	return <-done
}

type pinger struct{}

func (p pinger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
