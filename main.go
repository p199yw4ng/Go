package main

import (
	"fmt"
	"github.com/p199yw4ng/ExGo/alog"
)

func main() {
	// lru

	arr := []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5}
	alog.BubbleSort(arr)
	fmt.Println(arr)
}

func LRU(){
	lruCache := alog.NewLRUCache(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	fmt.Println(lruCache.Get(1)) // 1
	lruCache.Put(3, 3)
	fmt.Println(lruCache.Get(2)) // -1
	lruCache.Put(4, 4)
	fmt.Println(lruCache.Get(1)) // -1
	fmt.Println(lruCache.Get(3)) // 3
	fmt.Println(lruCache.Get(4)) // 4
}