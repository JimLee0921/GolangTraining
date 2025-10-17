package main

import "fmt"

func main() {
	ch := make(chan int, 2) // 有缓冲，容量 = 2

	fmt.Println("发送 1")
	ch <- 1 // 不阻塞，直接放进缓冲区

	fmt.Println("发送 2")
	ch <- 2 // 不阻塞，缓冲区还有空间

	// fmt.Println("发送 3")
	//ch <- 3 // 阻塞：缓冲区满了，必须等有人接收

	fmt.Println("接收:", <-ch)
	fmt.Println("接收:", <-ch)
}
