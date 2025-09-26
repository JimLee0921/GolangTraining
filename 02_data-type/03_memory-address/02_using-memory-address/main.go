package main

import "fmt"

// main 操作内存地址
func main() {
	/*
		变量可以用 & 取地址，用指针 * 操作地址
	*/
	x := 42
	p := &x // p 是 *int 类型，指向 x

	fmt.Println("x value:", x)
	fmt.Println("point p:", p)
	fmt.Println("get x value by point:", *p)

	*p = 100 // 修改指针指向的值
	fmt.Println("changed x value:", x)
	/*
		x value: 42
		point p: 0xc00000a088
		get x value by point: 42
		changed x value: 100
	*/
}
