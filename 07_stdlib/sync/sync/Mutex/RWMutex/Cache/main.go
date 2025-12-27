package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu sync.RWMutex
	m  map[string]string
}

func NewCache() *Cache {
	return &Cache{
		m: make(map[string]string),
	}
}

// Get 读操作，用 RLock/RUnlock 允许并发读
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.m[key]
	return v, ok
}

// Set 写操作，必须独占锁
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = value
}

func main() {
	cache := NewCache()
	cache.Set("lang", "go")

	var wg sync.WaitGroup

	// 10个读 goroutine 并发读
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for j := 0; j < 5; j++ {
				if v, ok := cache.Get("lang"); ok {
					fmt.Printf("[reader-%d] lang=%s\n", id, v)
				}
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	// 1 个写 goroutine 独占写
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(120 * time.Millisecond)
		cache.Set("lang", "golang")
		fmt.Println("[writer]update lang=golang")
	}()

	wg.Wait()
}
