package pokeapi

import (
	"encoding/json"
	"net/http"
)

func MockServer(addr string) error {
	http.HandleFunc("/pokemon/1", func(w http.ResponseWriter, r *http.Request) {
		p := Pokemon{ID: 1, Name: "bulbasaur"}
		o, _ := json.Marshal(p)
		w.Write(o)
	})
	return http.ListenAndServe(addr, nil)
}
