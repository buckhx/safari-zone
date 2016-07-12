package util

import (
	"fmt"
	"math/rand"
)

// GenUID generates a short 4-byte unique identifier
func GenUID() (uuid string) {
	u := RawUUID()
	uuid = fmt.Sprintf("%x", u[0:4])
	return
}

// GenUUID generates a RFC4122 UUID
func GenUUID() (uuid string) {
	u := RawUUID()
	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// Generates 16 random bytes that compliant w/ RFC4122
func RawUUID() []byte {
	u := make([]byte, 16)
	rand.Read(u)
	u[6] = (u[6] | 0x40) & 0x4F
	u[6] = (u[6] | 0x40) & 0x4F
	return u
}
