package main

import (
	"fmt"
	"sync"
)

// Counter 计数结构体
// Mutex 放在结构体里，而不是全局，可以保证锁与数据进行绑定，不会出现忘记哪个锁的问题，更利于维护和复用
type Counter struct {
	mu sync.Mutex // 使用互斥锁保证安全性
	n  int
}

// Inc 增加计数
func (c *Counter) Inc() {
	c.mu.Lock()         // 上锁
	defer c.mu.Unlock() // 保证解锁，推荐使用 defer
	c.n++               // 计数操作，保证临界区足够小
}

// Value 获取 c.n 的值，同样使用 Mutex 保证输入输出
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	// 启动 1000 个 goroutine 并发累加
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(counter.Value()) // 永远输出 1000 不会出现资源竞争问题
}
