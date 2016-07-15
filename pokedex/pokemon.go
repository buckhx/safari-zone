package pokedex

import "github.com/buckhx/safari-zone/proto/pbf"

// pokelist embeds a *pbf.Pokemon_Collection to make working w/ P_C easier
type pokelist struct {
	*pbf.Pokemon_Collection
}

func newPokelist(pokes ...*pbf.Pokemon) pokelist {
	return pokelist{&pbf.Pokemon_Collection{Pokemon: pokes}}
}

func (l pokelist) Append(pokes ...*pbf.Pokemon) {
	l.Pokemon = append(l.Pokemon, pokes...)
}

func (l pokelist) Empty() bool {
	return len(l.Pokemon) == 0
}

func (l pokelist) HasNumber(num int32) bool {
	for _, poke := range l.Pokemon {
		if poke.Number == num {
			return true
		}
	}
	return false
}

func unknown(num int32) *pbf.Pokemon {
	return &pbf.Pokemon{Number: num,
		Name: "???",
		Type: []pbf.Pokemon_Type{pbf.UNKNOWN},
	}
}
