package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		因为 ch2 任务耗时短，先输出 msg2
	*/
	// 创建两个管道
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 模拟两个异步任务
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "data from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "data from ch2"
	}()

	// 使用 select 等待两个通道的数据
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("receive:", msg1)
		case msg2 := <-ch2:
			fmt.Println("receive:", msg2)
		}
	}
}
