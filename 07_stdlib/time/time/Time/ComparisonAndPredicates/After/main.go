package main

import (
	"fmt"
	"time"
)

var c chan int

func handle(int) {

}

func main() {
	select {
	// 永远不会就绪
	case m := <-c:
		handle(m)
	// 立即执行 time.After(10*time.Second) 注册一个10秒后的定时时间，返回一个 channel 记为 ch
	// select 堵塞等待：c -> 永远堵塞 ch -> 10秒后可读
	// 10秒后 runtime 向 ch 发送一个 time.Time，select 被唤醒，执行该 case
	case <-time.After(10 * time.Second):
		fmt.Println("timed out")
	}
}
