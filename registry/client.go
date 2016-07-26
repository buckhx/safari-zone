package registry

import (
	"github.com/buckhx/safari-zone/proto/pbf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client struct {
	pbf.RegistryClient
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
	cli := pbf.NewRegistryClient(conn)
	return &Client{
		RegistryClient: cli,
		ClientConn:     conn,
		addr:           addr,
		//tok:            tok,
	}, nil
}

func FetchAccessToken(addr, key, secret string) (string, error) {
	reg, err := Dial(addr)
	if err != nil {
		return "", err
	}
	defer reg.Close()
	tok, err := reg.Access(context.Background(), &pbf.Token{Key: key, Secret: secret})
	if err != nil {
		return "", err
	}
	return tok.Access, nil
}
