package main

import "fmt"

// main defer 关键字
func main() {
	/*
		defer 用来 延迟执行一个函数调用，直到当前函数返回之前才会执行
		defer someFunction()
			等当前函数（通常是 main 或者别的函数）快结束时，再执行 someFunction()
	*/
	fmt.Println("start")
	defer fmt.Println("deferred call")
	fmt.Println("end")
	/*
		start
		end
		deferred call
	*/
}
