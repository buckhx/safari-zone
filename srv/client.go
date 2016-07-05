package srv

import (
	"github.com/buckhx/pokedex/pbf"
	"google.golang.org/grpc"
)

type Client struct {
	pbf.PokedexClient
	conn *grpc.ClientConn
}

func NewClient(addr string) (c *Client, err error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	c = &Client{
		pbf.NewPokedexClient(conn),
		conn,
	}
	return
}

func (c *Client) Close() {
	c.conn.Close()
}
