package main

import "fmt"

// main 函数表达式
func main() {
	/*
		函数表达式（function expressions / function values）
			在 Go 中，函数也是一种值（first-class value）：可以把函数赋值给变量、作为参数传递、作为返回值返回
			这个表达式指的就是：函数字面量 或 匿名函数，把它视作一个表达式来使用
		可以在需要的地方直接定义一个函数并赋予一个变量，用完就抛弃，不必每次都命名
		函数表达式的类型还是 func()
	*/
	tempFunc := func() {
		fmt.Println("Hello World")
	}
	tempFunc()
	fmt.Printf("tempFunc 类型为%T\n", tempFunc) // tempFunc 类型为func()
}
