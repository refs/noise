package main

import (
	"github.com/refs/gohack/sync-cache/cache"
	"time"
)

func main(){
	// create a new cache
	// use it?
	c := cache.NewCache(5)
	c.Store("test", 42, time.Now().Add(12 * time.Hour))
	print(c.Load("test").V.(int))
}
