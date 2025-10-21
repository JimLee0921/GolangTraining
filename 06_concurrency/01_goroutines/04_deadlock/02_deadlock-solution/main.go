package main

import "fmt"

func main() {
	c := make(chan int)
	// 并发发送
	go func() {
		c <- 666
	}()
	// 主 goroutine 接收
	fmt.Println(<-c)
}
