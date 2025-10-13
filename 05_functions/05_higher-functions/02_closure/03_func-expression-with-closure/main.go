package main

import "fmt"

func main() {
	/*
		Go 的闭包 (closure)、匿名函数 和 函数表达式
		闭包是函数 + 外部变量环境的组合
	*/
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}
