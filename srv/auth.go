package srv

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/buckhx/safari-zone/registry/mint"
	"github.com/dgrijalva/jwt-go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

const (
	AUTH_HEADER = "Authorization"
	CTX_CLAIMS  = "claims"
)

type Authorizer struct {
	pub *ecdsa.PublicKey
}

func (a *Authorizer) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "Authorization required")
	}
	tok, ok := md[AUTH_HEADER]
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "Authorization required")
	}
	token, err := a.Validate(tok[0])
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, err.Error())
	}
	_ = token

	ctx = context.WithValue(ctx, "claims", token.Claims.(mint.Claims))
	return handler(ctx, req)
}

func (a *Authorizer) Validate(tok string) (*jwt.Token, error) {
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
