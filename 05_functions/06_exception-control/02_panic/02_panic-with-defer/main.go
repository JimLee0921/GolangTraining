package main

import "fmt"

func main() {
	/*
		即使 panic 发生，panic 之前的 defer 依然会被调用
		panic 不会跳过 defer —— 所有已注册的 defer 一定执行
		这就是为什么可以在 defer 里用 recover() 来拦截 panic
	*/
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	panic("boom")
	defer fmt.Println("defer3") // 不会执行

	/*
		defer2
		defer1
		panic: boom
	*/
}
