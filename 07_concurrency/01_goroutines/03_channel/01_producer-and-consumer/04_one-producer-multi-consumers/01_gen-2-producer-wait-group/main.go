package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		一个生产者两个消费者使用wait group进行关闭
	*/
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		// 在生产者完成后直接关闭通道
		close(ch)
	}()

	// 两个消费者读取数据
	go func() {
		defer wg.Done()
		for n := range ch {
			fmt.Println("consumer one", n)
		}
	}()

	go func() {
		defer wg.Done()
		for n := range ch {
			fmt.Println("consumer two", n)
		}
	}()

	wg.Wait() // 等待两个消费者读取完毕
}
