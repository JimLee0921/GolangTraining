package main

import (
	"fmt"
	"time"
)

func AfterDemo() {
	fmt.Println("wait 2 seconds...")
	<-time.After(2 * time.Second) // 在这里阻塞等待 2 秒
	fmt.Println("2 seconds after, go!")
}
func TickDemo() {
	// 无法停止
	for t := range time.Tick(1 * time.Second) { // 每秒触发一次
		fmt.Println("Tick at:", t)
	}
}

func AfterFuncDemo() {
	timer := time.AfterFunc(3*time.Second, func() {
		fmt.Println("it's time to run")
	})
	defer timer.Stop()
	fmt.Println("afterfuncdemo is running")

	time.Sleep(5 * time.Second) // 等待回调函数执行

}

func main() {
	//AfterDemo()
	//TickDemo()
	AfterFuncDemo()
}
