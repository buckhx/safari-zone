package main

import (
	"log"

	"github.com/buckhx/safari-zone"
)

func main() {
	err := safaribot.GUI()
	if err != nil {
		log.Fatal(err)
	}
}
