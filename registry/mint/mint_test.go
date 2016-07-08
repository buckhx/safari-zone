package mint

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	pemlib "encoding/pem"
	"fmt"
	"testing"

	"gopkg.in/square/go-jose.v1"
)

func TestGenKey(t *testing.T) {

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Error(err)
	}
	der, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(der))
	pem := bytes.NewBuffer(nil)
	pemlib.Encode(pem, &pemlib.Block{Type: "EC PRIVATE KEY", Bytes: der})
	ikey, err := jose.LoadPrivateKey(pem.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	pk := ikey.(*ecdsa.PrivateKey)
	//pub, _ := x509.MarshalPKIXPublicKey(pk.Public())
	jwk := jose.JsonWebKey{
		Key:       pk.Public(),
		KeyID:     "derp",
		Algorithm: string(jose.ES256),
		Use:       "sig",
	}
	o, _ := jwk.MarshalJSON()
	fmt.Println(string(o))
	/*

		fmt.Println(string(buf.Bytes()))
		p, _ := pem.Decode(buf.Bytes())
		pk, err := x509.ParseECPrivateKey(p.Bytes)
		if err != nil {
			t.Error(err)
		}
		pub, err := x509.MarshalPKIXPublicKey(pk.Public())
		if err != nil {
			t.Error(err)
		}
		fmt.Println(o)
	*/
}
