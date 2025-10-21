package main

import (
	"fmt"
	"sync"
	"time"
)

// Job 工作者结构体
type Job struct {
	ID   int
	Data string
}

// Result 任务结果结构体
type Result struct {
	JobId int
	Value string
}

// WorkerPool 工作池结构体，包含工作者数量，任务管道，结果管道和wg
type WorkerPool struct {
	WorkerCount int
	Jobs        chan Job
	Results     chan Result
	wg          sync.WaitGroup
}

// NewWorkerPool 生成工作池，传入生成工作者数量
func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		WorkerCount: workerCount,
		Jobs:        make(chan Job, 20),
		Results:     make(chan Result, 20),
	}
}

func (wp *WorkerPool) Start() {
	for i := 1; i <= wp.WorkerCount; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for job := range wp.Jobs {
		fmt.Printf("worker %d 处理任务 %d: %s\n", id, job.ID, job.Data)
		time.Sleep(time.Millisecond * 500)
		wp.Results <- Result{
			JobId: job.ID,
			Value: fmt.Sprintf("结果_%s", job.Data),
		}
	}
}

func (wp *WorkerPool) Stop() {
	close(wp.Jobs)
	wp.wg.Wait()
	close(wp.Results)
}

func main() {
	/*
		封装成 WorkerPool 结构
		可复用、可配置
		适合批量任务（API 调用、爬虫、导入导出）
	*/
	pool := NewWorkerPool(3)
	pool.Start()

	for i := 1; i <= 8; i++ {
		pool.Jobs <- Job{
			ID:   i,
			Data: fmt.Sprintf("task_%d", i),
		}
	}
	pool.Stop()

	for r := range pool.Results {
		fmt.Printf("task %d -> %s\n", r.JobId, r.Value)
	}
}
