package main

import (
	"fmt"
	"time"
)

// withTimer 返回一个新的函数（包装了计时逻辑）
func withTimer(fn func()) func() {
	return func() {
		start := time.Now()
		fn()
		fmt.Println("Elapsed:", time.Since(start))
	}
}

// withRecover 包裹上一层异常捕获逻辑
func withRecover(fn func()) func() {
	return func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered from", r)
			}
		}()
		fn()
	}
}

func sayHello() {
	/*
		装饰器可以嵌套
		最外层函数最后执行
		执行顺序：withRecover -> withTimer -> sayHello
	*/
	for i := 1; i <= 100; i++ {
		fmt.Println("Hello World!")
		time.Sleep(20 * time.Millisecond)
		if i == 77 {
			panic("unknow error")
		}
	}
}

func main() {
	decorated := withRecover(withTimer(sayHello))
	decorated()
}
