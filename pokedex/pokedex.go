package pokedex

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/buckhx/safari-zone/proto/pbf"
)

const (
	hNAME  = "pokemon"
	hNUM   = "nat"
	hT1    = "type i"
	hT2    = "type ii"
	hCATCH = "catch"
	hSPEED = "spe"
)

type Pokedex struct {
	pokemon []*pbf.Pokemon
}

func (p *Pokedex) ByNumber(num int) *pbf.Pokemon {
	// Keep this inefficient to simulate work
	for _, pok := range p.pokemon {
		if int(pok.Number) == num {
			return pok
		}
	}
	return nil
}

func FromCsv(path string) (pdx *Pokedex, err error) {
	log.Println("Importing Pokedex from", path, "...")
	defer log.Println("Done importing Pokedex from", path)
	f, err := os.Open(path)
	if err != nil {
		return
	}
	r := csv.NewReader(bufio.NewReader(f))
	h, err := r.Read()
	if err != nil {
		return
	}
	hd := make(map[string]int, len(h))
	for i, v := range h {
		v = strings.ToLower(v)
		// there are dups, so take the first one
		if _, ok := hd[v]; !ok {
			hd[v] = i
		}
	}
	pdx = &Pokedex{}
	skp := 0
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil { // bad error, toss pdx
			return nil, err
		}
		num, e1 := strconv.ParseInt(rec[hd[hNUM]], 10, 32)
		var typ []pbf.Pokemon_Type //TODO errcheck
		t1, t2 := strings.ToUpper(rec[hd[hT1]]), strings.ToUpper(rec[hd[hT2]])
		if t2 == "" {
			typ = []pbf.Pokemon_Type{pbf.Pokemon_Type(pbf.Pokemon_Type_value[t1])}
		} else {
			typ = []pbf.Pokemon_Type{pbf.Pokemon_Type(pbf.Pokemon_Type_value[t1]), pbf.Pokemon_Type(pbf.Pokemon_Type_value[t2])}
		}
		cat, e2 := strconv.ParseInt(rec[hd[hCATCH]], 10, 32)
		spd, e3 := strconv.ParseInt(rec[hd[hSPEED]], 10, 32)
		if e1 != nil || e2 != nil || e3 != nil {
			skp++
		}
		pok := &pbf.Pokemon{
			Number:    int32(num),
			Name:      rec[hd[hNAME]],
			Type:      typ,
			CatchRate: int32(cat),
			Speed:     int32(spd),
		}
		pdx.pokemon = append(pdx.pokemon, pok)
	}
	if skp != 0 {
		log.Println("Skipped", skp, "lines")
	}
	return
}
