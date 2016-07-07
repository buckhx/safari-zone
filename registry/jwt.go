package registry

type Scope string

const ()

type Scopes map[Scope]bool

type Claims struct {
	Scopes Scopes
	jwt.StandardClaims
}
