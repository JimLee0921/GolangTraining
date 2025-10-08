package main

import "fmt"

func main() {
	c := factorial(10)
	fmt.Println(<-c)
}

func factorial(n int) <-chan int {
	// 异步计算阶乘，把结果通过 channel 返回
	out := make(chan int, 1) // 缓冲1：只发一次更合适
	go func() {
		defer close(out)
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
	}()
	return out
}
