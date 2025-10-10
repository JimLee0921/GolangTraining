package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Worker: Context已取消")
	case <-time.After(5 * time.Second):
		fmt.Println("Worker: 完成任务")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go worker(ctx)

	fmt.Println("Main: 等待5秒后取消Context")
	time.Sleep(3 * time.Second)
	cancel()
	fmt.Println("Main: Context已取消")
}
