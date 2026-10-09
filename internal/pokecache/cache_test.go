package pokecache

import (
	"testing"
	"time"
)

func TestAddCacheAndGet(t *testing.T) {

	cache := NewCache(5 * time.Second)
	tests := map[string]struct {
		key string
		val []byte
	}{
		"cacheAdd": {
			key: "test1",
			val: []byte("cacheresult"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cache.Add(tc.key, tc.val)

			val, ok := cache.ThreadSafeGet(tc.key)
			if !ok {
				t.Fatalf("Could not get cached value for %s", tc.key)
			}

			if string(val) != string(tc.val) {
				t.Fatalf("Cache Values dont match expected: %s got: %s", string(val), string(tc.val))
			}

		})
	}
}

func TestCacheRemoval(t *testing.T) {

	cache := NewCache(5 * time.Millisecond)
	tests := map[string]struct {
		key string
		val []byte
	}{
		"ToBeDeleted": {
			key: "test2",
			val: []byte("cacheresult"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cache.Add(tc.key, tc.val)

			// Needed to ensure that cache store exists
			_, ok := cache.ThreadSafeGet(tc.key)
			if !ok {
				t.Fatalf("Cache value not initially stored for:  %s", tc.key)
			}

			time.Sleep(20 * time.Millisecond)

			_, ok = cache.ThreadSafeGet(tc.key)
			if ok {
				t.Fatalf("Cache value expected to be deleted:  %s", tc.key)
			}

		})
	}
}
