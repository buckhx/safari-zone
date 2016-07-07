package registry

import "github.com/buckhx/safari-zone/proto/pbf"

type Claims struct {
	Scopes *pbf.Token_Scopes
	jwt.StandardClaims
}
