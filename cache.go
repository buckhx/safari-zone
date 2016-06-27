package pokedex

import (
	"fmt"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/golang/protobuf/proto"
)

var (
	CachePath = ".pokedex.cache"
	HttpBkt   = []byte("http")
	GrpcBkt   = []byte("grpc")
)

type Cache struct {
	db *bolt.DB
}

func NewCache() (c *Cache) {
	db, err := bolt.Open(CachePath, 0600, nil)
	if err != nil {
		panic(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(HttpBkt)
		tx.CreateBucketIfNotExists(GrpcBkt)
		return nil
	})
	return &Cache{
		db: db,
	}
}

func (c *Cache) Get(bkt, key []byte) (val []byte) {
	c.db.View(func(tx *bolt.Tx) error {
		val = tx.Bucket(bkt).Get(key)
		return nil
	})
	return
}

func (c *Cache) Set(bkt, key, val []byte) {
	c.db.Update(func(tx *bolt.Tx) error {
		tx.Bucket(bkt).Put(key, val)
		return nil
	})
}

func (c *Cache) GetProto(key proto.Message) []byte {
	k, _ := proto.Marshal(key)
	return c.Get(GrpcBkt, k)
}

func (c *Cache) SetProto(key, val proto.Message) error {
	k, err := proto.Marshal(key)
	if err != nil {
		return err
	}
	v, err := proto.Marshal(val)
	if err != nil {
		return err
	}
	c.Set(GrpcBkt, k, v)
	return nil
}

// caches based on method + url. not on body
func (c *Cache) FetchRequest(r *http.Request) (res []byte, ok bool) {
	k := []byte(fmt.Sprint(r.Method, r.URL))
	res = c.Get(HttpBkt, k)
	if len(res) != 0 {
		ok = true
	}
	return
}

// caches based on method + url. not on body
func (c *Cache) SaveResponse(r *http.Request, resp []byte) {
	k := []byte(fmt.Sprint(r.Method, r.URL))
	c.Set(HttpBkt, k, resp)
}
