package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 原子加 1
			counter.Add(1)
		}()
	}

	wg.Wait()

	// 原子读
	fmt.Println("final count:", counter.Load())
}
