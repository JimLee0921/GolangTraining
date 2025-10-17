package main

import (
	"fmt"
	"time"
)

// worker 工作函数
func worker(id int, in <-chan int, out chan<- int) {
	for v := range in {
		time.Sleep(100 * time.Millisecond) // 模拟任务耗时
		out <- v * v
	}
}

// main fan-out 多个工作者，已知数量任务数
func main() {
	/*
		关闭 out 的小技巧：单独起个收尾协程
		思路：当 in 关闭且所有 worker 都跑到结尾，out 也可以关
		这里先用已知任务数的方式，简单一点
		多个 goroutine 共同从 同一个 in 读，就是 fan-out 的本质
		结果无序，谁先算完谁先写
		这里先用“已知数量 jobs 来收尾
		真正使用还是需要用 WaitGroup
	*/
	in := make(chan int)
	out := make(chan int)

	//开启指定数量的 worker 从同一个 in 管道读取数据执行任务
	workers := 3
	for w := 1; w <= workers; w++ {
		go worker(w, in, out)
	}

	// 生产者添加指定数量的任务
	jobs := 10
	go func() {
		defer close(in)
		for j := 1; j <= jobs; j++ {
			in <- j
		}
	}()

	// 消费者（读取已知数量后主动结束并关闭 out）
	go func() {
		for i := 1; i <= jobs; i++ {
			fmt.Println("result:", <-out)
		}
		/*
			只有写入端才应该关 channel，这是 Go 的惯例
			这里 out 是消费者在读，它本身不是 out 的生产者
			如果在消费者里 defer close(out)，会导致还可能有其他 worker 在往 out 里写，从而 panic：send on closed channel
			所以 close 一定要在确认所有写操作都结束之后再执行
			这里它不是作为消费者，而是作为收尾者在关 channel
		*/
		close(out)
	}()
	time.Sleep(2 * time.Second) // 这里手动等待两秒等待全部 goroutine 完成
}
