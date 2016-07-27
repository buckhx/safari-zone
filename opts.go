package safaribot

import (
	"github.com/buckhx/safari-zone/registry"
	"github.com/buckhx/safari-zone/warden"
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

func (o Opts) DialSafari() (*warden.Client, error) {
	return warden.Dial(o.SafariAddress)
}
