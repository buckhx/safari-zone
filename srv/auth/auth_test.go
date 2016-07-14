package auth

import "testing"

func TestFetchCert(t *testing.T) {
	opts := Opts{CertURI: "../dev/reg.pem"}
	c, err := opts.fetchCert()
	if err != nil {
		t.Errorf("ERR %s", err)
	}
	_ = c
	//t.Errorf("CERT %+v", c)
}
