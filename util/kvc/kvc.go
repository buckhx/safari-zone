package kvc

import (
	"sync"
	"time"
)

type Key interface{}
type Value interface{}

type KVC interface {
	Get(Key) Value
	Has(Key) bool
	Set(Key, Value)
	SetTTL(Key, Value, time.Duration)
}

type MemKVC struct {
	sync.RWMutex
	items map[Key]Value
}

func NewMem() KVC {
	return &MemKVC{
		items: make(map[Key]Value),
	}
}

func (c *MemKVC) Get(k Key) Value {
	c.RLock()
	defer c.RUnlock()
	return c.items[k]
}

func (c *MemKVC) Has(k Key) bool {
	c.RLock()
	defer c.RUnlock()
	_, ok := c.items[k]
	return ok
}

func (c *MemKVC) Set(k Key, v Value) {
	c.Lock()
	defer c.Unlock()
	if v == nil {
		delete(c.items, k)
	} else {
		c.items[k] = v
	}
}

func (c *MemKVC) SetTTL(k Key, v Value, ttl time.Duration) {
	go func() {
		time.Sleep(ttl)
		c.Set(k, nil)
	}()
	c.Set(k, v)
}
