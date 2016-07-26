package pokedex

import (
	"fmt"
	"strings"

	"github.com/buckhx/safari-zone/registry"
	"github.com/buckhx/safari-zone/srv"
)

type Opts struct {
	srv.Opts
	Registry string
	Data     string
}

func (o Opts) RegistryClient() (reg *registry.Client, err error) {
	return registry.Dial(o.Registry) //, tok string)
}

func (o Opts) LoadData() (pdx *Pokedex, err error) {
	switch {
	case strings.HasSuffix(o.Data, ".csv"):
		pdx, err = FromCsv(o.Data)
	default:
		err = fmt.Errorf("Invalid data path")
	}
	return
}
