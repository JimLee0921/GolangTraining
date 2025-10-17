package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// init 不需要在 main() 里显式调用，Go 在执行 main.main() 之前，会自动调用当前包和依赖包里的所有 init() 函数
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var wg sync.WaitGroup

func main() {
	/*
		最大进程并行度
		1. 使用 runtime.GOMAXPROCS(runtime.NumCPU())
			这行的作用是告诉 Go 运行时：
				使用尽可能多的 CPU 核心来同时运行 goroutines（这里就是把逻辑 CPU 核心数设为最大值）
				默认情况下 Go 可能不会使用所有 CPU，把它调到 NumCPU() 后，可以充分并行
				GOMAXPROCS 设置后，如果机器有多核，两个 goroutine 可能真的是并行在不同核上执行，而不是单核抢时间片

		2. WaitGroup 控制并发退出
			wg.Add(2) 表示要等待两个 goroutine 完成
			foo() 和 bar() 的最后各自调用 wg.Done() 来 -1
			main() 调用 wg.Wait() 会阻塞，直到两个 goroutine 都结束才退出程序

		3. foo() 和 bar() 的区别
			foo() 每次循环都 Sleep(20ms) → 打印得比较快
			bar() 每次循环都 Sleep(50ms) → 打印得比较慢
			因为两个函数是并发执行的，输出会交错在一起，但总体耗时大约是 bar 的耗时，而不是两个时间相加
	*/
	// 添加两个 goroutine
	wg.Add(2)
	// 使用 goroutine 执行两个函数
	go Foo()
	go Bar()
	// main 等待两个 goroutine 执行完毕
	wg.Wait()

}

func Foo() {
	defer wg.Done()
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(20 * time.Millisecond) // 延时 0.02 秒
	}
}

func Bar() {
	defer wg.Done()
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(50 * time.Millisecond) // 延时 0.05 秒
	}
}
