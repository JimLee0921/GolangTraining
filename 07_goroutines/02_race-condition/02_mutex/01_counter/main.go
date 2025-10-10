package main

import (
	"fmt"
	"sync"
)

// Counter 是一个并发安全的计数器
type Counter struct {
	mu sync.Mutex
	n  int
}

// Inc 递增计数器
func (c *Counter) Inc() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

// Value 返回当前计数
func (c *Counter) Value() int {
	c.mu.Lock()
	v := c.n
	c.mu.Unlock()
	return v
}

func main() {
	var wg sync.WaitGroup
	c := &Counter{n: 2} // 零值就是可用状态，不需要额外初始化

	// 启动 10 个 goroutine，每个递增 1000 次
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				c.Inc()
				fmt.Println(c.Value())
			}
		}()
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	// 打印结果
	fmt.Println("Final Counter Value:", c.Value())
}
