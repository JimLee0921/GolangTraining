package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "task successful"
	}()

	select {
	case res := <-ch:
		fmt.Println("result: ", res)
	case <-time.After(2 * time.Second): // 等待任务执行超过两秒则为超时
		fmt.Println("task timeout")
	}
}
