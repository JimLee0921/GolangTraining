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
		解决方案，创建零值切片或者长度为 0、容量为 n 的切片
	*/
	//var xc []<-chan int
	xc := make([]<-chan int, 0, n)
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

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
