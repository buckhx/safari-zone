package kvc_test

import (
	"testing"
	"time"

	"github.com/buckhx/safari-zone/util/kvc"
)

func TestMemGetSet(t *testing.T) {
	k, v := 25, "25"
	c := kvc.NewMem()
	c.Set(k, v)
	if c.Get(k) != v {
		t.Errorf("Bad Get")
	}
}

func TestMemCmpSet(t *testing.T) {
	k, v := 25, "25"
	c := kvc.NewMem()
	c.Set(k, v)
	ok := c.CompareAndSet(k, "50", func() bool {
		return !c.(*kvc.MemKVC).UnsafeHas(k)
	})
	if ok || c.Get(k) != v {
		t.Errorf("Bad CmpSet: Get(%v) want: %v got: %v", k, "50", c.Get(k))
	}
}

func TestMemTTL(t *testing.T) {
	k, v := 25, "25"
	c := kvc.NewMem()
	c.SetTTL(k, v, 1*time.Second)
	if c.Get(k) != v {
		t.Errorf("Bad Set")
	}
	time.Sleep(1 * time.Second)
	if c.Has(k) {
		t.Errorf("Bad TTL")
	}
}
