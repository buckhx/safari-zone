package registry

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sync"
	"time"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/srv/auth"
	"github.com/buckhx/safari-zone/util"
	"github.com/buckhx/safari-zone/util/kvc"
)

var validator = regexp.MustCompile(`^[a-zA-Z0-9+]{3,32}$`)

const (
	Issuer       = "buckhx.safari.registry"
	TokenDur     = 24 * time.Hour
	ProfScope    = "role:prof"
	ServiceScope = "role:svc"
)

type registry struct {
	sync.Mutex
	usrs kvc.KVC
	svcs kvc.KVC
	mint *auth.Mint
}

// pem is the path to .pem private key used to sign tokens
func newreg(pemfile string) (r *registry, err error) {
	k, err := ioutil.ReadFile(pemfile)
	if err != nil {
		return
	}
	m, err := auth.NewEC256Mint(Issuer, k)
	if err != nil {
		return
	}
	r = &registry{
		usrs: kvc.NewMem(),
		svcs: kvc.NewMem(),
		mint: m,
	}
	r.bootstrap()
	return
}

func (r *registry) add(u *pbf.Trainer) (err error) {
	switch {
	case !validator.MatchString(u.Name):
		err = fmt.Errorf("User name must match /%s/", validator)
	case !validator.MatchString(u.Password):
		err = fmt.Errorf("Password must match /%s/", validator)
	case u.Age < 10:
		err = fmt.Errorf("Trainer is too young!")
	case u.Age > 99:
		err = fmt.Errorf("Trainer is too old!")
	}
	if err != nil {
		return
	}
	u.Password = util.Hash(u.Password)
	u.Start = &pbf.Timestamp{Unix: time.Now().Unix()}
	if u.Pc == nil {
		u.Pc = &pbf.Pokemon_Collection{}
	}
	ok := false
	for !ok { // make sure out short UID isn't taken
		uid := util.GenUID()
		u.Uid = uid
		ok = r.usrs.CompareAndSet(uid, u, func() bool {
			return !r.usrs.(*kvc.MemKVC).UnsafeHas(uid)
		})
	}
	return
}

func (r *registry) get(uid string) (*pbf.Trainer, error) {
	v := r.usrs.Get(uid)
	if v == nil {
		return nil, fmt.Errorf("not registered")
	}
	if u, ok := v.(*pbf.Trainer); !ok {
		return nil, fmt.Errorf("db assertion")
	} else {
		return u, nil
	}
}

func (r *registry) authenticate(req *pbf.Trainer) (tok *pbf.Token, err error) {
	v, err := r.get(req.Uid)
	switch {
	case err != nil:
		break
	case v.Password != util.Hash(req.Password):
		err = fmt.Errorf("invalid login credentials")
	case !auth.Claims{Scope: v.Scope}.HasScope(req.Scope...):
		err = fmt.Errorf("invalid scope")
	}
	if err != nil {
		return
	}
	if sig, err := r.mint.IssueToken(req.Uid, TokenDur, req.Scope...); err == nil {
		tok = &pbf.Token{Access: sig, Scope: req.Scope}
	}
	return
}

func (r *registry) access(req *pbf.Token) (tok *pbf.Token, err error) {
	v, ok := r.svcs.Get(req.Key).(string)
	switch {
	case !ok:
		err = fmt.Errorf("%s invalid token key/secret", req.Key)
	case v != util.Hash(req.Secret):
		err = fmt.Errorf("%s invalid token key/secret", req.Key)
	}
	if err != nil {
		return
	}
	scope := []string{ServiceScope}
	if sig, err := r.mint.IssueToken(req.Key, TokenDur, scope...); err == nil {
		req.Access = sig
		req.Secret = ""
		req.Scope = scope
		tok = req
	}
	return
}

// bootstrap hydrates the db with default data
func (r *registry) bootstrap() {
	adds := []*pbf.Trainer{
		{
			Name:     "oak",
			Password: "sam+delia4EVER",
			Age:      52,
			Scope:    []string{ProfScope},
		}, {
			Name:     "ash",
			Password: "THEverybest",
			Age:      11,
			Pc: &pbf.Pokemon_Collection{
				Pokemon: []*pbf.Pokemon{
					{Number: 25}, //TODO fill this own},
				},
			},
		},
	}
	for _, u := range adds {
		if err := r.add(u); err != nil {
			log.Printf("Could not bootstrap %s %s", u.Name, err)
		} else {
			log.Printf("Bootstrapped %T %s %s", u, u.Name, u.Uid)
		}
	}
	svcs := []struct {
		key, secret string
	}{
		{key: "buckhx.safari.pokedex", secret: util.Hash("pokedex-secret")},
		{key: "buckhx.safari.zone", secret: util.Hash("zone-secret")},
	}
	for _, svc := range svcs {
		r.svcs.Set(svc.key, svc.secret)
	}
}
