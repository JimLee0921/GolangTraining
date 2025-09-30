package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var balance = 1000 // 银行账户余额

func main() {
	/*
		race condition（资源竞争）：资源竞争指的是：
			在并发程序里，多个 goroutine / 线程 同时访问和修改同一个资源（变量、文件、网络连接等），而访问顺序和执行时间不确定，最终结果也就变得不可预测
			通俗点说：大家同时去抢一个东西，但没有人管秩序，谁先谁后不确定，最后就乱套了
		解决方案主要有三种：
			1. 用互斥锁 (sync.Mutex)，适合逻辑较复杂情况
			2. 用原子操作 (sync/atomic)，适合逻辑较简单情况
			3. 用 channel 传递数据（Go 推崇的方式，避免多个 goroutine 共享可变状态）
		可以使用 go run 命令查看是否存在资源竞争（Go 的 race detector 底层依赖 C 代码（需要 cgo 支持）默认情况下，Windows 下安装的 Go 有时会把 CGO_ENABLED=0，导致 -race 用不了）
		go run main.go：
			编译并运行 main.go
			不会检查资源竞争
			程序结果可能是错的，但不会提示


		go run -race main.go
			这是带竞态检测（Race Detector）的运行方式：
			-race 会在编译时插入额外的检查逻辑
			程序运行时会监控内存访问，发现 多个 goroutine 并发访问同一个变量且至少有一个写操作，就会报 DATA RACE 错误

	*/
	// 展示资源竞争
	wg.Add(2)
	go withdraw("JimLee", 800)
	go withdraw("BruceLee", 500)
	wg.Wait()
	fmt.Printf("最终余额为: %d", balance)

}

func withdraw(name string, amount int) {
	defer wg.Done()
	if balance >= amount {
		fmt.Printf("%s 正在取钱: %d. 当前余额: %d\n", name, amount, balance)
		time.Sleep(10 * time.Millisecond) // 模拟取钱耗时
		balance -= amount
		fmt.Printf("%s 取钱成功，取出: %d，剩余余额: %d\n", name, amount, balance)
	} else {
		fmt.Printf("%s 取钱失败，余额不足（尝试取: %d，余额: %d）\n", name, amount, balance)
	}

}

/*
最终输出:
	BruceLee 正在取钱: 500. 当前余额: 1000
	JimLee 正在取钱: 800. 当前余额: 1000
	JimLee 取钱成功，取出: 800，剩余余额: 200
	BruceLee 取钱成功，取出: 500，剩余余额: -300
	最终余额为: -300
两个 goroutine 同时读取 balance，都以为够取，结果都扣钱，最终余额出错，这就是典型的 资源竞争


race 检测输出
	PS C:\demo\GolangTraining> go run -race .\07_goroutines\02_race-condition\01_race-condition-show\main.go
	JimLee 正在取钱: 800. 当前余额: 1000
	BruceLee 正在取钱: 500. 当前余额: 1000
	JimLee 取钱成功，取出: 800，剩余余额: 200
	==================
	WARNING: DATA RACE
	Read at 0x00014011b208 by goroutine 9:
	  main.withdraw()
		  C:/demo/GolangTraining/07_goroutines/02_race-condition/01_race-condition-show/main.go:48 +0x1fa
	  main.main.gowrap2()
		  C:/demo/GolangTraining/07_goroutines/02_race-condition/01_race-condition-show/main.go:37 +0x3a

	Previous write at 0x00014011b208 by goroutine 8:
	  main.withdraw()
		  C:/demo/GolangTraining/07_goroutines/02_race-condition/01_race-condition-show/main.go:48 +0x212
	  main.main.gowrap1()
		  C:/demo/GolangTraining/07_goroutines/02_race-condition/01_race-condition-show/main.go:36 +0x3a

	Goroutine 9 (running) created at:
	  main.main()
		  C:/demo/GolangTraining/07_goroutines/02_race-condition/01_race-condition-show/main.go:37 +0x44

	Goroutine 8 (finished) created at:
	  main.main()
		  C:/demo/GolangTraining/07_goroutines/02_race-condition/01_race-condition-show/main.go:36 +0x38
	==================
	BruceLee 取钱成功，取出: 500，剩余余额: -300
	最终余额为: -300Found 1 data race(s)s
	exit status 66
*/
