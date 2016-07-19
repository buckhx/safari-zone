package auth

import (
	"os"
	"testing"
	"time"
)

var (
	mint *Mint
	pk   []byte
)

/*
func TestFetchCert(t *testing.T) {
	opts := Opts{CertURI: "../dev/reg.pem"}
	c, err := opts.fetchCert()
	if err != nil {
		t.Errorf("ERR %s", err)
	}
	_ = c
	//t.Errorf("CERT %+v", c)
}
*/

func TestClaimsFromToken(t *testing.T) {
	tok, err := mint.IssueToken("test-sub", 10*time.Minute, "test-scope")
	if err != nil {
		t.Error(err)
	}
	claims, ok := ClaimsFromToken(tok)
	if !ok {
		t.Errorf("Couldn't read claims from token")
	}
	switch {
	case claims.Subject != "test-sub":
		t.Errorf("Invalid claims subject %+v", claims)
	case claims.Issuer != "test-issuer":
		t.Errorf("Invalid claims issuer %+v", claims)
	case claims.Scope[0] != "test-scope":
		t.Errorf("Invalid claims scope %+v", claims)
	}

}

func TestMain(m *testing.M) {
	var err error
	pk, err = GenES256Key()
	if err != nil {
		panic(err)
	}
	mint, err = NewEC256Mint("test-issuer", pk)
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
