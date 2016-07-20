package main

import (
	"log"

	"github.com/buckhx/safari-zone"
)

func main() {
	c := safaribot.NewGUI()
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
