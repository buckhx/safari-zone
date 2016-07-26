package registry

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/srv"
	"github.com/buckhx/safari-zone/srv/auth"
	"github.com/gengo/grpc-gateway/runtime"
	"golang.org/x/net/context"
)

var unsecured = []string{
	"/buckhx.safari.registry.Registry/Certificate",
	"/buckhx.safari.registry.Registry/Register",
	"/buckhx.safari.registry.Registry/Access",
	"/buckhx.safari.registry.Registry/Enter",
}

type Service struct {
	*registry
	opts Opts
}

func NewService(opts Opts) (srv.Service, error) {
	r, err := newreg(opts.KeyPath)
	if err != nil {
		return nil, err
	}
	opts.Auth.UnsecuredMethods = unsecured
	return &Service{
		registry: r,
		opts:     opts,
	}, nil
}

func (s *Service) Name() string {
	return "registry"
}

func (s *Service) Version() string {
	return "v0"
}

// Register makes a creates a new trainer in the safari
//
// Trainer name, password, age & gender are required.
// Any other supplied fields will be ignored
func (s *Service) Register(ctx context.Context, in *pbf.Trainer) (out *pbf.Trainer, err error) {
	if err = s.add(in); err != nil {
		return
	}
	out = &(*in) //deep copy
	out.Password = ""
	return
}

// GetTrainer fetchs a trainer
//
// The populated fields will depend on the auth scope of the token
func (s *Service) GetTrainer(ctx context.Context, in *pbf.Trainer) (*pbf.Trainer, error) {
	claims, ok := auth.ClaimsFromContext(ctx)
	u, err := s.get(in.Uid)
	switch {
	case !ok:
		err = grpc.Errorf(codes.Unauthenticated, "Invalid Authorization: missing claims")
		u = nil
	case !claims.HasSubScope(in.Uid, ProfScope):
		err = grpc.Errorf(codes.PermissionDenied, "Invalid Scope")
		u = nil
	case err != nil:
		err = grpc.Errorf(codes.NotFound, "Trainer Not Found: %s", err)
	}
	return u, err
}

// UpdateTrainer updates a trainer
//
// The following fields can be updated w/ this method: Pc
func (s *Service) UpdateTrainer(ctx context.Context, in *pbf.Trainer) (*pbf.Trainer, error) {
	claims, ok := auth.ClaimsFromContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "Invalid Authorization: missing claims")
	}
	if claims.Subject != in.Uid {
		return nil, grpc.Errorf(codes.PermissionDenied, "Subject is not able to do this action")
	}
	u, ok := s.usrs.Get(in.Uid).(pbf.Trainer)
	if !ok {
		return nil, grpc.Errorf(codes.NotFound, "Trainer Not Found: %s", in.Uid)
	}
	trn := &u
	switch {
	case in.Pc != nil:
		trn.Pc = in.Pc
		fallthrough
	default:
		break
	}
	s.update(trn)
	return trn, nil
}

// Enter authenticates a user to retrieve a an access token to authorize requests for a safari
// TODO determine if the body of this method should move into the auth package
//
// HTTPS required w/ HTTP basic access authentication via a header
// Authorization: Basic BASE64({uid:pass})
func (s *Service) Enter(ctx context.Context, in *pbf.Trainer) (*pbf.Token, error) {
	key, pass, err := auth.GetBasicCredentials(ctx)
	if key != in.Uid {
		err = fmt.Errorf("user creds did not match request")
	}
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "Authorization Error: %s", err)
	}
	in.Password = pass
	tok, err := s.authenticate(in)
	if err != nil {
		err = grpc.Errorf(codes.Unauthenticated, "Authorization Error: %s", err)
	}
	return tok, err
}

// Enter authenticates a user to retrieve a an access token to authorize requests for a safari
// TODO figure out if BasicAuth should be used
//
// body must contain key & secret keys
func (s *Service) Access(ctx context.Context, tok *pbf.Token) (*pbf.Token, error) {
	access, err := s.access(tok)
	if err != nil {
		err = grpc.Errorf(codes.Unauthenticated, "Authorization Error: %s", err)
	}
	return access, err
}

// Certificate returns the cert used to verify token signatures
//
// The cert is in JWK form as described in https://tools.ietf.org/html/rfc7517
func (s *Service) Certificate(ctx context.Context, in *pbf.Trainer) (*pbf.Cert, error) {
	jwk, err := s.mint.MarshalPublicJwk()
	if err != nil {
		return nil, err
	}
	return &pbf.Cert{Jwk: jwk}, nil
}

func (s *Service) Listen() error {
	tcp, err := net.Listen("tcp", s.opts.Address)
	if err != nil {
		return err
	}
	cert, err := s.mint.MarshalPublicJwk()
	if err != nil {
		return err
	}
	s.opts.Opts.Auth.Cert = string(cert)
	rpc, err := srv.ConfigureGRPC(s.opts.Opts)
	if err != nil {
		return err
	}
	pbf.RegisterRegistryServer(rpc, s)
	log.Printf("Service %T listening at %s", s, s.opts.Address)
	return rpc.Serve(tcp)
}

func (s *Service) Mux() (http.Handler, error) {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pbf.RegisterRegistryHandlerFromEndpoint(ctx, mux, s.opts.Address, opts)
	if err != nil {
		mux = nil
	}
	return http.Handler(mux), err
}
