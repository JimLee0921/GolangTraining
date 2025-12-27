package main

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	Name string
}

var (
	once   sync.Once
	config *Config
)

// 返回全局配置，保证之初始化一次
func getConfig() *Config {
	once.Do(func() {
		fmt.Println("initializing config")
		time.Sleep(time.Second) // 模拟耗时初始化
		config = &Config{Name: "prod"}
	})
	return config
}

func main() {
	var wg sync.WaitGroup

	// 并发调用 getConfig
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cfg := getConfig()
			fmt.Printf("goroutine %d got config: %s\n", id, cfg.Name)
		}(i)
	}
	wg.Wait()
}
