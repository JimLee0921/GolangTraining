package main

import "fmt"

func withLog(fn func(int) int) func(int) int {
	return func(x int) int {
		fmt.Println("Start:", x)
		res := fn(x)
		fmt.Println("End:", res)
		return res
	}
}

func square(n int) int {
	return n * n
}

func main() {
	/*
		withLog 接收函数 fn
		返回一个包装了日志功能的新函数
		原逻辑不变，只是被装饰
		调用 loggedSquare(5) 实际执行：打印日志 ->调用原函数 ->打印结果
	*/
	loggedSquare := withLog(square)
	loggedSquare(5)

}
