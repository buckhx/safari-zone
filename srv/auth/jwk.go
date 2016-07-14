package auth

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	pemlib "encoding/pem"
	"fmt"

	"gopkg.in/square/go-jose.v1"
)

func LoadECPrivateKey(pem []byte) (*ecdsa.PrivateKey, error) {
	//k, err := jose.LoadPrivateKey(pem)
	block, _ := pemlib.Decode(pem)
	if block == nil {
		return nil, fmt.Errorf("No pem block")
	}
	return x509.ParseECPrivateKey(block.Bytes)
}

func GenES256Key() (pem []byte, err error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return
	}
	der, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(nil)
	err = pemlib.Encode(buf, &pemlib.Block{Type: "EC PRIVATE KEY", Bytes: der})
	pem = buf.Bytes()
	return
}

func MarshalJwkJSON(kid, alg string, key interface{}) ([]byte, error) {
	jwk := jose.JsonWebKey{
		Key:       key,
		KeyID:     kid,
		Algorithm: string(jose.ES256),
		Use:       "sig",
	}
	return jwk.MarshalJSON()
}
