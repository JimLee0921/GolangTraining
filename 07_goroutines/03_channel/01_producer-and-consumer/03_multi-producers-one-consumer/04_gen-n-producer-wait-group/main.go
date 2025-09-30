package main

import (
	"fmt"
	"sync"
)

// fanIn: 启动多个生产者，把它们的输出合并到一个 channel
func fanIn(n int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(n)

	// 启动 n 个生产者
	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				out <- id*100 + j // 每个生产者发 5 个数据
			}
		}(i)
	}

	// 关闭 out 的协程
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// 10 个生产者 → 1 个消费者
	c := fanIn(10)

	// 单消费者负责消费所有数据
	for v := range c {
		fmt.Println("接收到:", v)
	}
}
