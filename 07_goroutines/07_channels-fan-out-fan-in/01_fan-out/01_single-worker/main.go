package main

import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int)

	// 单 worker: 从 in 中读取计算读取到的数据取平方后写入 out
	go func() {
		defer close(out)    // 谁生产谁关闭原则
		for v := range in { // range 直到 in 管道关闭
			out <- v * v
		}
	}()

	// 生产者
	go func() {
		defer close(in) // 谁生产谁关闭原则
		for i := 0; i <= 5; i++ {
			in <- i
		}
	}()

	// main 消费者
	for r := range out {
		fmt.Println("result:", r)
	}

}
