package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var idGen atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 原子递增并返回新值
			id := idGen.Add(1)
			fmt.Println("generated id:", id)
		}()
	}

	wg.Wait()
}
