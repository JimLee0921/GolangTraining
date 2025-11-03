package main

import (
	"fmt"
	"time"
)

func heartbeat(stop <-chan struct{}) {
	ticker := time.NewTicker(1 * time.Second)
	// 一定要 defer ticker.Stop()
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("heartbeat at", t)
		case <-stop:
			fmt.Println("heartbeat stop")
			return
		}
	}
}

func main() {
	// 心跳/轮询/定时刷新，每 1s 执行一次，外部可停
	stop := make(chan struct{})
	go heartbeat(stop)

	time.Sleep(5 * time.Second)
	close(stop)
	// 等待 goroutine 退出
	time.Sleep(300 * time.Millisecond)
}
