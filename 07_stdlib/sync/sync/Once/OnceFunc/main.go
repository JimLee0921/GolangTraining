package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startWorker := sync.OnceFunc(func() {
		fmt.Println("start background worker")
		go func() {
			for {
				fmt.Println("working...")
				time.Sleep(time.Second)
			}
		}()
	})

	// 多处，并发调用，但是后台 worker 只会启动一次
	for i := 0; i < 3; i++ {
		go startWorker()
	}

	time.Sleep(3 * time.Second)
}
