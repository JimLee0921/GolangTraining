package main

import (
	"fmt"
	"sync"
)

func main() {
	initHeavy := func() {
		fmt.Println("heavy init runs exactly once")
	}
	onlyOnce := sync.OnceFunc(initHeavy) // 返回一个函数，多次调用也只执行一次

	onlyOnce()
	onlyOnce()
	onlyOnce()
	// 输出只会打印一行：heavy init runs exactly once
}
