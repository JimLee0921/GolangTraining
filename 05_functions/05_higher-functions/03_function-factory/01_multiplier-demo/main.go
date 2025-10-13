package main

import "fmt"

func multiplier(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}

func main() {
	/*
		multiplier 是工厂函数
		返回的函数记住了factor
		每次生成的函数都拥有独立的状态
		double 和 triple 是两个不同的闭包实例
	*/
	double := multiplier(2)
	triple := multiplier(3)
	fmt.Println(double(5)) // 10
	fmt.Println(triple(5)) // 15
}
