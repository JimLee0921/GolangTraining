package main

import "fmt"

// stage1：生成数据管道
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

// stage2：平方
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

// stage3：加一
func addOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v + 1
		}
	}()
	return out
}

func main() {
	/*
		每个函数就是一个 Filter
		每个 chan 是 Pipe
		串起来就是完整的 Pipeline
		每个阶段独立 goroutine 执行，自动并行
	*/
	src := gen(1, 2, 3, 4)
	sq := square(src)
	res := addOne(sq)

	for v := range res {
		fmt.Println(v)
	}
}
