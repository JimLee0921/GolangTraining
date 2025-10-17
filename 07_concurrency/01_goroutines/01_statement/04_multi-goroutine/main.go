package main

import (
	"fmt"
	"time"
)

func main() {
	// 如果是  Go 1.22  之前的版本会存在闭包捕获到问题，需要手动传入参数或在 goroutine 中重新声明一遍，否则会出现变量被多个 goroutine 重复操作的情况
	for i := 0; i <= 10; i++ {
		go func() {
			fmt.Println("get task:", i)
			time.Sleep(time.Millisecond * 100)
			fmt.Println("done task:", i)
		}()
	}
	time.Sleep(time.Second)
}
