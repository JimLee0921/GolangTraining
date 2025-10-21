package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		close(done) 表示事件完成或发出停止信号
		<-done 在等待
		一旦关闭，所有等待 <-done 的 goroutine 都会立刻收到信号
		done 一般不发送值，只靠关闭表示完成
	*/
	// 创建一个空结构体的 channel
	done := make(chan struct{})

	go func() {
		fmt.Println("working...")
		time.Sleep(2 * time.Second)

		fmt.Println("work done, send done semaphore to main goroutine")
		close(done) // 发送完成信号，关闭 channel

	}()
	// 主 goroutine 等待 done 关闭
	<-done
	fmt.Println("main goroutine get done semaphore, main exit")
}
