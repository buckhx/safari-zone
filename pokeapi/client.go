package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	BaseUrl         = "http://pokeapi.co/api/v2"
	PokemonResource = "pokemon"
)

type Client struct {
	url string
}

func NewClient() *Client {
	return &Client{url: BaseUrl}
}

func (c *Client) FetchPokemon(id int) (p *Pokemon, err error) {
	url := c.endpoint(PokemonResource, strconv.Itoa(id))
	res, err := http.Get(url)
	switch {
	case err != nil:
		break
	case res.StatusCode == http.StatusNotFound:
		err = fmt.Errorf("Unknown Pokemon #%d", id)
	case res.StatusCode != http.StatusOK:
		err = fmt.Errorf("Pokeapi not OK: %d", res.StatusCode)
	}
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &p)
	if p == (&Pokemon{}) {
		err = fmt.Errorf("Unknown Pokemon #%d", id)
	}
	return
}

func (c *Client) endpoint(resource string, args ...string) string {
	parts := append([]string{c.url, resource}, args...)
	return strings.Join(parts, "/")
}
