package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/buckhx/safari-zone"
	"github.com/buckhx/safari-zone/util/bot"
)

const (
	pdxAddr = "localhost:50051"
	regAddr = "localhost:50052"
	safAddr = "localhost:50053"
)

func main() {
	safbot := safari.NewBot(safari.Opts{
		RegistryAddress: regAddr,
		WardenAddress:   safAddr,
	})
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			safbot.Send(bot.Cmd(scanner.Text()))
		}
	}()
	for msg := range safbot.Msgs {
		fmt.Println(msg)
	}
}
