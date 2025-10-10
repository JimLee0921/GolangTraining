package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Config 定义配置结构体：版本号+限额
type Config struct {
	Version string
	Limit   int
}

// atomic.Pointer[T] 是 Go 1.19+ 引入的泛型类型，封装了对 *T 的原子读写，保证并发安全
var cfg atomic.Pointer[Config]

// GetConfig Load()：原子读指针，返回当前的配置快照
func GetConfig() *Config { return cfg.Load() }

// SetConfig Store()：原子写指针，一次性发布新配置
func SetConfig(c *Config) { cfg.Store(c) }

func main() {
	/*
		每次更新创建新实例再 Store；把配置视为不可变对象
		读侧 Load() 得到的是一致快照，不需要锁
	*/
	// 首次发布
	SetConfig(&Config{Version: "v1", Limit: 100})

	/*
		启动 4 个 goroutine 并发读配置
		每次读到的 *Config 是一个 快照指针，只读，不会被改
		因为 atomic.Pointer 是无锁的，所以这里并发安全
	*/
	var wg sync.WaitGroup
	readers := 4
	wg.Add(readers)
	for i := 0; i < readers; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c := GetConfig()
				fmt.Printf("[reader %d] version=%s limit=%d\n", id, c.Version, c.Limit)
				time.Sleep(80 * time.Millisecond)
			}
		}(i)
	}
	/*
		等 200ms，把配置更新为 v2
		再等 250ms，更新为 v3
		更新是一次性切换指针，不会影响已经拿到旧指针的读者
		老读者还在用旧快照，新读者立即能读到新配置
	*/
	go func() {
		time.Sleep(200 * time.Millisecond)
		SetConfig(&Config{Version: "v2", Limit: 250})
		time.Sleep(250 * time.Millisecond)
		SetConfig(&Config{Version: "v3", Limit: 400})
	}()

	wg.Wait()
}
