package main

import "fmt"

func main() {
	cache := NewLRUCache(3)
	cache.Put(1, 10)
	cache.Put(2, 20)
	fmt.Println(cache.Get(1))
	cache.Put(3, 30)
	fmt.Println(cache.Get(2))
	cache.Put(4, 40)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
}
