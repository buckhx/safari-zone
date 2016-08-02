package main

import (
	"os"

	"github.com/buckhx/safari-zone"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Safari Zone GUI"
	//app.Version = Version
	app.Action = func(c *cli.Context) error {
		opts := safari.Opts{
			RegistryAddress: c.String("registry"),
			WardenAddress:   c.String("warden"),
		}
		return safari.NewGUI(opts).Run()
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "registry",
			Value:  "localhost:50051",
			EnvVar: "REGISTRY_ADDR",
		},
		cli.StringFlag{
			Name:   "warden",
			Value:  "localhost:50053",
			EnvVar: "WARDEN_ADDR",
		},
	}
	app.Run(os.Args)
}
