package main

import "fmt"

// main 函数中出现多个 defer
func main() {
	/*
		多个 defer 会 逆序执行（像栈：后进先出 LIFO）
	*/
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
	/*
		third
		second
		first
	*/
}
