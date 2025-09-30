package main

import "fmt"

// mian go 中的 channel 类型具有严格的类型区分
func main() {
	var chInt chan int
	var chString chan string
	fmt.Println(chInt, chString)
	// chInt = chString 直接编译错误
	chInt = make(chan int) // 初始化
}
