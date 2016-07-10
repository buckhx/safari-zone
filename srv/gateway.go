package srv

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	API_VERSION = "v0"
	DOCS_ROUTE  = "/docs/"
)

type Gateway struct {
	addr string
	srvs []Service
}

func NewGateway(addr string, srvs ...Service) Gateway {
	return Gateway{
		addr: addr,
		srvs: srvs,
	}
}

func (gw Gateway) Serve() error {
	r := http.NewServeMux()
	log.Println("Registering docs at", DOCS_ROUTE)
	r.HandleFunc(DOCS_ROUTE, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./pbf/docs/static.html")
	})
	for _, srv := range gw.srvs {
		h, err := srv.Mux()
		if err != nil {
			return err
		}
		t := strings.ToLower(strings.Split(fmt.Sprintf("%T", srv), ".")[1])
		pre := fmt.Sprint("/", t, "/", API_VERSION)
		log.Printf("Registering service %T at %s", srv, pre)
		r.Handle(pre+"/", http.StripPrefix(pre, h))
	}
	log.Println("Starting Service Gateway at", gw.addr)
	return http.ListenAndServe(gw.addr, r) //mux)
}
