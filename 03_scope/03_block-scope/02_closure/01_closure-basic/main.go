package main

import "fmt"

func main() {
	x := 42 // 块级作用域中定义 x
	fmt.Println(x)
	{
		fmt.Println("内部函数可以访问外部函数的变量x: ", x) // 在闭包（函数嵌套）中外层定义的变量可以在内层访问
		y := "The credit belongs with the one who is in the ring."
		fmt.Println("内部函数自己定义的变量y: ", y)
	}
	// fmt.Println(y) // 内层函数定义的变量不能在外层访问
}
