package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"

	"github.com/buckhx/safari-zone/proto/pbf"

	"golang.org/x/net/context"
)

const (
	addr = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbf.NewPokedexClient(conn)
	scanner := bufio.NewScanner(os.Stdin)
	log.Printf("Waiting for input...")
	for scanner.Scan() {
		in := scanner.Text()
		id, err := strconv.Atoi(in)
		if err != nil {
			log.Printf("Invalid Pokemon ID %q", in)
			continue
		}
		r, err := c.GetPokemon(context.Background(), &pbf.Pokemon{Number: int32(id)})
		if err != nil {
			log.Printf("Error getting: %v", err)
			continue
		}
		msg, _ := json.Marshal(r)
		fmt.Println(string(msg))
	}
}
