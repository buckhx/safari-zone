package pokedex_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/buckhx/pokedex"
	"github.com/buckhx/pokedex/srv"

	"golang.org/x/net/context"
)

const (
	httpPort = ":8080"
	grpcPort = ":50051"
)

func setup() error {
	go pokedex.NewServer(httpPort).Run()
	go srv.New(grpcPort).Run()
	time.Sleep(1 * time.Second)
	return nil
}

func teardown() error { return nil }

func BenchmarkGrpc(b *testing.B) {
	c, err := srv.NewClient("localhost" + grpcPort)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer c.Close()
	b.RunParallel(func(pb *testing.PB) {
		var nm string
		for pb.Next() {
			r, err := c.GetPokemon(context.Background(), &srv.PokemonQuery{ID: 1})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			nm = r.Name
			fmt.Println(nm)
		}
		_ = nm
	})
}

func BenchmarkHttp(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var nm string
		for pb.Next() {
			res, err := http.Get("http://localhost" + httpPort + "/pokemon/1")
			if err != nil {
				log.Fatal(err)
			}
			b, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
			var p pokedex.Pokemon
			if err := json.Unmarshal(b, &p); err != nil {
				log.Fatal(err)
			}
			nm = p.Name
		}
		_ = nm
	})
}

func TestMain(m *testing.M) {
	setup()
	r := m.Run()
	teardown()
	os.Exit(r)
}
