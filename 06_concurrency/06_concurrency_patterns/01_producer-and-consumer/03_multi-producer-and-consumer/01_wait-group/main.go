package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		val := id*10 + i
		fmt.Printf("producer %d -> %d\n", id, val)
		ch <- val
		time.Sleep(time.Millisecond * 200)
	}
}

func consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("consumer %d <- %d\n", id, v)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	/*
		使用 sync.WaitGroup 等待所有生产者结束
		通道关闭后消费者会自动退出
		可扩展为任意数量的生产者/消费者
	*/
	ch := make(chan int, 5)
	var wg sync.WaitGroup

	// 启动多个生产者
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go producer(i, ch, &wg)
	}

	// 启动多个消费者
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go consumer(i, ch, &wg)
	}

	go func() {
		wg.Wait() // 等生产者完成
		close(ch)
	}()

	time.Sleep(time.Second * 3)
}
