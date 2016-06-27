package pokedex

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/buckhx/pokedex/pokeapi"
	"github.com/gorilla/mux"
)

type Server struct {
	addr string
	api  *pokeapi.Client
	c    *Cache
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
		api:  pokeapi.NewClient(),
		c:    NewCache(),
	}
}

func (s *Server) Run() error {
	r := mux.NewRouter()
	r.HandleFunc("/pokemon/{id:[0-9]+}", s.handlePokemon())
	return http.ListenAndServe(s.addr, r)
}

func (s *Server) handlePokemon() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			if res, ok := s.c.fetchRequest(r); ok {
				w.Write(res)
				return
			}
		*/
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		p, err := s.api.FetchPokemon(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res, _ := json.Marshal(Pokemon{ID: p.ID, Name: p.Name})
		go s.c.SaveResponse(r, res)
		w.Write(res)
	}
}

type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
