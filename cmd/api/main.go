package main

import (
	"fmt"

	"github.com/buckhx/pokedex"
)

func main() {
	api := pokedex.NewServer(":8080")
	err := api.Run()
	fmt.Println(err)
}
