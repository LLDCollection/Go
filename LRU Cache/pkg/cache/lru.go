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

func NewLRUCache(capacity int) *lruCache {
	return &lruCache{
		mp: make(map[string]*list.Element),
		cap: capacity,
		lst: list.New(),
	}
}

func (lru *lruCache) Get(key string) (string, bool) {
	v, exists := lru.mp[key]
	if exists {
		lru.lst.MoveToFront(v)
		return v.Value.(element).val, exists
	}
	
	return "", false
}

func (lru *lruCache) Put(key, value string) {
	e := lru.lst.PushFront(element{key: key, val: value})

	if lru.lst.Len() > lru.cap {
		lastEle := lru.lst.Back()
		lru.lst.Remove(lastEle)
	}

	lru.mp[key] = e
}
