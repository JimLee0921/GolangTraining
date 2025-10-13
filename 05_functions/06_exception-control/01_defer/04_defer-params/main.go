package main

import "fmt"

func main() {
	/*
		defer 的 参数在注册时求值，不是在执行时
		当执行 defer fmt.Println("x =", x) 时，x 的值就被拷贝成 10
		后续 x = 20 不会影响打印
	*/
	x := 10
	defer fmt.Println("defer one x =", x)
	x = 20
	defer fmt.Println("defer two x =", x)
	fmt.Println("main x=", x)
	/*
		main x= 20
		defer two x = 20
		defer one x = 10
	*/
}
