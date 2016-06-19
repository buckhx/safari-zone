package pokedex

import (
	"fmt"
	"net/http"

	"github.com/boltdb/bolt"
)

const CachePath = ".pokedex.cache"

var requestBucket = []byte("request")

type cache struct {
	db *bolt.DB
}

func newCache() (c *cache) {
	db, err := bolt.Open(CachePath, 0600, nil)
	if err != nil {
		panic(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(requestBucket)
		return nil
	})
	return &cache{
		db: db,
	}
}

func (c *cache) get(bkt, key []byte) (val []byte) {
	c.db.View(func(tx *bolt.Tx) error {
		val = tx.Bucket(bkt).Get(key)
		return nil
	})
	return
}

func (c *cache) set(bkt, key, val []byte) {
	c.db.Update(func(tx *bolt.Tx) error {
		tx.Bucket(bkt).Put(key, val)
		return nil
	})
}

// caches based on method + url. not on body
func (c *cache) fetchRequest(r *http.Request) (res []byte, ok bool) {
	k := []byte(fmt.Sprint(r.Method, r.URL))
	res = c.get(requestBucket, k)
	if len(res) != 0 {
		ok = true
	}
	return
}

// caches based on method + url. not on body
func (c *cache) saveResponse(r *http.Request, resp []byte) {
	k := []byte(fmt.Sprint(r.Method, r.URL))
	c.set(requestBucket, k, resp)
}
