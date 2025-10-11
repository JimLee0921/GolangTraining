package main

import (
	"fmt"
	"time"
)

func safeGo(fn func()) {
	/*
		goroutine 捕获到异常就打印，没有捕获到就正常执行
	*/
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("goroutine recovered:", r)
			}
		}()
		fn()
	}()
}

func main() {
	// panic 只在当前 goroutine 中传播，因此要在每个 goroutine 中单独 recover
	safeGo(func() {
		panic("worker panic")
	})
	time.Sleep(time.Second)
	fmt.Println("main still alive")
	/*
		goroutine recovered: worker panic
		main still alive
	*/
}
