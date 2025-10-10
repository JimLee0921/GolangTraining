package main

import (
	"fmt"
	"time"
)

// main 测试并发执行
func main() {
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
