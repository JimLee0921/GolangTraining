package main

import "fmt"

func main() {
	/*
		一个生产者两个消费者使用信号进行关闭
	*/
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		// 在生产者完成后直接关闭通道
		close(ch)
	}()

	// 两个消费者读取数据
	go func() {
		for n := range ch {
			fmt.Println("consumer one", n)
		}
		done <- true
	}()

	go func() {
		for n := range ch {
			fmt.Println("consumer two", n)
		}
		done <- true
	}()

	// 等待消费者结束后再关闭 main goroutine，这里不能放在 go func() 中，否则还是会直接结束 main
	<-done
	<-done

}
