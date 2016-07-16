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
	GetAndSet(Key, func(Value) Value)
	CompareAndSet(Key, Value, func() bool) bool
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
	return c.UnsafeGet(k)
}

func (c *MemKVC) Has(k Key) bool {
	c.RLock()
	defer c.RUnlock()
	return c.UnsafeHas(k)
}

func (c *MemKVC) Set(k Key, v Value) {
	c.Lock()
	defer c.Unlock()
	c.UnsafeSet(k, v)
}

func (c *MemKVC) SetTTL(k Key, v Value, ttl time.Duration) {
	go func() {
		time.Sleep(ttl)
		c.Set(k, nil)
	}()
	c.Set(k, v)
}

// CompareAndSet sets the value if the cmp function returns true
// Only Unsafe* method's should be used in the cmp func since they do not acquire locks
func (c *MemKVC) CompareAndSet(k Key, v Value, cmp func() bool) bool {
	c.Lock()
	defer c.Unlock()
	ok := cmp()
	if ok {
		c.UnsafeSet(k, v)
	}
	return ok
}

func (c *MemKVC) GetAndSet(k Key, fn func(cur Value) Value) {
	c.Lock()
	defer c.Unlock()
	cur := c.UnsafeGet(k)
	v := fn(cur)
	c.UnsafeSet(k, v)
}

func (m *MemKVC) UnsafeGet(k Key) Value {
	return m.items[k]
}

func (m *MemKVC) UnsafeHas(k Key) bool {
	_, ok := m.items[k]
	return ok
}

func (m *MemKVC) UnsafeSet(k Key, v Value) {
	if v == nil {
		delete(m.items, k)
	} else {
		m.items[k] = v
	}
}
