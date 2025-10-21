package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
	wg      sync.WaitGroup
)

func main() {
	/*
		这里启用一千个 goroutine 同时修改 counter
		如果不使用 sync.Mutex 会导致最终结果出现偏差
		正确结果应该是 1000
	*/
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()         // 上锁，独占访问
			defer mu.Unlock() // 解锁，使用 defer 关键字保证 goroutine 出错也能正常解锁
			counter++         // 临界区，修改数据
		}()
	}
	wg.Wait()
	fmt.Println("final counter:", counter)
}
