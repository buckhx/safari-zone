package main

import (
	"fmt"
	"os"

	"github.com/buckhx/safari-zone/registry/mint"
)

func main() {
	key, err := mint.GenES256Key()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("%s", key)
	}
}
