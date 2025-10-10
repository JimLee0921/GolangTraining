package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 设置 1s 延迟
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// 模拟一个需要两秒耗时的任务
	err := doWork(ctx, 2*time.Second)

	if err != nil {
		fmt.Println("main task fail", err)
	} else {
		fmt.Println("main task successful")
	}
}

func doWork(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		fmt.Println("work task successful")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
