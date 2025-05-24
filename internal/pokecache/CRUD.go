package pokecache

import (
	"fmt"
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Printf("Cache key: %s has been added to cache\n", key)
	c.item[key] = cacheEntry{
		createdAt: 	time.Now(),
		val: 		val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Printf("Cache key: %s has been accessed though cache\n", key)
	value, found := c.item[key]
	return value.val,found
}
