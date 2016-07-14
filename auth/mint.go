package auth

import (
	"crypto/ecdsa"
	"time"

	"github.com/buckhx/safari-zone/util"
	"github.com/dgrijalva/jwt-go"
)

// TODO mint interface
type Mint struct {
	owner string
	alg   jwt.SigningMethod
	key   *ecdsa.PrivateKey
}

// NewEC creates a new mint w/ a ES256 private key from a .pem file at the given path
func NewEC256Mint(owner string, pem []byte) (m *Mint, err error) {
	key, err := LoadECPrivateKey(pem)
	if err != nil {
		return
	}
	m = &Mint{
		owner: owner,
		alg:   jwt.SigningMethodES256,
		key:   key,
	}
	return
}

func (m *Mint) IssueToken(sub string, dur time.Duration, scope ...string) (string, error) {
	now := time.Now()
	claims := Claims{
		Scope: scope,
		StandardClaims: jwt.StandardClaims{
			Subject:   sub,
			Issuer:    m.owner,
			ExpiresAt: now.Add(dur).Unix(),
			IssuedAt:  now.Unix(),
			Id:        util.GenUUID(),
		},
	}
	return jwt.NewWithClaims(m.alg, claims).SignedString(m.key)
}

func (m *Mint) MarshalPublicJwk() ([]byte, error) {
	return MarshalJwkJSON(m.owner, m.alg.Alg(), m.key.Public())
}
