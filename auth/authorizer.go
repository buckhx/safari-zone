package auth

import (
	"crypto"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/mwitkow/go-grpc-middleware"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// Authorizer verifies authorization for a RPC calls by intercepting request metadata
type Authorizer struct {
	opts Opts
	pub  crypto.PublicKey
}

func NewAuthorizer(opts Opts) (*Authorizer, error) {
	pub, err := opts.fetchCert()
	if err != nil {
		return nil, err
	}
	return &Authorizer{opts: opts, pub: pub}, nil
}

func (a *Authorizer) HandleUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if a.skip(info.FullMethod) {
		return handler(ctx, req)
	}
	ctx, err := a.Context(ctx)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func (a *Authorizer) HandleStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if a.skip(info.FullMethod) {
		return handler(srv, stream)
	}
	wrap := grpc_middleware.WrapServerStream(stream)
	ctx := wrap.Context()
	ctx, err := a.Context(ctx)
	if err != nil {
		return err
	}
	wrap.WrappedContext = ctx
	return handler(srv, stream)
}

// Verify checks a token string and returns a jwt.Token if valid
func (a *Authorizer) Verify(tok string) (*jwt.Token, error) {
	if token, err := jwt.ParseWithClaims(tok, Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return a.pub, nil
	}); err == nil && token.Valid {
		return token, nil
	} else {
		return nil, err
	}
}

// Context validates the context's authorization params and populates claims if there is no error
func (a *Authorizer) Context(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "Authorization required: no context metadata")
	}
	payload, ok := md[AUTH_HEADER]
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "Authorization required: missing header")
	}
	tok := strings.TrimPrefix(payload[0], BEARER_PREFIX)
	token, err := a.Verify(tok)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, err.Error())
	}
	ctx = token.Claims.(Claims).Context(ctx)
	return ctx, nil
}

// skip checks if this method is in the whitelist (skips authorization)
func (a *Authorizer) skip(method string) (ok bool) {
	for _, m := range a.opts.UnsecuredMethods {
		if m == method {
			return true
		}
	}
	return
}
