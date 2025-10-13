package main

import "fmt"

func main() {
	fmt.Println("before panic")
	panic("something went wrong")
	fmt.Println("after panic") // 不会执行
	/*
		before panic
		panic: something went wrong
	*/
}
