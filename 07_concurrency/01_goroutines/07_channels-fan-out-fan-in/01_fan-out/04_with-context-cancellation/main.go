package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// main 加入 context 可取消/超时（防泄漏）
func main() {
	/*
		使用 context 使得任何阻塞读写都可以被取消，避免协程泄露
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel() // 保底

	in := make(chan int)
	out := make(chan int, 32)

	var wg sync.WaitGroup

	workers := 4
	wg.Add(workers)

	for w := 1; w <= workers; w++ {
		go worker(ctx, w, in, out, &wg)
	}

	// 生产者
	go func() {
		defer close(in)
		for i := 0; i <= 1000; i++ { // 设置过量任务，超时后会自动退出
			select {
			case in <- i:
			case <-ctx.Done():
				fmt.Println("producer context cancel operation")
				return
			}
		}
	}()

	// 关闭通道
	go func() {
		wg.Wait()
		close(out)
	}()

	for r := range out {
		fmt.Println("result:", r)
	}
}

func worker(ctx context.Context, id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker", id, "context cancel operation")
			return
		case v, ok := <-in:
			if !ok {
				fmt.Println("worker", id, "have no new task")
				return
			}
			fmt.Println("worker", id, "get task", v)
			// 模拟任务耗时
			time.Sleep(100 * time.Millisecond)
			select {
			case out <- v * v:
			case <-ctx.Done():
				fmt.Println("worker", id, "context cancel operation")
				return
			}
		}
	}
}
