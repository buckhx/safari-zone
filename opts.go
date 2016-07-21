package safaribot

import (
	"github.com/buckhx/safari-zone/registry"
	"github.com/buckhx/safari-zone/safari"
)

type Opts struct {
	SafariAddress   string
	RegistryAddress string
	/*
		Safari struct {
			Addr string
		}
		Registry struct {
			Addr string
		}
	*/
}

func (o Opts) DialRegistry() (*registry.Client, error) {
	return registry.Dial(o.RegistryAddress)
}

func (o Opts) DialSafari() (*safari.Client, error) {
	return safari.Dial(o.SafariAddress)
}
