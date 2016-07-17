package auth

import (
	"encoding/base64"
	"os"
	"strings"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
func decodeTokBlock(block string) ([]byte, error) {
	if l := len(block) % 4; l > 0 {
		block += strings.Repeat("=", 4-l)
	}
	return base64.URLEncoding.DecodeString(block)
}
