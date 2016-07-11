package srv

import (
	"crypto"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"gopkg.in/square/go-jose.v1"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/registry/mint" //TOD rm this dep
	"github.com/dgrijalva/jwt-go"
	"github.com/mwitkow/go-grpc-middleware"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	AUTH_HEADER   = "authorization"
	BEARER_PREFIX = "Bearer "
	BASIC_PREFIX  = "Basic "
)

// AuthOpts configures a Authorizer
type AuthOpts struct {
	// UnsecuredMethods are grpc method strings that skip authorization
	UnsecuredMethods []string
	// CertURI is the uri for the publc JWK that verifies access tokens
	CertURI string
}

func (o AuthOpts) fetchCert() (pub crypto.PublicKey, err error) {
	switch {
	case strings.HasPrefix(o.CertURI, "https"):
		r, e := http.Get(o.CertURI)
		if e != nil {
			err = e
			break
		}
		if r.StatusCode != http.StatusOK {
			err = fmt.Errorf("CertURI not OK: %d", r.StatusCode)
			break
		}
		defer r.Body.Close()
		raw, e := ioutil.ReadAll(r.Body)
		if e != nil {
			err = e
			break
		}
		var cert *pbf.Cert
		if err = cert.Unmarshal(raw); err != nil {
			break
		}
		var jwk *jose.JsonWebKey
		if err = jwk.UnmarshalJSON(cert.Jwk); err != nil {
			// jwk.Valid()
			break
		}
		var ok bool
		if pub, ok = jwk.Key.(crypto.PublicKey); !ok {
			err = fmt.Errorf("JWK.Key not a crypto.PublicKey")
		}
	case strings.HasPrefix(o.CertURI, "http"):
		err = fmt.Errorf("HTTPS required for network AuthOpts.CertURI")
	case exists(o.CertURI):
		f, e := os.Open(o.CertURI)
		if e != nil {
			err = e
			break
		}
		raw, e := ioutil.ReadAll(f)
		if e != nil {
			err = e
			break
		}
		if key, e := mint.LoadECPrivateKey(raw); e == nil {
			pub = key.Public()
		} else {
			err = e
		}
	default:
		err = fmt.Errorf("AuthOpts.CertURI must be a local file or HTTPS network resource")
	}
	return
}

// Authorizer verifies authorization for a RPC calls by intercepting request metadata
type Authorizer struct {
	opts AuthOpts
	pub  crypto.PublicKey
}

func NewAuthorizer(opts AuthOpts) (*Authorizer, error) {
	pub, err := opts.fetchCert()
	if err != nil {
		return nil, err
	}
	return &Authorizer{opts: opts, pub: pub}, nil
}

func (a *Authorizer) HandleUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if a.skip(info.FullMethod) {
		return handler(ctx, req)
	}
	ctx, err := a.ValidateContext(ctx)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func (a *Authorizer) HandleStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if a.skip(info.FullMethod) {
		return handler(srv, stream)
	}
	wrap := grpc_middleware.WrapServerStream(stream)
	ctx := wrap.Context()
	ctx, err := a.ValidateContext(ctx)
	if err != nil {
		return err
	}
	wrap.WrappedContext = ctx
	return handler(srv, stream)
}

// Verify checks a token string and returns a jwt.Token if valid
func (a *Authorizer) Verify(tok string) (*jwt.Token, error) {
	if token, err := jwt.ParseWithClaims(tok, &mint.Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return a.pub, nil
	}); err == nil && token.Valid {
		return token, nil
	} else {
		return nil, err
	}
}

// ValidateContext checks for valid metadata from a context and adds CtxClaims
func (a *Authorizer) ValidateContext(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "Authorization required: no context metadata")
	}
	payload, ok := md[AUTH_HEADER]
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "Authorization required: missing header")
	}
	tok := strings.TrimPrefix(payload[0], BEARER_PREFIX)
	token, err := a.Verify(tok)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, err.Error())
	}
	ctx = context.WithValue(ctx, CtxClaims, token.Claims)
	return ctx, nil
}

// skip checks if this method is in the whitelist (skips authorization)
func (a *Authorizer) skip(method string) (ok bool) {
	for _, m := range a.opts.UnsecuredMethods {
		if m == method {
			return true
		}
	}
	return
}

type security int

const (
	public security = iota
	access
	basic
)

type creds struct {
	payload  string
	security security
}

// AccessCredentials generates grpc credentials based on the access token string
func BasicCredentials(key, secret string) credentials.PerRPCCredentials {
	payload := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s")))
	return creds{payload: payload, security: basic}
}

// AccessCredentials generates grpc credentials based on the access token string
func AccessCredentials(token string) credentials.PerRPCCredentials {
	return creds{payload: token, security: access}
}

// EmptyCredentials generates grpc credentials that can be used to call unsecured methods such as authentication
func PublicCredentials() credentials.PerRPCCredentials {
	return creds{security: public}
}

func (c creds) GetRequestMetadata(ctx context.Context, uri ...string) (md map[string]string, err error) {
	switch c.security {
	case public:
		break //TODO make sure md == nil is ok
	case access:
		md = map[string]string{AUTH_HEADER: BEARER_PREFIX + c.payload}
	case basic:
		md = map[string]string{AUTH_HEADER: BASIC_PREFIX + c.payload}
	default:
		err = fmt.Errorf("Invalid Credentials Security")
	}
	return
}
func (c creds) RequireTransportSecurity() bool {
	return false //TODO change
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
