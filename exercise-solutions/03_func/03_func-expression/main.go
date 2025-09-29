package main

import "fmt"

func main() {
	/*
		使用函数表达式
	*/
	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
	h, even := half(5)
	fmt.Println(h, even) // 2 false
}
