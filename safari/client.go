package safari

import (
	"github.com/buckhx/safari-zone/proto/pbf"
	"google.golang.org/grpc"
)

type Client struct {
	pbf.SafariClient
	*grpc.ClientConn
	addr string
}

func Dial(addr string) (*Client, error) {
	//creds := auth.AccessCredentials(tok)
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		//grpc.WithPerRPCCredentials(creds),
	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	cli := pbf.NewSafariClient(conn)
	return &Client{
		SafariClient: cli,
		ClientConn:   conn,
		addr:         addr,
	}, nil
}
