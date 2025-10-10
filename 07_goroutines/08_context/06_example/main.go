package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 总体超时 2s，创建处就 defer cancel，防止资源泄露
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := doWork(ctx); err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("done")
}

// doWork 链式调用再创建一个子 context 收缩为一秒
func doWork(ctx context.Context) error {
	// 子步骤最多 1s
	subCtx, subCancel := context.WithTimeout(ctx, 1*time.Second)
	defer subCancel()

	select {
	case <-time.After(1500 * time.Millisecond): // 比 1s 慢，会触发取消
		fmt.Println("sub step finished")
		return nil
	case <-subCtx.Done():
		return subCtx.Err()
	}
}
