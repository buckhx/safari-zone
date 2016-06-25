package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/net/context"

	"github.com/buckhx/pokedex/srv"
)

const (
	address = "localhost:50051"
)

func main() {
	c, err := srv.NewClient(address)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer c.Close()
	scanner := bufio.NewScanner(os.Stdin)
	log.Printf("Waiting for input...")
	for scanner.Scan() {
		in := scanner.Text()
		id, err := strconv.Atoi(in)
		if err != nil {
			log.Printf("Invalid Pokemon ID %q", in)
			continue
		}
		r, err := c.GetPokemon(context.Background(), &srv.PokemonQuery{ID: int32(id)})
		if err != nil {
			log.Printf("Error getting: %v", err)
			continue
		}
		msg, _ := json.Marshal(r)
		fmt.Println(string(msg))
	}
}
