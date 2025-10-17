package main

import "fmt"

func main() {
	/*
		指针的零值是 nil
		可以与 nil 比较
		同类型指针之间可以比较是否相等
	*/
	var p1, p2 *int
	fmt.Println(p1 == p2) // true 都为 nil

	x := 5
	y := 5
	fmt.Println(x == y)   // true
	fmt.Println(&x == &y) //false （不同变量地址不同）
}
