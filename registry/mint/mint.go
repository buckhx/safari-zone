package mint

import (
	"crypto"
	"time"

	"github.com/buckhx/safari-zone/util"
	"github.com/dgrijalva/jwt-go"
)

// TODO mint interface
type Mint struct {
	owner string
	alg   jwt.SigningMethod
	key   crypto.PrivateKey
}

// NewEC creates a new mint w/ a ES256 private key from a .pem file at the given path
func NewEC256(owner string, pem []byte) (m *Mint, err error) {
	k, err := LoadECPrivateKey(pem)
	if err != nil {
		return
	}
	m = &Mint{
		owner: owner,
		alg:   jwt.SigningMethod,
		key:   k,
	}
	return
}

func (m *Mint) IssueToken(sub string, dur time.Duration, scope ...string) (string, error) {
	now := time.Now()
	claims := claims{
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

func (m *Mint) MarshalPublicJwk() ([]byte, err) {
	return MarshalJwkJson(m.owner, m.alg.Alg(), m.key.Public())
}

type claims struct {
	jwt.StandardClaims
	Scope []string `json:"scope,omitempty"`
}
