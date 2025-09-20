package main

import "fmt"

func main() {
	x := 42 // x 函数在函数块内作用域 不能直接被其它函数访问
	fmt.Println(x)
	foo()
}

func foo() {
	// fmt.Println(x) 这里不能访问 x 会报错无法编译
	fmt.Println("there is func foo ~")
}
