package ancache

import (
	"container/list"
	"sync"
	"time"
)

type CacheItem struct {
	key        interface{}
	value      interface{}
	timestamp  time.Time
	expiration time.Duration
}

type LRUCache struct {
	maxSize    int
	expiryTime time.Duration
	cache      map[interface{}]*list.Element
	keys       *list.List
	mu         sync.Mutex
}

type keyValuePair struct {
	key   interface{}
	value *CacheItem
}

func Create(maxSize int, expiryTime time.Duration) *LRUCache {
	return &LRUCache{
		maxSize:    maxSize,
		expiryTime: expiryTime,
		cache:      make(map[interface{}]*list.Element),
		keys:       list.New(),
	}
}

func (c *LRUCache) Set(key interface{}, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Create a new cache item
	item := &CacheItem{
		key:        key,
		value:      value,
		timestamp:  time.Now(),
		expiration: c.expiryTime,
	}

	// Add the item to the cache and the linked list
	element := c.keys.PushBack(keyValuePair{key, item})
	c.cache[key] = element

	// If the cache exceeds the maximum size, remove the oldest item
	if c.keys.Len() > c.maxSize {
		c.removeOldest()
	}
}

func (c *LRUCache) Get(key interface{}) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	element, ok := c.cache[key]
	if !ok {
		return nil
	}

	pair := element.Value.(keyValuePair)
	item := pair.value

	// If the item has expired, remove it from the cache and return nil
	if time.Since(item.timestamp) > item.expiration {
		c.remove(key)
		return nil
	}

	// Update the item's timestamp and move it to the end of the list
	item.timestamp = time.Now()
	c.keys.MoveToBack(element)

	return item.value
}

func (c *LRUCache) removeOldest() {
	element := c.keys.Front()
	pair := element.Value.(keyValuePair)
	c.remove(pair.key)
}

func (c *LRUCache) remove(key interface{}) {
	element, ok := c.cache[key]
	if !ok {
		return
	}
	delete(c.cache, key)
	c.keys.Remove(element)
}

func (c *LRUCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache = make(map[interface{}]*list.Element)
	c.keys = list.New()
}
