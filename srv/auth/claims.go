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

// HasSubScope checks if the scope is in these claims OR if the subject is matchec
func (c Claims) HasSubScope(sub string, scp ...string) bool {
	return c.Subject == sub || c.HasScope(scp...)
}

// HasScope checks that every scope is covered
func (c Claims) HasScope(scp ...string) bool {
	ok := true //
	for _, rs := range scp {
		ok = false
		for _, us := range c.Scope {
			if us == rs {
				ok = true
				break
			}
		}
		if !ok {
			return ok
		}
	}
	return ok
}
