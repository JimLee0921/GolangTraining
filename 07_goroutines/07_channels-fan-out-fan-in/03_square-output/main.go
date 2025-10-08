package main

import (
	"fmt"
	"sync"
)

func main() {
	source := genSource(2, 3, 5, 6, 46, 4, 3, 3)
	// fan-out：这里是抢占式的，把同一条输入分给多个 worker
	c1 := square(source, "worker1")
	c2 := square(source, "worker2")
	// fan-in：把多个输出通道合并成一个
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func genSource(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(c <-chan int, workerName string) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range c {
			fmt.Println(workerName, "get", i)
			out <- i * i
		}
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	// 需要使用 waitGroup 确保 close(out) 的执行时机
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch <-chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
