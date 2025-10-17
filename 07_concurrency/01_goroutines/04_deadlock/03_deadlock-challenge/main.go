package main

import "fmt"

func main() {
	/*
		发送端：匿名 goroutine 会往 c 连续发送 0,1,...,9
		接收端：主 goroutine 只执行了一次 <-c，所以只会收到第一个值（通常是 0）
		发送端还想继续发送 1..9，但没人再接收了 -> goroutine 卡在 c <- i 阻塞
		程序退出时 Go 检测到 goroutine 阻塞，就报 deadlock
	*/
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	fmt.Println(<-c)
}
