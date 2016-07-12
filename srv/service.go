package srv

import (
	"log"
	"net/http"

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
	Auth AuthOpts
}

// ConfigureGRPC configures a GRPC Server w/ the given opts
func ConfigureGRPC(opts Opts) (*grpc.Server, error) {
	auth, err := NewAuthorizer(opts.Auth)
	if err != nil {
		return nil, err
	}
	return NewGRPC(auth), nil
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
