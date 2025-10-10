package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	/*
		1. wg.Add(2) 表示要等待两个 goroutine 完成
		2. foo() 和 bar() 结束时调用 wg.Done() 各减 1
		3. main() 里调用 wg.Wait()，阻塞等待，直到计数器归零
		4. main 不会提前退出
	*/
	wg.Add(2)
	go Foo()
	go Bar()
	wg.Wait()
}
func Foo() {
	defer wg.Done()
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(20 * time.Millisecond) // 延时 0.02 秒
	}
	//wg.Done()
}

func Bar() {
	defer wg.Done()
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(50 * time.Millisecond) // 延时 0.05 秒
	}
	//wg.Done() // 使用 defer wg.Done()
}
