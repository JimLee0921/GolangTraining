package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. 创建可取消的 context
	ctx, cancel := context.WithCancel(context.Background())

	// 2. 启动两个 goroutine
	go worker(ctx, "A")
	go worker(ctx, "B")

	// 3. 主线程运行一段时候
	time.Sleep(1 * time.Second)
	fmt.Printf(">>> cancel now")
	cancel() // 通知所有持有 ctx 的 goroutine 停止

	// 4. 等待子 goroutine 退出后关闭 main 主线程
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("main done")
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker", name, "stopped:", ctx.Err())
			return
		default:
			fmt.Println("worker", name, "working...")
			time.Sleep(200 * time.Millisecond)
		}
	}
}
