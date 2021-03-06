package util

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

var salt = []byte(os.Getenv("SAFARI_SALT"))

func Hash(k string) string {
	if len(salt) == 0 {
		fmt.Println("WARNING: SAFARI_SALT not set. Using DERP")
		salt = []byte("DERP")
	}
	return base64.StdEncoding.EncodeToString(pbkdf2.Key([]byte(k), salt, 4096, 32, sha1.New))
}
