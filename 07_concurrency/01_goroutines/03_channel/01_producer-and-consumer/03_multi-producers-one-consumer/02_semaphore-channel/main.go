package main

import "fmt"

func main() {
	/*
		定义一个 done channel 来做生产者完成信号的同步
	*/

	ch := make(chan int)    // 主数据通道，两个生产者往里面送数据，消费者从里面读
	done := make(chan bool) // 信号通道，用来告诉“协调者”生产者已经结束

	// 两个生产者
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		done <- true // 发送完后往 done 通道里写一个信号，表示此生产者任务已完成
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		done <- true
	}()

	go func() {
		<-done // 等待两个 done 信号，说明两个生产者都已完成
		<-done
		close(ch) // 关闭 channel 告诉消费者不会再有新数据了
	}()
	// 主 goroutine 开始读取并打印 channel 中的数据
	for i := range ch {
		fmt.Println(i)
	}
}
