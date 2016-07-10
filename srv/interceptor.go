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
	HandleUnary(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error)
	HandleStream(interface{}, grpc.ServerStream, *grpc.StreamServerInfo, grpc.StreamHandler) error
}
