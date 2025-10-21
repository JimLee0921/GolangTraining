package main

import (
	"context"
	"fmt"
	"time"
)

// 带有 ctx 的 filter 模板
func squareCtx(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("get context cancel")
				return
			case v, ok := <-in:
				if !ok {
					fmt.Println("in is closed")
					return
				}
				out <- v * v
			}
		}
	}()
	return out
}

func genCtx(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	return out
}

func main() {
	/*
		每个阶段 select { case <-ctx.Done(): return }
		上游和下游都能优雅退出
		避免 goroutine 泄漏
	*/
	// 给 500 毫秒超时，最终任务无法全部完成就会被 context 取消
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	src := genCtx(ctx, 1, 2, 3, 4, 5, 6)
	sq := squareCtx(ctx, src)

	for v := range sq {
		fmt.Println(v)
		time.Sleep(200 * time.Millisecond)
	}
}
