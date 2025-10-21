package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		sync.Once.Do(f) 不会重试
		即使 f 发生 panic，Once 也会被标记为“已执行过”，后续对 Do 的调用都会直接返回，不会再次调用 f
	*/
	var once sync.Once
	calls := 0

	f := func() {
		calls++
		fmt.Println("call #", calls)
		panic("boom")
	}

	for i := 0; i < 3; i++ {
		func() {
			defer func() { _ = recover() }()
			once.Do(f)
		}()
	}

	fmt.Println("total calls actually executed:", calls) // 输出 1

}
