package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("producer: %d\n", i)
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for item := range ch {
		fmt.Printf("consumer: %d\n", item)
		time.Sleep(time.Millisecond * 400)
	}
}

func main() {
	/*
		通道有缓冲区，因此生产者可提前生产多个数据
		当缓冲区满时，ch <- 会阻塞，直到消费者消费
	*/
	ch := make(chan int, 2) // 缓冲区容量为 2

	go producer(ch)
	consumer(ch)
}
