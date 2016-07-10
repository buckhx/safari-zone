package srv

import (
	"context"
	"crypto"
	"fmt"

	"gopkg.in/square/go-jose.v1"

	"github.com/dgrijalva/jwt-go"
	"github.com/mwitkow/go-grpc-middleware"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

const (
	AUTH_HEADER = "Authorization"
)

// AuthOpts configures a Authorizer
type AuthOpts struct {
	// UnsecuredMethods are grpc method strings that skip authorization
	UnsecuredMethods []string
	// PublicJwkURI is the uri for the publc JWK that verifies access tokens
	PublicJwkURI string
}

// Authorizer verifies authorization for a RPC calls by intercepting request metadata
type Authorizer struct {
	opts AuthOpts
	pub  crypto.PublicKey
}

func NewAuthorizer(opts AuthOpts) (*Authorizer, error) {
	raw := []byte{}
	jwk := &jose.JsonWebKey{}
	if err := jwk.UnmarshalJSON(raw); err != nil {
		// jwk.Valid()
		return nil, err
	}
	pub := jwk.Key.(crypto.PublicKey)
	return &Authorizer{opts: opts, pub: pub}, nil
}

func (a *Authorizer) HandleUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if a.skip(info.FullMethod) {
		return handler(ctx, req)
	}
	ctx, err := a.ValidateContext(ctx)
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
	ctx, err := a.ValidateContext(ctx)
	if err != nil {
		return err
	}
	wrap.WrappedContext = ctx
	return handler(srv, stream)
}

// Verify checks a token string and returns a jwt.Token if valid
func (a *Authorizer) Verify(tok string) (*jwt.Token, error) {
	if token, err := jwt.Parse(tok, func(t *jwt.Token) (interface{}, error) {
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

// ValidateContext checks for valid metadata from a context and adds CtxClaims
func (a *Authorizer) ValidateContext(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "Authorization required")
	}
	tok, ok := md[AUTH_HEADER]
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "Authorization required")
	}
	token, err := a.Verify(tok[0])
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, err.Error())
	}
	ctx = context.WithValue(ctx, CtxClaims, token.Claims)
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
