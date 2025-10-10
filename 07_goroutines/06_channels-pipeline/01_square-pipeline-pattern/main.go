package main

import (
	"fmt"
)

// stage: 生成 0..n-1
func gen(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- i
		}
	}()
	return out
}

// stage: 计算平方
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()
	return out
}

// stage: 转字符串
func toString(in <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for v := range in {
			out <- fmt.Sprintf("val=%d", v)
		}
	}()
	return out
}

func main() {
	nums := gen(1000)
	sq := square(nums)
	strs := toString(sq)

	for s := range strs {
		fmt.Println(s)
	}
}
