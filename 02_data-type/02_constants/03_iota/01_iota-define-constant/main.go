package main

import "fmt"

const (
	x = iota
	y = iota
	z = iota
)

// main iota 常量标识符
func main() {
	/*
		iota 是 Go 预定义的常量标识符
		只能在 const 常量声明块中使用
		它的值是从 0 开始，每往下写一行，自动加 1
		常用来简化 枚举值 或 位运算标志 的定义
	*/

	fmt.Println(x) // 1
	fmt.Println(y) // 2
	fmt.Println(z) // 3
}
