package auth

import (
	"encoding/json"
	"strings"

	"google.golang.org/grpc/metadata"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
)

type Claims struct {
	jwt.StandardClaims
	Scope []string `json:"scope,omitempty"`
}

func ClaimsFromContext(ctx context.Context) (Claims, bool) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return Claims{}, false
	}
	hd, ok := md[AUTH_HEADER]
	if !ok {
		return Claims{}, false
	}
	tok := strings.TrimPrefix(hd[0], BEARER_PREFIX)
	return ClaimsFromToken(tok)
}

// ClaimsFromToken reads the claims from a token string.
// It DOES NOT verify the signature
func ClaimsFromToken(token string) (c Claims, ok bool) {
	blocks := strings.Split(token, ".")
	if len(blocks) != 3 {
		return
	}
	raw, err := decodeTokBlock(blocks[1])
	if err != nil {
		return
	}
	if err := json.Unmarshal(raw, &c); err == nil {
		ok = true
	}
	return
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

// HasRole checks if any of these roles are in the claims
func (c Claims) HasRole(roles ...string) bool {
	for _, role := range roles {
		for _, scp := range c.Scope {
			if scp == role {
				return true
			}
		}
	}
	return false
}
