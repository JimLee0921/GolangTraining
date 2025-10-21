package main

import (
	"fmt"
	"time"
)

func worker(id int, done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("worker", id, "get done semaphore")
			return
		default:
			fmt.Println("worker", id, "is working...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func main() {
	done := make(chan struct{})
	for i := 1; i <= 2; i++ {
		go worker(i, done)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("main goroutine: close done")
	close(done)
	// 手动延时等待最后的打印
	time.Sleep(300 * time.Millisecond)
}

/*
用 select 同时监听多个 channel
一旦 <-done 可读（即 channel 关闭），立即退出
这种写法可以安全地中止任何“长期循环”的 goroutine
*/
