package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建根 context
	ctx := context.Background()

	// 使用 WithValue 传递 requestID 和 userID
	ctx = context.WithValue(ctx, "requestID", "abc-123")
	ctx = context.WithValue(ctx, "userID", 32)
	// 把 ctx 传给子 goroutine
	go doSomething(ctx)

	time.Sleep(500 * time.Millisecond)
}

func doSomething(ctx context.Context) {
	// 从 ctx 中取值（需要类型断言）
	requestID, ok := ctx.Value("requestID").(string)
	if ok {
		fmt.Printf("Request ID: %s\n", requestID)
	} else {
		fmt.Println("assert request id fail")
	}
	userID, ok := ctx.Value("userID").(int)
	if ok {
		fmt.Printf("user ID: %d\n", userID)
	} else {
		fmt.Println("assert user id fail")

	}
	// 模拟执行任务
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task finished")
}
