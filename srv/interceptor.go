package srv

import (
	"context"

	"google.golang.org/grpc"
)

type CtxKey int

const (
	CtxClaims = iota
)

type Interceptor interface {
	UnaryInterceptor(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error)
	StreamingInterceptor(interface{}, grpc.ServerStream, *grpc.StreamServerInfo, grpc.StreamHandler) error
}
