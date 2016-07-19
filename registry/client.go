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
	/*
		payload, err := json.Marshal(&pbf.Token{Key: key, Secret: secret})
		if err != nil {
			return "", err
		}
		buf := bytes.NewBuffer(payload)
		url := fmt.Sprintf("http://%s/registry/v0/access", addr) // TODO change to HTTPS
		resp, err := http.Post(url, "application//javascript", buf)
		if err != nil {
			return "", err
		}
		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("Status != OK: %d", resp.StatusCode)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		var token *pbf.Token
		if err = json.Unmarshal(body, token); err != nil {
			return "", err
		}
		return token.Access, nil
	*/
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
	/*
		if err == nil {
			s.ctx = auth.AuthorizeContext(s.ctx, tok.Access)
		} else {
			return err
		}
	*/
}
