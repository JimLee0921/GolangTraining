package main

import "fmt"

func main() {
	/*
		立即调用函数表达式 IIFE = Immediately Invoked Function Expression
		也就是自执行函数
	*/
	func() {
		fmt.Println("Initializing...")
	}()
}
