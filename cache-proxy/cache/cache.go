package cache

import (
	"fmt"
	"sync"
)

// thread-safety on
const Threaded = true

// we will check for that -> maybe it needs to be a json
type Key string

// basically the "Any" type in GO -> This kind of approach
// allows flexibility but makes it impossible to apply type safety
type Value interface{}

type Cache struct {
	data map[Key]Value
	lock sync.RWMutex // syncronization purposes
}

func (c *Cache) SetValueToCache(key Key, value Value) {
	if Threaded {
		c.lock.Lock()
		defer c.lock.Unlock()
	}
	c.data[key] = value
}

func (c *Cache) GetCachedValue(key Key) (Value, error) {
	if Threaded {
		// since the data are not about to be edited, it is considered wiser to
		// use RLock instead of the classic Lock.
		c.lock.RLock()
		defer c.lock.RUnlock()
	}
	value, exists := c.data[key]

	if !exists {
		return nil, fmt.Errorf("Key does not exist in cache")
	}

	return value, nil
}

func (c *Cache) DropFromCache(key Key) {
	// delete a Cache key,value
	if Threaded {
		c.lock.Lock()
		defer c.lock.Unlock()
	}
	// delete is no-op in case of not
	// finding the element so no need to take cases there.
	delete(c.data, key)
}

func NewCache() *Cache {
	// no need to initialize lock -> Go does it automatically to 0
	return &Cache{
		data: make(map[Key]Value),
	}
}
