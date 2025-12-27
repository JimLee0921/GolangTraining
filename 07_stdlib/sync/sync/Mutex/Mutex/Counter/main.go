package main

import (
	"fmt"
	"sync"
	"time"
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

// TryReset 缓存刷新，失败也不影响，适合可选性操作，不影响主流程，不要求一定成功
func (c *Counter) TryReset() bool {
	if !c.mu.TryLock() {
		// 正在被使用，返回
		return false // 说明正在被其他 goroutine 使用
	}
	defer c.mu.Unlock() // 如果没有执行到这里不会被执行
	c.n = 0
	return true
}

// 后台定时任务，不影响主流程
func resetJob(c *Counter) {
	if ok := c.TryReset(); ok {
		fmt.Println("counter reset successful")
	} else {
		fmt.Println("counter busy, skip reset")
	}
}

func main() {
	counter := &Counter{}

	// 模拟高并发写入
	for i := 0; i < 10; i++ {
		go func() {
			for {
				counter.Inc()
				fmt.Println(counter.Value())
				time.Sleep(1000 * time.Millisecond)
			}
		}()
	}

	// 后台定时 reset job
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		resetJob(counter)
	}
}
