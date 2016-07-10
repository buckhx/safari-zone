package registry

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/srv"
	"github.com/gengo/grpc-gateway/runtime"
	"golang.org/x/net/context"
)

type RegistrySrv struct {
	*registry
	addr string
}

func NewService(pemfile, addr string) (srv.Service, error) {
	r, err := newreg(pemfile)
	if err != nil {
		return nil, err
	}
	return &RegistrySrv{
		registry: r,
		addr:     addr,
	}, nil
}

// Register makes a creates a new trainer in the safari
//
// Trainer name, password, age & gender are required.
// Any other supplied fields will be ignored
func (s *RegistrySrv) Register(ctx context.Context, in *pbf.Trainer) (*pbf.Response, error) {
	err := s.add(in)
	if err != nil {
		return nil, err
	}
	u, err := s.get(in.Uid)
	if err != nil {
		return nil, err
	}
	msg := fmt.Sprintf("Registered %s with uid %s", u.Name, u.Uid)
	return &pbf.Response{Msg: msg, Ok: true}, nil
}

// Get fetchs a trainer
//
// The populated fields will depend on the auth scope of the token
func (s *RegistrySrv) Get(ctx context.Context, in *pbf.Trainer) (*pbf.Trainer, error) {
	u, err := s.get(in.Uid)
	if err != nil {
		return nil, err
	}
	u.Password = ""
	return u, nil
}

// Enter authenticates a user to retrieve a an access token to authorize requests for a safari
//
// HTTPS required w/ HTTP basic access authentication via a header
// Authorization: Basic BASE64({user:pass})
func (s *RegistrySrv) Enter(ctx context.Context, in *pbf.Trainer) (*pbf.Token, error) {
	return s.authenticate(in)
}

// Certificate returns the cert used to verify token signatures
//
// The cert is in JWK form as described in https://tools.ietf.org/html/rfc7517
func (s *RegistrySrv) Certificate(ctx context.Context, in *pbf.Trainer) (*pbf.Cert, error) {
	jwk, err := s.mint.MarshalPublicJwk()
	if err != nil {
		return nil, err
	}
	return &pbf.Cert{Jwk: jwk}, nil
}

func (s *RegistrySrv) Listen() error {
	tcp, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	opts := srv.Opts{
		Auth: srv.AuthOpts{
			CertURI: "dev/reg.pem",
			UnsecuredMethods: []string{
				"/buckhx.safari.registry.Registry/Certificate",
			},
		},
	}
	rpc, err := srv.NewGRPC(opts)
	if err != nil {
		return err
	}
	pbf.RegisterRegistryServer(rpc, s)
	log.Printf("%T listening at %s", s, s.addr)
	return rpc.Serve(tcp)
}

func (s *RegistrySrv) Mux() (http.Handler, error) {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pbf.RegisterRegistryHandlerFromEndpoint(ctx, mux, s.addr, opts)
	if err != nil {
		mux = nil
	}
	return http.Handler(mux), err
}
