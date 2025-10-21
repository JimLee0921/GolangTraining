package main

import (
	"fmt"
	"sync"
	"time"
)

// parallelSquare 同时运行多个平方计算 filter
func parallelSquare(in <-chan int, workerCount int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	// 启动多个 filter（worker）
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for v := range in {
				fmt.Printf("worker %d deal %d\n", id, v)
				time.Sleep(200 * time.Millisecond)
				out <- v * v
			}
		}(i)
	}

	// 所有 worker 完成后关闭 out
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func main() {
	src := gen(1, 2, 3, 4, 5, 6)
	sq := parallelSquare(src, 3)
	for v := range sq {
		fmt.Println("result:", v)
	}
}
