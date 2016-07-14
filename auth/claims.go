package auth

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
)

type CtxKey int

const (
	CtxClaims = iota
)

type Claims struct {
	jwt.StandardClaims
	Scope []string `json:"scope,omitempty"`
}

func ClaimsFromContext(ctx context.Context) (Claims, bool) {
	c, ok := ctx.Value(CtxClaims).(Claims)
	return c, ok
}

func (c Claims) Context(ctx context.Context) context.Context {
	return context.WithValue(ctx, CtxClaims, c)
}
