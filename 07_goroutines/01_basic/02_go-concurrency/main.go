package main

import (
	"fmt"
	"time"
)

// main 测试并发执行
func main() {
	/*
		在 Go 语言里，go 关键字的作用就是：把一个函数或方法放到 新的 goroutine 中去运行
		goroutine 是 Go 的并发执行单元，类似超轻量级线程（可以理解为协程）
		一个 Go 程序里可以同时运行成千上万个 goroutines，Go 的运行时调度器会自动管理它们
		与操作系统线程相比，goroutine 占用的内存极小（只有几 KB）
		语法：go someFunction()
		效果：
			会启动一个新的 goroutine 来执行 someFunction
			当前 goroutine（比如 main）不会等待，直接往下执行
			多个 goroutines 会并发运行，输出可能交错
		main() 里用 go foo() 和 go bar() 启动了两个 goroutine
			foo() 会打印 "Foo: 0" 到 "Foo: 44"
			bar() 会打印 "Bar: 0" 到 "Bar: 44"
			因为 go 关键字的存在，这两个函数会 并发执行
			输出结果会交错，不再是先全部 Foo，再全部 Bar
			每次运行的结果可能不一样，有时候先打印几行 Foo，有时候先打印几行 Bar
		问题点：
			main() 本身也是一个 goroutine
			当 main() 执行到末尾时，整个程序会直接退出，不会等子 goroutines 结束
			所以如果直接运行，有可能什么都看不到（因为 main 退出得太快）
			需要在 main 中添加足够长的延时等待 Foo 和 Bar 都执行完毕
		与协程的异同：
			相同点
				轻量级：相比操作系统线程，goroutine 占用资源更少，可以轻松创建成千上万个
				用户态调度：不是由操作系统直接调度，而是由 Go 运行时在用户态里做调度（M:N 模型）
				并发执行：多个 goroutine 可以同时运行（特别是在多核 CPU 上）
			不同点
				传统协程一般是 协作式调度：
					一个协程要主动“让出控制权”，另一个协程才能运行
					程序员需要手动管理 yield、resume 之类的操作
				而 goroutine 是抢占式调度：
					Go 运行时会在合适的点（如函数调用、IO 阻塞、系统调用等）自动切换 goroutine
					程序员不需要关心什么时候切换，写法和普通函数没区别，只要在前面加 go 就行
			goroutine = Go runtime 调度的协程。
			比“线程”轻，比“协程”自动化。
			所以有时候说 goroutine 是 增强版协程 或者 Go 内置的绿色线程
	*/
	go Foo()
	go Bar()
	time.Sleep(10 * time.Second) // 主程序等待十秒等待 goroutines 跑完
	fmt.Println("func main is down")
}
func Foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(100 * time.Millisecond) // 延时 0.1 秒
	}
}

func Bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(100 * time.Millisecond) // 延时 0.1 秒
	}
}
