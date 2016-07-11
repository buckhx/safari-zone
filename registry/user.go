package registry

import "github.com/buckhx/safari-zone/proto/pbf"

type user struct {
	*pbf.Trainer
}

func (u user) hasScope(scp ...string) (ok bool) {
	return hasScope(u.Scope, scp...)
}

func hasScope(scope []string, scp ...string) (ok bool) {
	ok = true //
	for _, rs := range scp {
		ok = false
		for _, us := range scope {
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
