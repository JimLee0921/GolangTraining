package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	go func() {
		fmt.Println("send 1")
		ch <- 1
		fmt.Println("send 2")
		ch <- 2
		fmt.Println("send 3 blocked")
		ch <- 3
		fmt.Println("send 3 completed")
	}()

	time.Sleep(2 * time.Second) // 暂停接收，观察发送阻塞
	fmt.Println("receive: ", <-ch)
	time.Sleep(1 * time.Second)
	fmt.Println("receive: ", <-ch)
	time.Sleep(1 * time.Second)
	fmt.Println("receive: ", <-ch)
}
