package main

import "fmt"

const (
	x = iota
	y
	z
)

const (
	a = iota
	b
	c
)

// main 多个 const 声明块
func main() {
	/*
		如果有多个 const 声明块需要使用 iota
		每个声明块直接没有任何联系（每个const块中的iota值都会被重置，从0重新开始）
	*/

	fmt.Println(x) // 0
	fmt.Println(y) // 1
	fmt.Println(z) // 2
	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // 2
}
