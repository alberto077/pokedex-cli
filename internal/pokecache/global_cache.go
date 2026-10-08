package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}
type Cache struct {
	CacheValues map[string]CacheEntry
	Mu          sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.CacheValues[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) ThreadSafeGet(key string) ([]byte, bool) {

	c.Mu.Lock()
	defer c.Mu.Unlock()
	val, ok := c.CacheValues[key]

	if !ok {
		return []byte{}, ok
	}

	return val.Val, ok

}

func NewCache() *Cache {
	c := Cache{
		CacheValues: make(map[string]CacheEntry),
		Mu:          sync.Mutex{},
	}

	return &c
}
