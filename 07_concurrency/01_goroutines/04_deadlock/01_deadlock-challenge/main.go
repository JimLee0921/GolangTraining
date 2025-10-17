package main

import "fmt"

func main() {
	/*
		这里创建的是无缓冲 channel：c := make(chan int)
		无缓冲通道的发送必须和接收同时就绪，否则发送方会阻塞
		c <- 666 发生在主 goroutine里，此时没有任何接收者在并发地 <-c，所以主 goroutine 被卡在发送处，永远到不了下一行的接收
		结果：唯一的 goroutine 阻塞 -> 运行时报 deadlock
	*/
	c := make(chan int)
	c <- 666
	fmt.Println(<-c)
}
