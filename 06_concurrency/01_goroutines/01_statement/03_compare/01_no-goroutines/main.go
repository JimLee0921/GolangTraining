package main

import (
	"fmt"
	"time"
)

// main 同步执行
func main() {
	/*
		main 函数里依次调用 foo() 和 bar()
		foo() 会打印 "Foo: 0" 到 "Foo: 44"
		bar() 会打印 "Bar: 0" 到 "Bar: 44"
		因为是顺序同步调用，所以会先完整打印 45 行 Foo，然后再打印 45 行 Bar
		而这里是延时进行的，所以消耗时间就是 foo + bar 的执行时间
	*/
	Foo()
	Bar()
}

func Foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(100 * time.Millisecond) // 延时 0.1 秒
	}
}

func Bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(100 * time.Millisecond) // 延时 0.1 秒
	}
}
