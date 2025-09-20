package main

import "fmt"

func main() {
	/*
		这个版本 x 是 main 内的局部变量，被闭包捕获，相对更安，不会被别的函数随意改动
	*/
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

}
