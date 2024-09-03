package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type Cache struct {
	mu      *sync.Mutex
	entries map[string]cacheEntry
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{val: val, createdAt: time.Now()}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.entries[key]
	if exists {
		return entry.val, exists
	} else {
		return nil, exists
	}
}

func (c Cache) reapLoop(ticker *time.Ticker, interval time.Duration) {
	<-ticker.C
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.entries {
		if time.Since(entry.createdAt) > interval {
			delete(c.entries, key)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{mu: &sync.Mutex{}, entries: make(map[string]cacheEntry)}
	ticker := time.NewTicker(interval)

	go cache.reapLoop(ticker, interval)

	return cache
}
