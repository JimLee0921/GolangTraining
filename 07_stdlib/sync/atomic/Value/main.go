package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Config struct {
	Version int
	Rate    int
}

func main() {
	var cfg atomic.Value

	// 第一次 Store：决定类型
	cfg.Store(Config{
		Version: 1,
		Rate:    100,
	})

	// reader
	for i := 0; i < 3; i++ {
		go func(id int) {
			for {
				c := cfg.Load().(Config) // 类型断言
				fmt.Printf("reader %d: version=%d rate=%d\n",
					id, c.Version, c.Rate)
				time.Sleep(300 * time.Millisecond)
			}
		}(i)
	}

	// writer：整体替换
	go func() {
		version := 1
		for {
			time.Sleep(1 * time.Second)
			version++

			cfg.Store(Config{
				Version: version,
				Rate:    version * 100,
			})
			fmt.Println("config updated:", version)
		}
	}()

	time.Sleep(4 * time.Second)
}
