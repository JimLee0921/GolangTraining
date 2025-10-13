package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SafeGo 启动一个 goroutine，并在内部捕获 panic，防止程序崩溃。
func safeGo(wg *sync.WaitGroup, name string, fn func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("[RECOVER] goroutine %s recovered from panic: %v\n", name, r)
			}
		}()

		// 模拟执行任务
		fmt.Printf("[%s] start working\n", name)
		fn()
		fmt.Printf("[%s] finished\n", name)
	}()
}

// workerTask 模拟任务，随机 panic
func workerTask(id int) func() {
	return func() {
		time.Sleep(time.Duration(rand.Intn(500)))
		if rand.Float32() < 0.3 { // 30% 概率 panic
			panic(fmt.Sprintf("worker-%d unexpected error!", id))
		}
		fmt.Printf("worker-%d done normally\n", id)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 100; i++ {
		safeGo(&wg, fmt.Sprintf("W%d", i), workerTask(i))
	}
	wg.Wait()

}
