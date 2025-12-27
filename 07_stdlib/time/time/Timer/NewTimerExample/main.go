package main

import (
	"fmt"
	"time"
)

func main() {
	// 模拟请求输入
	requestCh := make(chan int)

	// 启动一个 goroutine 模拟请求到来
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Duration(i) * time.Second)
			requestCh <- i
		}
		close(requestCh)
	}()

	// 空闲超时
	idleTimeout := 3 * time.Second

	// 创建一个一次性 Timer
	timer := time.NewTimer(idleTimeout)
	defer timer.Stop() // 兜底

	for {
		select {
		case req, ok := <-requestCh:
			if !ok {
				fmt.Println("request channel closed, exit")
				return
			}
			fmt.Println("received request:", req)

			// 重置空闲超时（Go >=1.23 可以直接 Reset）
			timer.Reset(idleTimeout)
		case <-timer.C:
			// 空闲超时触发
			fmt.Println("idle timeout reached, flush data")
			// flush 后 继续等待下一轮空闲
			timer.Reset(idleTimeout)
		}

	}
}
