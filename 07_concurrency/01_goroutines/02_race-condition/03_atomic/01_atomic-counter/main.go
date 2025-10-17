package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	n atomic.Int64 // 新版本写法：直接使用结构体
}

func (c *Counter) Inc()        { c.n.Add(1) }
func (c *Counter) Add(k int64) { c.n.Add(k) }
func (c *Counter) Get() int64  { return c.n.Load() }

func main() {
	var c Counter
	var wg sync.WaitGroup

	worker := 8
	per := 10000

	wg.Add(worker)

	for i := 0; i < worker; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < per; j++ {
				c.Inc()
				fmt.Println(c.Get())
			}
		}()
	}
	wg.Wait()

	fmt.Println("total:", c.Get()) // 期望 8 * 10000 = 80000
}
