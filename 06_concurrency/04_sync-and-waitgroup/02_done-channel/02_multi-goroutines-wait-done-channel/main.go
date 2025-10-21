package main

import (
	"fmt"
	"time"
)

func worker(id int, done <-chan struct{}) {
	fmt.Println("worker", id, "start work")
	<-done // 阻塞等待
	fmt.Println("worker", id, "end work")
}

func main() {
	done := make(chan struct{})

	for i := 1; i <= 5; i++ {
		go worker(i, done)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("main send done semaphore")
	close(done)
	// 这里等待 goroutine 最后的结束打印
	time.Sleep(500 * time.Millisecond)
}

/*
三个 worker 都在 <-done 阻塞
一旦关闭 done，所有 <-done 都立刻返回
这种广播式通知在并发控制里非常常见（一次 close，处处停止）
*/
