package main

import (
	"context"
	"fmt"
	"time"
)

// 定义一个类型避免 key 冲突（推荐做法）
type contextKey string

func main() {
	// 创建根 context
	ctx := context.Background()

	// 使用 WithValue 传递 requestID
	ctx = context.WithValue(ctx, contextKey("requestID"), "abc-123")

	// 把 ctx 传给子 goroutine
	go doSomething(ctx)

	time.Sleep(500 * time.Millisecond)
}

func doSomething(ctx context.Context) {
	// 从 ctx 中取值
	if v := ctx.Value(contextKey("requestID")); v != nil {
		fmt.Println("requestID =", v)
	} else {
		fmt.Println("no requestID found")
	}

	// 模拟执行任务
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task finished")
}
