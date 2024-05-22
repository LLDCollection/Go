package cache

import (
	"container/list"
)

type Cache interface {
	Get(key string) (string, bool)
	Put(key, value string)
}

type element struct {
	key string
	val string
}

type lruCache struct {
	mp  map[string]*list.Element
	cap int
	lst *list.List
}

// Interface implementation to prevent compile time errors
var _ Cache = &lruCache{}

// Create new LRU Cache with capacity
func NewLRUCache(capacity int) *lruCache {
	return &lruCache{
		mp: make(map[string]*list.Element),
		cap: capacity,
		lst: list.New(),
	}
}

// Get is used to obtain the value of a key stored in the LRU Cache
// If the key is present, it returns the value and updates the LRU items
// If the key is not present, it returns ("", false)
func (lru *lruCache) Get(key string) (string, bool) {
	v, exists := lru.mp[key]
	if exists {
		lru.lst.MoveToFront(v)
		return v.Value.(element).val, exists
	}
	
	return "", false
}

// Put is used to store the value against a key
// If the key exists, it will remove the old key-value pair and creates a new key-value pair
// The new value is pushed to the front of LRU
func (lru *lruCache) Put(key, value string) {
	old, exists := lru.mp[key]
	if exists {
		lru.lst.Remove(old)
		delete(lru.mp, key)
	}

	e := lru.lst.PushFront(element{key: key, val: value})

	if lru.lst.Len() > lru.cap {
		lastEle := lru.lst.Back()
		delete(lru.mp, lastEle.Value.(element).key)
		lru.lst.Remove(lastEle)
	}

	lru.mp[key] = e
}
