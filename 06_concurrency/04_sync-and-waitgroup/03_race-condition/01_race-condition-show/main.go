package main

import (
	"fmt"
	"sync"
)

var (
	wg      sync.WaitGroup
	counter = 0
)

func main() {
	/*
		race condition（资源竞争）
		这里多个 goroutine 同时操作 counter 存在资源竞争问题，导致运行结果不确定
	*/
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // 多个 goroutine 同时操作 counter
		}()

	}
	wg.Wait()

	fmt.Println("final counter:", counter)
}
