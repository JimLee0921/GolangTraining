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
	var cfg atomic.Pointer[Config]

	// 初始化，必须先 Store
	cfg.Store(&Config{
		Version: 1,
		Rate:    100,
	})

	// 多个 reader
	for i := 0; i < 3; i++ {
		go func(id int) {
			for {
				c := cfg.Load() // 原子读取指针
				fmt.Printf("reader %d : version=%d rate=%d\n", id, c.Version, c.Rate)
				time.Sleep(300 * time.Millisecond)
			}
		}(i)
	}

	// writer 定期整体替换配置
	go func() {
		version := 1
		for {
			for {
				time.Sleep(1 * time.Second)
				version++
				newCfg := &Config{
					Version: version,
					Rate:    version * 100,
				}
				cfg.Store(newCfg)
				fmt.Println("config updated:", version)
			}

		}
	}()
	time.Sleep(10 * time.Second)
}
