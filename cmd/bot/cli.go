package main

import (
	"fmt"
	"time"

	"github.com/buckhx/safari-zone"
)

const (
	pdxAddr = "localhost:50051"
	regAddr = "localhost:50052"
	sfrAddr = "localhost:50053"
)

func main() {
	bot := safaribot.DerpBot()
	for msg := range bot.Msgs {
		fmt.Println(msg)
		time.Sleep(100 * time.Millisecond)
	}
	//log.Fatal(bot.Get())
	/*
		bot := safaribot.New()
		if err := bot.Run(); err != nil {
			log.Fatal(err)
		}
	*/
}
