package pokedex

import (
	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/srv/auth"
	"google.golang.org/grpc"
)

type Client struct {
	pbf.PokedexClient
	*grpc.ClientConn
	addr, tok string
}

func Dial(addr, tok string) (*Client, error) {
	creds := auth.AccessCredentials(tok)
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(creds),
	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	cli := pbf.NewPokedexClient(conn)
	return &Client{
		PokedexClient: cli,
		ClientConn:    conn,
		addr:          addr,
		tok:           tok,
	}, nil
}
