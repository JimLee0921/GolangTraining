package main

import "fmt"

// main defer 关键字
func main() {
	/*
		defer 用来 延迟执行一个函数调用，直到当前函数返回之前才会执行
		多个 defer 会 逆序执行（像栈：后进先出 LIFO）
			即使函数中间 return 了，defer 语句依然会执行
			常用于 资源清理（关闭文件、解锁、关闭连接）
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
