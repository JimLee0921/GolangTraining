package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 截止时间：当前时间 + 1 秒
	deadline := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// 模拟一个需要 500ms 的任务（能完成）
	err := doWork(ctx, 500*time.Millisecond)
	if err != nil {
		fmt.Println("case1:", err)
	} else {
		fmt.Println("case1: 成功")
	}

	// 模拟一个需要 2 秒的任务（会超时）
	err = doWork(ctx, 2*time.Second)
	if err != nil {
		fmt.Println("case2:", err) // context deadline exceeded
	} else {
		fmt.Println("case2: 成功")
	}
}
func doWork(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		fmt.Println("任务完成")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
