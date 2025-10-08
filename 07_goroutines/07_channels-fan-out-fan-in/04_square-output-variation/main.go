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
	for n := range mergeAnother(c1, c2) {
		fmt.Println(n)
	}

	//for n := range merge(c1, c2) {
	//	fmt.Println(n)
	//}
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
		defer close(out)
		wg.Wait()
	}()

	return out
}

func mergeAnother(cs ...<-chan int) <-chan int {
	// 把 goroutine 体提取成 output 内联函数
	var wg sync.WaitGroup
	out := make(chan int)
	// 为 cs 中的每个输入通道启动一个输出 goroutine。
	// 输出将值从 c 复制到 out，直到 c 关闭，然后调用 wg.Done。
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}
	// 启动一个 Goroutine，在所有输出 Goroutine 完成后关闭
	// 此 Goroutine 必须在 wg.Add 调用之后启动
	go func() {
		defer close(out)
		wg.Wait()
	}()
	return out
}
