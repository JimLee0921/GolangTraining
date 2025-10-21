package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 新版本类型
var (
	counter atomic.Int64
	wg      sync.WaitGroup
)

func main() {
	/*
		Go 之后的版本（尤其是 Go 1.19+）在 sync/atomic 包中新增了一些类型化封装的方式
		让写法更现代、更安全。
	*/
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1) // 不再需要传指针
		}()
	}
	wg.Wait()
	fmt.Println("final counter:", counter.Load()) // 读取
}
