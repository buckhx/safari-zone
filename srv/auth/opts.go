package auth

import (
	"crypto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"gopkg.in/square/go-jose.v1"

	"github.com/buckhx/safari-zone/proto/pbf"
)

const (
	AUTH_HEADER   = "authorization"
	BEARER_PREFIX = "Bearer "
	BASIC_PREFIX  = "Basic "
)

// AuthOpts configures a Authorizer
type Opts struct {
	// UnsecuredMethods are grpc method strings that skip authorization
	UnsecuredMethods []string
	// CertURI is the uri for the publc JWK that verifies access tokens
	CertURI string
}

func (o Opts) fetchCert() (pub crypto.PublicKey, err error) {
	switch {
	//case strings.HasPrefix(o.CertURI, "https"):
	case strings.HasPrefix(o.CertURI, "http"):
		r, e := http.Get(o.CertURI)
		if e != nil {
			err = e
			break
		}
		if r.StatusCode != http.StatusOK {
			err = fmt.Errorf("CertURI not OK: %d", r.StatusCode)
			break
		}
		defer r.Body.Close()
		raw, e := ioutil.ReadAll(r.Body)
		if e != nil {
			err = e
			break
		}
		cert := &pbf.Cert{}
		if err = json.Unmarshal(raw, cert); err != nil {
			break
		}
		jwk := &jose.JsonWebKey{}
		if err = jwk.UnmarshalJSON(cert.Jwk); err != nil {
			// jwk.Valid()
			break
		}
		var ok bool
		if pub, ok = jwk.Key.(crypto.PublicKey); !ok {
			err = fmt.Errorf("JWK.Key not a crypto.PublicKey")
		}
		//case strings.HasPrefix(o.CertURI, "http"):
		// TODO verify that this is the correct behavior (HTTPS required to fetch cert)
		//err = fmt.Errorf("HTTPS required for network AuthOpts.CertURI")
	case exists(o.CertURI):
		f, e := os.Open(o.CertURI)
		if e != nil {
			err = e
			break
		}
		raw, e := ioutil.ReadAll(f)
		if e != nil {
			err = e
			break
		}
		if key, e := LoadECPrivateKey(raw); e == nil {
			pub = key.Public()
		} else {
			err = e
		}
	default:
		err = fmt.Errorf("AuthOpts.CertURI must be a local file or HTTPS network resource")
	}
	return
}
