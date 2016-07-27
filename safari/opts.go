package safari

import (
	"fmt"

	"github.com/buckhx/safari-zone/pokedex"
	"github.com/buckhx/safari-zone/registry"
	"github.com/buckhx/safari-zone/srv"
)

type Opts struct {
	srv.Opts
	Registry string
	Pokedex  string
	srvtok   string
}

func (o Opts) RegistryClient() (*registry.Client, error) {
	return registry.Dial(o.Registry)
}

func (o Opts) PokedexClient() (*pokedex.Client, error) {
	if o.srvtok == "" {
		return nil, fmt.Errorf("missing token for pokedex")
	}
	return pokedex.Dial(o.Pokedex, o.srvtok)
}
