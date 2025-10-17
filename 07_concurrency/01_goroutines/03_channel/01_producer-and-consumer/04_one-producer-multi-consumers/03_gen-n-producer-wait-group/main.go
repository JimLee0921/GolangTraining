package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	numWorkers := 3
	wg.Add(numWorkers)

	// 启动 n 个消费者
	for w := 0; w < numWorkers; w++ {
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", id, job)
				time.Sleep(100 * time.Millisecond) // 模拟耗时任务

			}
			fmt.Printf("Worker %d done\n", id)
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

	wg.Wait()
	fmt.Println("All jobs processed.")
}
