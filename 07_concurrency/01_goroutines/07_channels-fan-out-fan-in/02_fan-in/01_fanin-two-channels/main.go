package main

import (
	"fmt"
	"time"
)

// main 最简单的 fanin 合并两个 channels
func main() {
	/*
		两个搬运工把各自输入转发到 out
		这里为了简单，用已知总数来收尾（不推荐用于生产）
	*/
	a := make(chan int)
	b := make(chan int)
	out := make(chan int)
	// 给定已知总数
	nums := 5

	// 生产者A
	go func() {
		defer close(a) // 谁生产谁关闭原则
		for i := 0; i < nums; i++ {
			time.Sleep(50 * time.Millisecond) // 模拟生产耗时
			a <- i
		}
	}()
	//	生产者B
	go func() {
		defer close(b) // 谁生产谁关闭原则
		for i := 100; i < 100+nums; i++ {
			time.Sleep(80 * time.Millisecond) // 模拟生产耗时
			b <- i
		}
	}()

	// 最简单的 fan-in 转发两个 goroutine
	go func() {
		for v := range a {
			out <- v
		}
	}()
	go func() {
		for v := range b {
			out <- v
		}
	}()

	// 已知总数为 10 个，读完后进行关闭
	go func() {
		for i := 0; i < nums*2; i++ {
			fmt.Println(<-out)
		}
		close(out)
	}()
	// main 等待子 goroutine 完成
	time.Sleep(1 * time.Second)
}
