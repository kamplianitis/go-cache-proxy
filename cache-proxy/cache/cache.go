package cache

import (
	"sync"
)

// we will check for that -> maybe it needs to be a json
type Key string

type Value interface{}

type Cache struct {
	data map[Key]Value
	lock sync.Mutex
}
