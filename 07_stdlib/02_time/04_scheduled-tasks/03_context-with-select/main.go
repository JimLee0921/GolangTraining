package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 真实项目里优先用 context 控制超时，可向下游透传取消信号
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	select {
	case res := <-doWork():
		fmt.Println("OK:", res)
	case <-ctx.Done():
		fmt.Println("超时/取消：", ctx.Err())
	}
}

func doWork() <-chan string {
	ch := make(chan string, 1)
	go func() {
		// 模拟耗时
		time.Sleep(5 * time.Second)
		ch <- "result"
	}()
	return ch
}
