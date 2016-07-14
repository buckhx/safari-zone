package srv

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Interceptor interface {
	HandleUnary(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error)
	HandleStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error
}
