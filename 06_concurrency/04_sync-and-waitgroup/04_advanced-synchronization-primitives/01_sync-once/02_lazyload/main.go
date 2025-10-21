package main

import (
	"fmt"
	"sync"
	"time"
)

type Client struct {
	createdAt time.Time
}

var (
	once   sync.Once
	client *Client
)

// main once.Do 内的初始化只会跑一次，后续调用直接复用已建好的 client
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			c := getClient()
			fmt.Println("use client created at:", c.createdAt.Format(time.RFC3339Nano))
		}()
	}
	wg.Wait()
}

// getClient 懒加载：并发安全且只初始化一次
func getClient() *Client {
	once.Do(func() {
		// 这里可以进行各种配置等操作的初始化
		client = &Client{createdAt: time.Now()}
		fmt.Println("Client initialized")
	})
	return client
}
