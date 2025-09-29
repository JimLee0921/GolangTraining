package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	/*
		上一版中直接在 main 中使用 time.Sleep()等待 Foo 和 Bar 执行完毕并不会在开发时真正使用
		实际开发中最常用的时：用 sync.WaitGroup 来等待 goroutines 结束
		引入 sync.WaitGroup
			1. wg.Add(2) 表示要等待两个 goroutine 完成
			2. foo() 和 bar() 结束时调用 wg.Done() 各减 1
			3. main() 里调用 wg.Wait()，阻塞等待，直到计数器归零
			4. main 不会提前退出
		前面版本的问题：main 可能在 goroutines 还没打印完就直接退出或 mian 等待时间太久
		使用 WaitGroup，main 会等待两个 goroutine 完成后才退出
		输出是交错的，foo() 和 bar() 并发执行，打印的结果会交错，但两个函数都会正常执行完毕
		注意事项：
			foo 和 bar 中最后的 wg.Done() 最好写成在前面写为 defer wg.Done()
			写在最后 wg.Done() 如果函数中途 return 或发生 panic，就可能没来得及执行 wg.Done()，导致 Wait() 永远卡住

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
