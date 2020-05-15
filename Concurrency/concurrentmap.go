package main

import (
	"sync"
)

// ConcurrentMap is a wrapper around map that is safe for concurrent use
type ConcurrentMap struct {
	sync.RWMutex
	m map[string]int
}

// Get gets a value for a key
func (r *ConcurrentMap) Get(key string) (value int, ok bool) {
	r.RLock()
	defer r.RUnlock()
	if val, ok := r.m[key]; ok {
		return val, ok
	}
	return
}

// Set sets a key to a given value
func (r *ConcurrentMap) Set(key string, val int) {
	r.Lock()
	defer r.Unlock()
	r.m[key] = val
}

// Add increases the value under a key by n
func (r *ConcurrentMap) Add(key string, n int) int {
	r.Lock()
	defer r.Unlock()
	r.m[key] += n
	return r.m[key]
}

// Count return the map length
func (r *ConcurrentMap) Count() int {
	r.Lock()
	defer r.Unlock()
	return len(r.m)
}
