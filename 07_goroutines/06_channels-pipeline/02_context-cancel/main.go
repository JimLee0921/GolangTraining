package main

import "fmt"

func main() {
	// pipeline 模式计算一百个数的阶乘
	res := factorialChannel(genSource())
	for i := range res {
		fmt.Println(i)
	}
}

func genSource() <-chan int {
	// 生成数据源管道
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

func factorialChannel(c <-chan int) <-chan int {
	// 计算数据源管道中每一个数的阶乘并保存到新的管道返回
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range c {
			out <- factorial(i)
		}
	}()
	return out
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
