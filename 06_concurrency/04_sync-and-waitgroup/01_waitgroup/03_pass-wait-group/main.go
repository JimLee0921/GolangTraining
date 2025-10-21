package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker 多个 goroutine 要共享同一个计数器，这里必须接收指针
func Worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("process:", id)
	time.Sleep(time.Second)
	fmt.Println("done:", id)
}
func main() {
	// 这里定义到 main 中，需要手动传递给外部函数才能使用
	var wg sync.WaitGroup

	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go Worker(i, &wg)
	}
	wg.Wait()
}
