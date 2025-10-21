package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	// 使用 atomic.AddInt64 这种通用方法，各个版本都适合
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1) // 将 counter 变量原子+1
		}()
	}
	wg.Wait()

	fmt.Println("final counter", counter)
}
