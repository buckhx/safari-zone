package registry

import (
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/buckhx/safari-zone/proto/pbf"
	"github.com/buckhx/safari-zone/util"
	"github.com/buckhx/safari-zone/util/kvc"
)

var validator = regexp.MustCompile(`^[a-zA-Z0-9]{3,32}$`)

type Registry struct {
	sync.Mutex
	db kvc.KVC
}

func NewRegistry() *Registry {
	return &Registry{
		db: kvc.NewMem(),
	}
}

func (r *Registry) Add(user *pbf.Trainer) (err error) {
	switch {
	case !validator.MatchString(user.Name):
		err = fmt.Errorf("User name must match /%s/", validator)
	case !validator.MatchString(user.Password):
		err = fmt.Errorf("Password must match /%s/", validator)
	case user.Age < 10:
		err = fmt.Errorf("Trainer is too young!")
	case user.Age > 99:
		err = fmt.Errorf("Trainer is too old!")
	}
	if err != nil {
		return
	}
	user.Password = util.Hash(user.Password)
	user.Start = &pbf.Timestamp{Unix: time.Now().Unix()}
	user.Pc = &pbf.Pokemon_Collection{}
	r.Lock() // for race w/ GenUID
	uid := util.GenUID()
	for r.db.Has(uid) {
		uid = util.GenUID()
	}
	user.Uid = uid
	r.db.Set(uid, user)
	defer r.Unlock()
	return
}

func (r *Registry) Get(uid string) *pbf.Trainer {
	return r.db.Get(uid).(*pbf.Trainer) // should check err
}

func (r *Registry) Authenticate(user *pbf.Trainer) *pbf.Token {
	o := r.Get(user.Uid)
	if o == nil || o.Password == util.Hash(user.Password) {
		return nil
	}

	return nil
}
