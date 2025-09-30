package main

import (
	"fmt"
)

func main() {
	/*
		使用  close(ch) 关闭 channel，for range 消费数据
		和前一个版本的区别
			没有死锁/阻塞风险：之前消费者用无限循环 for { <-c }，如果没人再发送，会一直阻塞，使用 close(c) + for range，循环会自动退出
			退出机制更优雅：不需要 time.Sleep 强行等待，程序会在数据消费完后自然结束
	*/
	// 1. 创建一个无缓冲 channel
	ch := make(chan int)

	// 2. 启动一个 goroutine，负责生产数据。 把 0 ~ 9 依次送进 channel。 最后 close(c)，通知不会再有新数据了
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 3. 主 goroutine 作为消费者，从 channel ch 里不断取数据，for range ch 会一直接收，直到 channel 被关闭并且所有数据都被取完为止，一旦 channel 关闭，循环自动结束，不会死等
	for n := range ch {
		fmt.Println(n)
	}
}

/*
输出结果：
0
1
2
3
4
5
6
7
8
9
*/
