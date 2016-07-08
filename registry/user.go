package registry

import "github.com/buckhx/safari-zone/proto/pbf"

type user struct {
	*pbf.Trainer
}

func (u user) hasScope(scp ...string) (ok bool) {
	ok = true //
	for _, rs := range scp {
		ok = false
		for _, us := range u.Scope {
			if us == rs {
				ok = true
				break
			}
		}
		if !ok {
			return
		}
	}
	return
}
