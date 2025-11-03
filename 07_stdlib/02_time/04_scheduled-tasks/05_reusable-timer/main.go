package main

import (
	"fmt"
	"time"
)

func Worker(stop <-chan struct{}) {
	// 可复用的 timer
	timer := time.NewTimer(0)
	if !timer.Stop() {
		<-timer.C // 取出
	}
	for i := 1; i <= 10; i++ {
		fmt.Println("do work", i)
		// 下次间隔 1s
		timer.Reset(1 * time.Second)
		select {
		case <-timer.C:
		case <-stop:
			fmt.Println("worker stop")
			return
		}
	}
}

func main() {
	// 工作循环里有等待间隔，但需要可取消、可复用，避免泄漏
	stop := make(chan struct{})
	go Worker(stop)

	time.Sleep(4500 * time.Millisecond)
	close(stop)
	time.Sleep(200 * time.Millisecond)
}
