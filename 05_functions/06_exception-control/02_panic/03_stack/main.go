package main

import "fmt"

func main() {
	/*
		panic 在 B 发生
		先执行 B 的 defer
		再执行 A 的 defer
		最后崩溃打印堆栈
	*/
	fmt.Println("main start")
	A()
	fmt.Println("main end") // 不会执行
	/*
		main start
		B defer
		A defer
		panic: something wrong in B
	*/
}

func A() {
	defer fmt.Println("A defer")
	B()
}

func B() {
	defer fmt.Println("B defer")
	panic("something wrong in B")
}
