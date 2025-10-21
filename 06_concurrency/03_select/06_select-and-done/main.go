package main

import (
	"fmt"
	"time"
)

func main() {
	// done 通道是 Go 并发中最常见的停止信号手段
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("receive done sign")
				return
			default:
				fmt.Println("executing~")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(time.Second)
	done <- struct{}{} // 通知退出
	time.Sleep(500 * time.Millisecond)
}
