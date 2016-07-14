package main

import (
	"fmt"
	"os"

	"github.com/buckhx/safari-zone/auth"
)

func main() {
	key, err := auth.GenES256Key()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("%s", key)
	}
}
