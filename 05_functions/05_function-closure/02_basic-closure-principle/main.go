package main

import "fmt"

var x int

func increment() int {
	x++
	return x
}

// main 包作用域变量
func main() {
	/*
		由于块作用域的原因，两个或多个函数要访问同一个变量，该变量就需要属于包作用域（这里的 x）
		而闭包就是为了解决这个问题
	*/
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2

}
