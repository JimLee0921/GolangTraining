package main

import "fmt"

func main() {
	/*
		使用 recover 捕获 panic
	*/
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from:", r)
		}
	}()

	fmt.Println("before panic")
	panic("something went wrong")
	fmt.Println("after panic") // 不会执行

	/*
		before panic
		recovered from: something went wrong
	*/
}
