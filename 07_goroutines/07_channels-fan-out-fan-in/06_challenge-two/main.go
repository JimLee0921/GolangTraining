package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	// fan-out 扇出 多个函数从同一个通道读取数据，直到该通道关闭，将工作分配给所有从 in 读取数据的函数（10 个 Goroutine）
	xc := fanOut(in, 10)

	// fan-in 扇入 将多个通道复用到单个通道 将 c0 到 c9 的通道合并到单个通道
	for n := range merge(xc...) {
		fmt.Println(n)
	}

}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
	}()
	return out
}

func fanOut(in <-chan int, n int) []<-chan int {
	/*
		长度为 n的切片（里头已有 n 个 nil 通道），然后又 append 了 n 个新的通道，最终长度是 2n，且前 n 个是 nil
		在 merge 里会对每个通道使用 range 遇到 nil 通道时，for n := range c 会永远阻塞（对 nil 通道的接收/发送都会永久阻塞）
		导致 wg.Done() 永远不执行，out 也永远不关闭导致死锁
	*/
	xc := make([]<-chan int, n)
	for i := 0; i < n; i++ {
		xc = append(xc, factorial(in))
	}
	return xc
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- fact(n)
		}
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

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
	// 启动一个 Goroutine，在所有输出 Goroutine 完成后关闭，此 Goroutine 必须在 wg.Add 调用之后启动
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
