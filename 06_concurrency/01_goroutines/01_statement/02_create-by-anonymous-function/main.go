package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		goroutine 可以直接启动匿名函数，常用于
			内联逻辑
			启动临时任务
			向 goroutine 传参
	*/
	go func(name string) {
		fmt.Println("Hello", name)
	}("Alice")

	fmt.Println("Main done!")
	time.Sleep(time.Second)
}
