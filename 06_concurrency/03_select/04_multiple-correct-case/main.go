package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ { // 多跑几次看随机性
		// 每轮新建 3 个有缓冲通道，并预填数据
		ch1 := make(chan int, 1)
		ch2 := make(chan int, 1)
		ch3 := make(chan int, 1)

		ch1 <- 1
		ch2 <- 2
		ch3 <- 3

		// 此时三个 case 都就绪，select 会随机选择一个执行
		select {
		case v := <-ch1:
			fmt.Println("picked ch1:", v)
		case v := <-ch2:
			fmt.Println("picked ch2:", v)
		case v := <-ch3:
			fmt.Println("picked ch3:", v)
		}
	}
}
