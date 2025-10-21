package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func main() {
	/*
		如果不加 time.Sleep，可能会什么都没输出（因为主 goroutine 退出后程序就结束）
		注意：
			Go 程序中所有 goroutine 运行在同一个进程内
			主 goroutine（即 main()）结束时，整个程序会立即退出
			因此如果希望看到 goroutine 执行结果，必须同步或等待
	*/
	go sayHello()             // 创建一个新的 goroutine
	fmt.Println("Main done.") // 主 goroutine 继续执行
	time.Sleep(time.Second)   // 等待一会儿让 goroutine 运行完
}
