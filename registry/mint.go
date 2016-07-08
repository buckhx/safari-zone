package registry

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"time"

	"github.com/buckhx/safari-zone/util"
	"github.com/dgrijalva/jwt-go"
)

const (
	Issuer          = "buckhx.safari.registry"
	DefaultTokenDur = 24 * time.Hour
)

type Mint struct {
	signer jwt.SigningMethod
	key    *ecdsa.PrivateKey
}

func NewJwtMint() *Mint {
	k, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	return &Mint{
		signer: jwt.SigningMethodES384,
		key:    k,
	}
}

func (m *Mint) IssueToken(sub string, dur time.Duration, scope ...string) (string, error) {
	now := time.Now()
	claims := claims{
		Scope: scope,
		StandardClaims: jwt.StandardClaims{
			Subject:   sub,
			Issuer:    Issuer,
			ExpiresAt: now.Add(dur).Unix(),
			IssuedAt:  now.Unix(),
			Id:        util.GenUUID(),
		},
	}
	return jwt.NewWithClaims(m.signer, claims).SignedString(m.key)
}

type claims struct {
	jwt.StandardClaims
	Scope []string `json:"scope,omitempty"`
}
