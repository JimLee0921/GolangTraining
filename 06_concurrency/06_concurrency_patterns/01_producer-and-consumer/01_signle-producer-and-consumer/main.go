package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	defer close(ch) // 关闭管道
	for i := 1; i <= 5; i++ {
		fmt.Println("producing:", i)
		ch <- i // 发送到管道
		time.Sleep(time.Millisecond * 200)
	}
}

func consumer(ch <-chan int) {
	// 使用 for range 读取管道直到关闭
	for v := range ch {
		fmt.Println("consuming:", v)
		time.Sleep(time.Millisecond * 300)
	}
	fmt.Println("consumer done")
}

func main() {
	/*
		使用 chan<- / <-chan 明确通道方向
		close(ch) 通知消费者生产已结束
		消费者使用 for range ch 自动检测通道关闭
	*/
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
