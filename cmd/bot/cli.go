package main

import (
	"log"

	"github.com/buckhx/safari-zone"
)

const (
	pdxAddr = "localhost:50051"
	regAddr = "localhost:50052"
	sfrAddr = "localhost:50053"
)

func main() {
	bot := safaribot.New()
	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}
