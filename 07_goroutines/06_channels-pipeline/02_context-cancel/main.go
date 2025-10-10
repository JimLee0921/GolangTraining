package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				// 写也要可取消
				select {
				case out <- v * v:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}

func toString(ctx context.Context, in <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				s := fmt.Sprintf("val=%d", v)
				select {
				case out <- s:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}

// main 加上上下文取消机制
func main() {
	/*
		在每个 stage 的读/写处监听 ctx.Done()
		上游或外部取消时，所有 goroutine 都能及时退出，防泄露
		生成的 nums 管道数量多但是只给 30ms 耗时，所以不能完成，但是正常退出
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()

	nums := gen(ctx, 10000)
	sq := square(ctx, nums)
	strs := toString(ctx, sq)

	for s := range strs {
		fmt.Println(s)
	}
	// 到 300ms 时，ctx 超时，全部 stage 连锁退出。
}
