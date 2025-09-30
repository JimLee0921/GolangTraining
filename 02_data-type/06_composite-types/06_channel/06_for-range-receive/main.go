package main

import "fmt"

func main() {
	/*
		接收 channel 数据既可以用单次 <-ch，也可以用 for range ch。它们的使用场景有点不同
			单次接收：x := <-ch
				适合场景：只需要接收 一次或有限几次
				如果 channel 已经关闭并且没有数据了，返回的是零值
				一般搭配 value, ok := <-ch 来判断是否关闭
			循环接收：for v := range ch（必须关闭 channel否则会永久堵塞）
				适合场景：需要把 channel 里的所有数据都读完
				会一直接收，直到 channel 被关闭并且数据读完为止
				常见在 生产者关闭 channel，消费者消费所有数据 的模式里
	*/

	ch := make(chan int, 10) // 有缓冲 channel 演示

	// 发送数据到 ch
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // 手动关闭，告诉接收者没有其他数据了，否则会导致for range一直堵塞

	// 单次接收
	fmt.Println("单次接收: ", <-ch)

	// for range 循环接收剩下所有数据
	for i := range ch {
		fmt.Println("循环接受: ", i)
	}
}
