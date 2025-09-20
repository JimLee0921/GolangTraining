package main

import "fmt"

var x = 0 // 定义包级变量 x

func increment() int {
	// 每次被调用修改的是同一个包级变量 x
	x++
	return x
}

func main() {
	/*
		用函数访问并修改包级变量达到状态的持久化效果
		闭包常见用途：函数维护一个外部变量，从而“保存状态”
	*/
	fmt.Println(increment())
	fmt.Println(increment())
}
