package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries  map[string]CacheEntry
	Mu       sync.Mutex
	Duration time.Duration
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{Entries: make(map[string]CacheEntry), Duration: interval}
	go func(dt time.Duration) {
		ticker := time.NewTicker(dt)
		for time := range ticker.C { // for every interval pass
			cache.reapLoop(dt, time) // Loop though cache Entries to remove old Entries
		}
	}(interval)
	return &cache
}

func (c *Cache) Add(entryName string, data []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Entries[entryName] = CacheEntry{CreatedAt: time.Now(), Val: data}
}

func (c *Cache) Get(entryName string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if data, ok := c.Entries[entryName]; !ok {
		return []byte{}, ok
	} else {
		return data.Val, ok
	}
}

func (c *Cache) reapLoop(interval time.Duration, passed time.Time) {
	for entryName, cacheEntry := range c.Entries {
		c.Mu.Lock()
		Duration := passed.Sub(cacheEntry.CreatedAt) // calculate the Duration between now and when the cache entry was created
		if Duration > interval {                     // if the cache is older than the interval
			delete(c.Entries, entryName)
		}
		c.Mu.Unlock()
	}
}
