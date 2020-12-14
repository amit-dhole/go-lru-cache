package main

import (
	"container/list"
)

//CacheElements elements of cache
type CacheElements struct {
	Key   int
	Value int
}

//LRUCache struct
type LRUCache struct {
	cap        int
	linkedList *list.List
	elements   map[int]*list.Element
}

//NewLRUCache initalize instance of LRUCache
func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		cap:        capacity,
		linkedList: list.New(),
		elements:   make(map[int]*list.Element, capacity),
	}
}

//Put to add key and value elements
func (cache *LRUCache) Put(key, val int) {
	if node, ok := cache.elements[key]; ok {
		cache.linkedList.MoveToFront(node)
		node.Value.(*list.Element).Value = CacheElements{Key: key, Value: val}
	} else {
		if cache.linkedList.Len() == cache.cap {
			node := cache.linkedList.Back().Value.(*list.Element).Value.(CacheElements).Key
			delete(cache.elements, node)
			cache.linkedList.Remove(cache.linkedList.Back())
		}
		node := &list.Element{
			Value: CacheElements{
				Key:   key,
				Value: val,
			},
		}
		element := cache.linkedList.PushFront(node)
		cache.elements[key] = element
	}
}

//Get the value from the element
func (cache *LRUCache) Get(key int) int {
	if element, ok := cache.elements[key]; ok {
		val := element.Value.(*list.Element).Value.(CacheElements).Value
		cache.linkedList.MoveToFront(element)
		return val
	}
	return -1
}
