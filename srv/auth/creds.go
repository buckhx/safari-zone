package auth

import (
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

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
	payload := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", key, secret)))
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

func AuthorizeContext(ctx context.Context, token string) context.Context {
	md := metadata.MD{AUTH_HEADER: []string{BEARER_PREFIX + token}}
	return metadata.NewContext(ctx, md)
}

func AuthenticateContext(ctx context.Context, key, secret string) context.Context {
	payload := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", key, secret)))
	md := metadata.MD{AUTH_HEADER: []string{BASIC_PREFIX + payload}}
	return metadata.NewContext(ctx, md)
}

func GetBasicCredentials(ctx context.Context) (key, secret string, err error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		err = fmt.Errorf("no context metadata")
		return
	}
	payload, ok := md[AUTH_HEADER]
	if !ok {
		err = fmt.Errorf("missing auth header")
		return
	}
	creds := strings.TrimPrefix(payload[0], BASIC_PREFIX)
	if payload[0] == creds {
		err = fmt.Errorf("missing basic authorization")
		return
	}
	raw, err := base64.StdEncoding.DecodeString(creds)
	kv := strings.Split(string(raw), ":")
	if err != nil || len(kv) != 2 {
		err = fmt.Errorf("invalid basic authorization payload")
		return
	}
	key = kv[0]
	secret = kv[1]
	return
}
