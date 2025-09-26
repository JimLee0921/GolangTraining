package main

import "fmt"

// main	匿名自调用立即执行函数
func main() {
	func() {
		fmt.Println("i'm driving")
	}()

	// 携带参数
	func(name string) {
		fmt.Println(name)
	}("JimLee")
}
