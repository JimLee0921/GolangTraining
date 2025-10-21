package main

import (
	"context"
	"fmt"
	"time"
)

func producer(ctx context.Context, ch chan<- int) {
	defer close(ch)
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("producer done")
			return
		case ch <- i:
			fmt.Println("produce:", i)
			time.Sleep(time.Millisecond * 200)
		}
	}
}

func consumer(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("consumer done")
			return
		case v, ok := <-ch:
			if !ok {
				fmt.Println("channel closed, consumer exit")
				return
			}
			fmt.Println("consume:", v)
			time.Sleep(time.Millisecond * 300)
		}
	}
}

func main() {
	/*
		通过 context.WithTimeout 实现超时取消
		可扩展到带优雅关闭的并发任务系统
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan int, 5)
	go producer(ctx, ch)
	consumer(ctx, ch)
}
