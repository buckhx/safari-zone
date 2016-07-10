package srv

import (
	"net/http"

	"google.golang.org/grpc"
)

type Service interface {
	Listen() error
	Mux() (http.Handler, error)
}

type Opts struct {
	Auth AuthOpts
}

func NewGRPC(opts Opts) *grpc.Server {
	return grpc.NewServer()
	/*
		auth := NewAuthorizer(opts.Auth)
		stm_opts := []grpc.ServerOption{auth.StreamInterceptor()}
		uni_opts := []grpc.ServerOption{auth.UnaryInterceptor()}
		return grpc.NewServer(
			grpc_middleware.WithStreamServerChain(stm_opts...),
			grpc_middleware.WithUnaryServerChain(uni_tops...),
		)
	*/
}
