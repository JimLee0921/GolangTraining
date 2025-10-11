package main

import "fmt"

// main 函数中出现多个 defer
func main() {
	/*
		defer 语句在函数退出前执行
		多个 defer 是 后进先出（LIFO） 执行顺序
		所以先注册的最后执行
	*/
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("main running")
	/*
		main running
		defer 2
		defer 1
	*/
}
