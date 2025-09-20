package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	/*
		最终版本实现闭包
			x 是 wrapper 的局部变量
			按正常逻辑，函数结束后局部变量就会消失
			但闭包让匿名函数“捕获”了 x，所以 x 会随着闭包一起存活
			每次调用 increment()，都会修改捕获的那个 x，所以输出依次递增
			没有闭包时，想让多个函数共享一个变量，就得把变量放在 包级作用域（全局变量）
			有了闭包，可以把变量封装在函数里，让它只暴露给需要的函数使用，更安全、更模块化
		闭包还能生成多个相互独立的互不干扰的环境
	*/
	incrementA := wrapper()
	incrementB := wrapper()
	fmt.Println(incrementA()) // 都是独立的闭包环境
	fmt.Println(incrementB())
	fmt.Println(incrementA())
	fmt.Println(incrementB())
	fmt.Println(incrementA())
	fmt.Println(incrementA())
	fmt.Println(incrementB())
	fmt.Println(incrementB())
}
