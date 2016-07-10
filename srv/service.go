package srv

import (
	"net/http"

	"github.com/mwitkow/go-grpc-middleware"

	"google.golang.org/grpc"
)

type Service interface {
	Listen() error
	Mux() (http.Handler, error)
}

type Opts struct {
	Auth AuthOpts
}

func NewGRPC(opts Opts) (*grpc.Server, error) {
	auth, err := NewAuthorizer(opts.Auth)
	if err != nil {
		return nil, err
	}
	stm := []grpc.StreamServerInterceptor{auth.HandleStream}
	uni := []grpc.UnaryServerInterceptor{auth.HandleUnary}
	return grpc.NewServer(
		grpc_middleware.WithStreamServerChain(stm...),
		grpc_middleware.WithUnaryServerChain(uni...),
	), nil
}
