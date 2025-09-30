package main

import "fmt"

func main() {
	/*
		无缓冲区 channel
	*/
	ch := make(chan int) // 无缓冲

	go func() {
		fmt.Println("准备发送 1")
		ch <- 1 // 会阻塞，直到有接收方
		fmt.Println("发送完成")
	}()

	fmt.Println("准备接收")
	val := <-ch // 接收，解除发送阻塞
	fmt.Println("接收到:", val)
}
