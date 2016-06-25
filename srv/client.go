package srv

import "google.golang.org/grpc"

type Client struct {
	PokedexClient
	conn *grpc.ClientConn
}

func NewClient(addr string) (c *Client, err error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	c = &Client{
		NewPokedexClient(conn),
		conn,
	}
	return
}

func (c *Client) Close() {
	c.conn.Close()
}
