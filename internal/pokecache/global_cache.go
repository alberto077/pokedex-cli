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

func (c *Cache) reapLoop(interval time.Duration) {
	var ticker *time.Ticker = time.NewTicker(interval)

	for t := range ticker.C {
		c.Mu.Lock()
		for req, entry := range c.CacheValues {
			if entry.CreatedAt.Add(interval).Before(t) {
				delete(c.CacheValues, req)
			}
		}
		c.Mu.Unlock()

		ticker.Reset(interval)
	}
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		CacheValues: make(map[string]CacheEntry),
		Mu:          sync.Mutex{},
	}

	go c.reapLoop(interval)

	return &c
}
