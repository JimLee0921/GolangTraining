package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		发送和接收必须同时准备好，否则发送方或接收方会阻塞
		这是一种 同步通信机制
		用于需要确保数据“立即被消费的场景
		这里演示给一些延时查看运行效果
	*/
	ch := make(chan int) // 无缓冲通道

	// 启动一个 goroutine 向无缓冲通道发送数据发送数据
	go func() {
		fmt.Println("prepare send data 1")
		ch <- 1 // 会阻塞，直到有接收方
		fmt.Println("send data successful")
	}()

	fmt.Println("prepare receive")
	time.Sleep(time.Second)
	val := <-ch // 此处接收前，上面的 goroutine 一直阻塞
	fmt.Println("received:", val)
	time.Sleep(time.Second)

}
