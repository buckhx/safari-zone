package srv

import "net/http"

type Service interface {
	Listen() error
	Mux() (http.Handler, error)
}
