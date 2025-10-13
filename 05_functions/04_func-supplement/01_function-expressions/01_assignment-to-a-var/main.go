package main

import "fmt"

// main 函数表达式赋值给一个变量
func main() {
	tempFunc := func() {
		fmt.Println("Hello World")
	}
	tempFunc()
	fmt.Printf("tempFunc 类型为%T\n", tempFunc) // tempFunc 类型为func()
}
