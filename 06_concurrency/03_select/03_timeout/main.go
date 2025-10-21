package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		加入超时机制防止永远堵塞
		time.After() 会在指定时间后向返回的通道发送一个值
		可以用来实现超时控制或请求取消机制
	*/
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "task successful"
	}()

	select {
	case res := <-ch:
		fmt.Println("received: ", res)
	case <-time.After(2 * time.Second): // 等待任务执行超过两秒则为超时
		fmt.Println("task timeout")
	}
}
