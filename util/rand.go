package util

import (
	"math/rand"
	"strings"

	"gopkg.in/petname.v1"
)

func RandRng(min, max int) int {
	//rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func RandName() string {
	name := strings.Title(petname.Generate(2, " "))
	return strings.Replace(name, " ", "", -1)
}
