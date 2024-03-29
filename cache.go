package cache

import (
	"sync"
)

const (
	fillSize = 0.7
	minSize  = 16
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})

	Flush()
}

type cache struct {
	ring1 map[string]interface{}
	ring2 map[string]interface{}
	mt    sync.RWMutex
	limit int
	used  int
}

func New(size int) Cache {
	if size < minSize {
		size = minSize
	}

	c := &cache{
		limit: int(float64(size) * fillSize),
	}

	c.Flush()

	return c
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.mt.RLock()
	defer c.mt.RUnlock()

	if v, h := c.ring1[key]; h {
		return v, true
	}

	if v, h := c.ring2[key]; h {
		return v, true
	}

	return nil, false
}

func (c *cache) Set(key string, value interface{}) {
	c.mt.Lock()
	defer c.mt.Unlock()

	if _, h := c.ring1[key]; h {
		c.ring1[key] = value
		return
	}

	if c.used == c.limit {
		c.ring2 = c.ring1
		c.ring1 = make(map[string]interface{}, c.limit)
		c.used = 0
	}

	c.ring1[key] = value

	c.used++
}

func (c *cache) Flush() {
	c.mt.Lock()
	defer c.mt.Unlock()

	c.ring1 = make(map[string]interface{}, c.limit*2)
	c.ring2 = make(map[string]interface{})

	c.used = 0
}
