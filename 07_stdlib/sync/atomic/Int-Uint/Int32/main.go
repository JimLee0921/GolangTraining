package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Worker struct {
	state atomic.Int32
}

const (
	StateInit    int32 = 0
	StateRunning int32 = 1
	StateStopped int32 = 2
)

// NewWorker 显式初始化状态
func NewWorker() *Worker {
	w := &Worker{}
	w.state.Store(StateInit)
	return w
}

// Start 只允许从 Init -> Running
func (w *Worker) Start() bool {
	swapped := w.state.CompareAndSwap(StateInit, StateRunning)
	if swapped {
		fmt.Println("worker started")
		go w.loop()
	}
	return swapped
}

// Stop 无论当前什么状态都进入 Stopped
func (w *Worker) Stop() {
	old := w.state.Swap(StateStopped)
	if old != StateStopped {
		fmt.Println("Worker stopped")
	}
}

// IsRunning 只做只读判断
func (w *Worker) IsRunning() bool {
	return w.state.Load() == StateRunning
}

// loop 模拟后台工作
func (w *Worker) loop() {
	for {
		// 原子读状态作为退出条件
		if w.state.Load() != StateRunning {
			fmt.Println("Worker loop exited")
			return
		}
		fmt.Println("Worker working...")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	worker := NewWorker()

	var wg sync.WaitGroup

	// 多个 goroutine 并发开启调用
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ok := worker.Start()
			fmt.Printf("goroutine %d Start() result = %v\n", id, ok)
		}(i)
	}
	wg.Wait()

	time.Sleep(2 * time.Second)

	fmt.Println("IsRunning:", worker.IsRunning())

	worker.Stop()

	time.Sleep(1 * time.Second)
	fmt.Println("IsRunning:", worker.IsRunning())
}
