package main

import (
	"fmt"
	"time"
)

func TimerDemo() {
	fmt.Println("wait three seconds...")
	timer := time.NewTimer(3 * time.Second)
	<-timer.C // 阻塞等待 Timer 触发
	fmt.Println("three times after, go!")

}

func TickerDemo() {
	ticker := time.NewTicker(1 * time.Second) // 每 1 秒触发一次
	defer ticker.Stop()                       // 用完要停止，否则可能泄漏

	count := 0

	for t := range ticker.C {
		fmt.Println("Tick at:", t)
		count++
		if count == 5 {
			break // 跳出循环，停止
		}
	}

	fmt.Println("Ticker stop")
}

func main() {
	//TimerDemo()
	TickerDemo()
}
