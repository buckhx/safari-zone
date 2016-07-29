package srv

import (
	"fmt"
	"net/http"
)

const (
	API_VERSION = "v0"
	DOCS_ROUTE  = "/docs/"
)

type Route struct {
	Path    string
	Handler http.Handler
	//Docs Route
	//Name string
}

type Gateway struct {
	Address string
	Routes  []Route
}

func (gw Gateway) Mux() http.Handler {
	m := http.NewServeMux()
	for _, rte := range gw.Routes {
		fmt.Println("Registering Route:", rte.Path)
		m.Handle(rte.Path+"/", http.StripPrefix(rte.Path, rte.Handler))
	}
	return http.Handler(m)
}

func (gw Gateway) Serve() error {
	m := gw.Mux()
	return http.ListenAndServe(gw.Address, m)
}
