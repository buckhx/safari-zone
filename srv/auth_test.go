package srv

import "testing"

func TestFetchCert(t *testing.T) {
	opts := AuthOpts{CertURI: "../dev/reg.pem"}
	c, err := opts.fetchCert()
	if err != nil {
		t.Errorf("ERR %s", err)
	}
	_ = c
	//t.Errorf("CERT %+v", c)
}
