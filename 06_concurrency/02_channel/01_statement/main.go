package main

import "fmt"

// main channel 基本声明
func main() {
	// 1. 基本声明（零值是 nil）
	var intChan chan int    // 声明一个传输 int 的 channel
	var strChan chan string // 声明一个传输 string 的 channel

	fmt.Printf("%v, %T\n", intChan, intChan) // <nil>, chan int
	fmt.Printf("%v, %T\n", strChan, strChan) // <nil>, chan string

	// 2. 使用 make 创建（最常用）
	unbufferedChan := make(chan int)          // 创建一个无缓冲区的 channel
	bufferedChan := make(chan int, 5)         // 创建一个容量为 5 的缓冲 channel
	fmt.Println(unbufferedChan, bufferedChan) // 默认打印的引用地址 0xc00001a0e0 0xc0000200a0

	// 3. channel 类型，在函数参数或变量中，可以限制 channel 的方向
	var sendOnly chan<- int
	var recvOnly <-chan int
	fmt.Println(sendOnly, recvOnly) // <nil> <nil>

}
