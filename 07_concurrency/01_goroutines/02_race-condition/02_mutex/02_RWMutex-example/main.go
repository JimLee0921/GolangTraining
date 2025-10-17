package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	initBalance = 1000
	nReaders    = 2000
	nWriters    = 10
	readDelay   = 1 * time.Millisecond
	writeDelay  = 2 * time.Millisecond
)

// main RWMutex 读写锁
func main() {
	/*
		RWMutex 是一种读写锁，允许多个 Goroutine 同时读取，但在写入时需要独占锁
	*/
	fmt.Println("==== 基准：Mutex（读=写，完全串行化） ====")
	runWithMutex()

	fmt.Println("\n==== 基准：RWMutex（读可并发，写独占） ====")
	runWithRWMutex()

}

func runWithMutex() time.Duration {
	var (
		mu      sync.Mutex
		wg      sync.WaitGroup
		balance = initBalance
	)

	start := time.Now()

	// 启动读取（但用 Mutex，读也要独占）
	wg.Add(nReaders)
	for i := 0; i < nReaders; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			_ = balance           // 模拟读取共享数据
			time.Sleep(readDelay) // 模拟读取时间
			mu.Unlock()
		}()
	}

	// 启动写入（独占）
	wg.Add(nWriters)
	for i := 0; i < nWriters; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			time.Sleep(writeDelay) // 模拟写入时间
			balance--              // 模拟写入
			mu.Unlock()
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Mutex 总用时：%v（balance=%d）\n", elapsed, balance)
	return elapsed
}

func runWithRWMutex() time.Duration {
	var (
		rw      sync.RWMutex
		wg      sync.WaitGroup
		balance = initBalance
	)

	start := time.Now()

	// 读者用 RLock：可并发
	wg.Add(nReaders)
	for i := 0; i < nReaders; i++ {
		go func() {
			defer wg.Done()
			rw.RLock()
			_ = balance           // 模拟读取共享数据
			time.Sleep(readDelay) // 模拟读处理
			rw.RUnlock()
		}()
	}

	// 写者用 Lock：独占
	wg.Add(nWriters)
	for i := 0; i < nWriters; i++ {
		go func() {
			defer wg.Done()
			rw.Lock()
			time.Sleep(writeDelay) // 模拟写处理
			balance--              // 模拟写
			rw.Unlock()
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("RWMutex 总用时：%v（balance=%d）\n", elapsed, balance)
	return elapsed
}
