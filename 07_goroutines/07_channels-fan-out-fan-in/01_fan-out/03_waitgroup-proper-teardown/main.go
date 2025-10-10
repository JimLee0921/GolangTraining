package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range in {
		fmt.Println("worker", id, "get task", v)
		time.Sleep(100 * time.Millisecond) // 模拟任务耗时
		out <- v * v
	}
}

func main() {
	/*
		用 WaitGroup 统计 worker 的生命周期
		谁启动 worker，谁在 wg.Wait() 后关闭 ou
	*/
	in := make(chan int)
	out := make(chan int, 16) // 设置容量适度缓冲，削峰

	var wg sync.WaitGroup

	workers := 3
	wg.Add(workers)

	for w := 1; w <= workers; w++ {
		go worker(w, in, out, &wg)
	}

	// 生产者，未知数量
	go func() {
		defer close(in)
		for i := 1; i <= 100; i++ {
			in <- i
		}
	}()

	// 统一收尾，所有 worker 结束后关闭 呕吐
	go func() {
		wg.Wait()
		close(out)
	}()

	// main goroutine 消费
	for v := range out {
		fmt.Println("result:", v)
	}
}
