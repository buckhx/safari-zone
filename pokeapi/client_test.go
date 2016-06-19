package pokeapi_test

import (
	"testing"

	"github.com/buckhx/pokedex/pokeapi"
)

func TestFetchPokemon(t *testing.T) {
	tests := []struct {
		id   int
		name string
	}{
		{id: 1, name: "bulbasaur"},
	}
	c := pokeapi.NewClient()
	for _, test := range tests {
		if p, err := c.FetchPokemon(test.id); err != nil {
			t.Error(err)
		} else if p.Name != test.name {
			t.Errorf("Invalid c.FetchPokemon(%v): %v -> %v", test.id, test.name, p.Name)
		}
	}
}
