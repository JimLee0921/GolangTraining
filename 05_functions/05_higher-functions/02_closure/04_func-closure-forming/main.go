package main

import "fmt"

func main() {
	/*
		闭包最终版本，创建一个 wrapper 函数，将内部函数作为返回值
		变量 x 放于 wrapper 函数内使得内部函数可以访问
		匿名函数捕获了外部的 x，即使 wrapper 执行完毕，x 也不会消失
		这种函数携带外部变量环境的行为就叫闭包
		使用：
			调用 wrapper()，返回那个匿名函数，赋值给 increment
			此时 increment 就是一个闭包函数，它记住了属于自己的 x
	*/
	// 使用闭包
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
