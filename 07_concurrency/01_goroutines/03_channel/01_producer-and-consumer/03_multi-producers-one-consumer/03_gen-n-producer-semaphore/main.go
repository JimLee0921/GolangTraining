package main

import (
	"fmt"
)

// fanIn: 启动多个生产者，把它们的输出合并到一个 channel，返回时定义为只读channel(<-chan int)
func fanIn(n int) <-chan int {
	out := make(chan int)
	done := make(chan bool)

	// 启动 n 个生产者
	for i := 0; i < n; i++ {
		go func(id int) {
			for j := 0; j < 5; j++ {
				out <- id*100 + j // 每个生产者发 5 个数据
			}
			done <- true
		}(i)
	}

	// 关闭 out 的协程
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
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
