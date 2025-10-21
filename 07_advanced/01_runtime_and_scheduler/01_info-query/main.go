package main

import "C"
import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())       // 返回当前系统的逻辑 CPU 数量（例如机器 8 核，就返回 8）
	fmt.Println(runtime.NumGoroutine()) // 返回当前活跃的 goroutine 数量
	fmt.Println(runtime.NumCgoCall())   // 返回当前进程中已执行过的 cgo 调用次数（即 Go 调用 C 函数的次数）
}
