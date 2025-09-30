package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	jobs := make(chan int)

	numWorkers := 3

	// 启动 n 个消费者
	for w := 0; w < numWorkers; w++ {
		go func(id int) {
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", id, job)
				time.Sleep(100 * time.Millisecond) // 模拟耗时任务
			}
			fmt.Printf("Worker %d done\n", id)
			done <- true
		}(w)
	}

	// 生产者
	go func() {
		for j := 0; j <= 10; j++ {
			jobs <- j
			fmt.Printf("Produced job %d\n", j)
		}
		close(jobs)
	}()

	// 等待所有消费者发送完成信号
	for w := 0; w < numWorkers; w++ {
		<-done
	}

	fmt.Println("All jobs processed.")
}
