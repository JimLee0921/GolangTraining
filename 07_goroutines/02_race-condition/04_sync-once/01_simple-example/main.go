package main

import (
	"fmt"
	"sync"
)

// main 打印只执行一次
func main() {
	/*
		无论并发多少次调用 once.Do(printOnce)，都只会执行一次
	*/
	var once sync.Once

	printOnce := func() {
		fmt.Println("I run exactly once")
	}

	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			once.Do(printOnce)
		}()
	}
	wg.Wait()
}
