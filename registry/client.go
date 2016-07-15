package registry

import (
	"github.com/buckhx/safari-zone/proto/pbf"
	"google.golang.org/grpc"
)

type SrvClient struct {
	reg  pbf.RegistryClient
	addr string
}

func NewSrvClient(addr string) *SrvClient {
	return &SrvClient{addr: addr}
}

func (sc *SrvClient) BindRegistry() error {
	if conn, err := grpc.Dial(sc.addr, grpc.WithInsecure()); err == nil {
		sc.reg = pbf.NewRegistryClient(conn)
	} else {
		return err
	}
	return nil
}
