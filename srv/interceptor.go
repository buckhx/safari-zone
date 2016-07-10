package srv

type CtxKey int

const (
	CtxClaims = iota
)

/*
type Interceptor interface {
	UnaryHandler(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error)
	StreamHandler() grpc.StreamServerInterceptor
}
*/
