package pokecache

import "time"


func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		item: make(map[string]cacheEntry),
	}

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			c.reapLoop(interval)
		}
	}()

	return c
}


func (c *Cache) reapLoop(d time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	for k, val := range c.item {
		if now.Sub(val.createdAt) > d {
			delete(c.item, k)
		}
	}
}