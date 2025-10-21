package main

import "fmt"

func main() {
	ch := make(chan int, 2) // 有缓冲，容量 = 2

	fmt.Println("send data 1")
	ch <- 1 // 不阻塞，直接放进缓冲区

	fmt.Println("send data 2")
	ch <- 2 // 不阻塞，缓冲区还有空间

	// fmt.Println("发送 3")
	//ch <- 3 // 阻塞：若再发送一个，会阻塞（因为缓冲已满）

	fmt.Println("received:", <-ch)
	fmt.Println("received:", <-ch)
}
