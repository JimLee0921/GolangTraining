package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		// 插入数据
		for i := 0; i < 10; i++ {
			c <- i
		}
		// 通知接收者不会再有新的数据插入
		close(c)
	}()
	// 接收数据
	for i := range c {
		fmt.Println(i)
	}
}
