package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu    sync.Mutex
	cond  = sync.NewCond(&mu)
	queue []int
	wg    sync.WaitGroup

	taskCount = 10 // 总任务数
	done      = false
)

// 消费者
func consumer(id int) {
	defer wg.Done()
	for {
		mu.Lock()
		for len(queue) == 0 && !done {
			fmt.Printf("consumer: %d is waiting\n", id)
			cond.Wait() // 等待任务
		}

		if len(queue) == 0 && done {
			// 所有任务都生产完且队列为空 -> 退出
			mu.Unlock()
			fmt.Printf("consumer %d: no more task\n", id)
			return
		}

		// 取出任务
		item := queue[0]
		queue = queue[1:]
		mu.Unlock()

		fmt.Printf("consumer %d is consuming: %d\n", id, item)
		time.Sleep(200 * time.Millisecond)
	}
}

// 生产者
func producer() {
	for i := 1; i <= taskCount; i++ {
		mu.Lock()
		queue = append(queue, i)
		fmt.Printf("producer is producing: %d\n", i)
		mu.Unlock()
		cond.Signal() // 唤醒一个消费者（随机唤醒）
		time.Sleep(100 * time.Millisecond)
	}

	// 所有任务生产完毕
	mu.Lock()
	done = true
	mu.Unlock()
	cond.Broadcast() // 唤醒所有等待的消费者，让他们检查 done 状态进行优雅退出
}

func main() {
	consumerCount := 3
	wg.Add(consumerCount)

	for i := 1; i <= consumerCount; i++ {
		go consumer(i)
	}

	producer()

	wg.Wait()
	fmt.Println("all task done")
}
