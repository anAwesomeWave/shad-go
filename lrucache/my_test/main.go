package main

import (
	"fmt"
	"gitlab.com/slon/shad-go/lrucache"
)

func main() {
	cache := lrucache.New(10)
	for i := 0; i < 1000; i++ {
		cache.Set(i, i+10)
	}
	var keys, values []int
	cache.Range(func(key, value int) bool {
		keys = append(keys, key)
		values = append(values, value)
		return true
	})
	fmt.Println("Keys:", keys)
	fmt.Println("Values:", values)
}
