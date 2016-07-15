package srv

import (
	"log"
	"net/http"

	"github.com/buckhx/safari-zone/srv/auth"
	"github.com/mwitkow/go-grpc-middleware"

	"google.golang.org/grpc"
)

type Service interface {
	Listen() error
	Mux() (http.Handler, error)
	Name() string
	Version() string
}

type Opts struct {
	Addr string
	Auth auth.Opts
}

// ConfigureGRPC configures a GRPC Server w/ the given opts
func ConfigureGRPC(opts Opts) (*grpc.Server, error) {
	a, err := auth.NewAuthorizer(opts.Auth)
	if err != nil {
		return nil, err
	}
	return NewGRPC(a), nil
}

// NewGRPC builds a grpc.Server w/ the given inceptors registered in order
func NewGRPC(incpts ...Interceptor) *grpc.Server {
	stm := make([]grpc.StreamServerInterceptor, len(incpts))
	uni := make([]grpc.UnaryServerInterceptor, len(incpts))
	for i, incpt := range incpts {
		log.Printf("Registering Interceptor - %T", incpt)
		stm[i] = incpt.HandleStream
		uni[i] = incpt.HandleUnary
	}
	return grpc.NewServer(
		grpc_middleware.WithStreamServerChain(stm...),
		grpc_middleware.WithUnaryServerChain(uni...),
	)
}
