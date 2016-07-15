package main

import "github.com/buckhx/safari-zone"

const (
	pdxAddr = "localhost:50051"
	regAddr = "localhost:50052"
	sfrAddr = "localhost:50053"
)

func main() {
	bot := safaribot.NewSafariBot()
	bot.Run()
}
