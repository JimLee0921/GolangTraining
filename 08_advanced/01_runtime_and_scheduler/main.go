package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func main() {
	// 查看 CPU 信息
	fmt.Println("CPU 核心数:", runtime.NumCPU())

	// 控制最大并行数
	old := runtime.GOMAXPROCS(0)
	fmt.Println("当前 GOMAXPROCS:", old)
	runtime.GOMAXPROCS(runtime.NumCPU()) // 设置为 CPU 数
	fmt.Println("设置后 GOMAXPROCS:", runtime.GOMAXPROCS(0))

	// 查看当前 goroutine 数量
	fmt.Println("当前 goroutine 数:", runtime.NumGoroutine())

	// 手动让出 CPU 调度权
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程执行:", i)
			runtime.Gosched() // 让出 CPU
		}
	}()
	for i := 0; i < 3; i++ {
		fmt.Println("主协程执行:", i)
	}

	// 等待 goroutine 执行完
	time.Sleep(200 * time.Millisecond)
	fmt.Println("当前 goroutine 数:", runtime.NumGoroutine())

	// 手动触发 GC
	fmt.Println("\n===== 垃圾回收示例 =====")
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	fmt.Printf("GC 前: Alloc=%.2fMB, NumGC=%d\n", float64(m1.Alloc)/1024/1024, m1.NumGC)

	junk := make([]byte, 20<<20) // 20MB
	_ = junk

	runtime.GC() // 触发一次 GC
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)
	fmt.Printf("GC 后: Alloc=%.2fMB, NumGC=%d\n", float64(m2.Alloc)/1024/1024, m2.NumGC)

	// 打印当前栈信息
	fmt.Println("\n===== 打印堆栈信息 =====")
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	fmt.Println(string(buf[:n]))

	// 打印调用信息
	fmt.Println("\n===== 调用者信息 =====")
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		fn := runtime.FuncForPC(pc)
		fmt.Printf("函数: %s\n文件: %s\n行号: %d\n", fn.Name(), file, line)
	}

	// 内存状态监控
	fmt.Println("\n===== 读取内存统计 =====")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc=%.2fMB, TotalAlloc=%.2fMB, Sys=%.2fMB, NumGC=%d\n",
		float64(m.Alloc)/1024/1024,
		float64(m.TotalAlloc)/1024/1024,
		float64(m.Sys)/1024/1024,
		m.NumGC)

	// 归还空闲内存
	debug.FreeOSMemory()
	fmt.Println("已请求将空闲内存归还操作系统。")
}
