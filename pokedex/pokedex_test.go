package pokedex_test

import (
	"testing"

	"github.com/buckhx/safari-zone/pokedex"
	"github.com/buckhx/safari-zone/proto/pbf"
)

func TestFromCsv(t *testing.T) {
	tests := []*pbf.Pokemon{
		{
			Number:    574,
			Name:      "Gothita",
			Type:      []pbf.Pokemon_Type{pbf.PSYCHIC},
			CatchRate: 200,
			Speed:     45,
		},
	}
	pdx, err := pokedex.FromCsv("pokedex.csv")
	if err != nil {
		t.Error(err)
	}
	for _, test := range tests {
		pok := pdx.ByNumber(int(test.Number))
		if !pok.Equal(test) {
			t.Errorf("Pokemon does not match %s -> %s", test, pok)
		}
	}
}
