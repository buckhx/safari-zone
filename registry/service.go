package registry

import (
	"net/http"

	"github.com/buckhx/safari-zone/proto/pbf"
	"golang.org/x/net/context"
)

type RegistrySrv struct {
	*Registry
}

// Register makes a creates a new trainer in the safari
//
// Trainer name, password, age & gender are required.
// Any other supplied fields will be ignored
func (s *RegistrySrv) Register(context.Context, *pbf.Trainer) (*pbf.Response, error) { return nil, nil }

// Get fetchs a trainer
//
// The populated fields will depend on the auth scope of the token
func (s *RegistrySrv) Get(context.Context, *pbf.Trainer) (*pbf.Trainer, error) { return nil, nil }

// Enter authenticates a user to retrieve a an access token to authorize requests for a safari
//
// HTTPS required w/ HTTP basic access authentication via a header
// Authorization: Basic BASE64({user:pass})
func (s *RegistrySrv) Enter(context.Context, *pbf.Trainer) (*pbf.Token, error) { return nil, nil }

func (s *RegistrySrv) Listen() error { return nil }

func (s *RegistrySrv) Mux() (http.Handler, error) { return nil, nil }
