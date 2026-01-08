package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]CacheEntry
	mu       sync.Mutex
	duration time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{entries: make(map[string]CacheEntry), duration: interval}
	go func(dt time.Duration) {
		ticker := time.NewTicker(dt)
		for time := range ticker.C { // for every interval pass
			cache.reapLoop(dt, time) // Loop though cache entries to remove old entries
		}
	}(interval)
	return &cache
}

func (c *Cache) Add(entryName string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[entryName] = CacheEntry{createdAt: time.Now(), val: data}
}

func (c *Cache) Get(entryName string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if data, ok := c.entries[entryName]; !ok {
		return []byte{}, ok
	} else {
		return data.val, ok
	}
}

func (c *Cache) reapLoop(interval time.Duration, passed time.Time) {
	for entryName, cacheEntry := range c.entries {
		c.mu.Lock()
		duration := passed.Sub(cacheEntry.createdAt) // calculate the duration between now and when the cache entry was created
		if duration > interval {                     // if the cache is older than the interval
			delete(c.entries, entryName)
		}
		c.mu.Unlock()
	}
}
