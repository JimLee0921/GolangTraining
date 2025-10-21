package main

import (
	"fmt"
	"sync"
	"time"
)

// worker 模拟一个工人，不断从 jobs 通道取任务执行
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("worker %d deal task %d\n", id, job)
		time.Sleep(time.Millisecond * 300) // 模拟耗时任务
		results <- job * 2
	}
}

func main() {
	/*
		任务生产者 -> jobs 通道
		多个工人从 jobs 通道取任务
		WaitGroup 等待所有工人退出
		主 goroutine 汇总结果
	*/
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	// 启动指定数量个工人
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// 生成任务
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// 等待所有工人处理完毕
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集结果
	for r := range results {
		fmt.Println("result:", r)
	}
}
