package main

import "fmt"

func trace(msg string) func() {
	fmt.Println("enter:", msg) // 立即执行：记录进入
	return func() {            // 返回一个收尾函数
		fmt.Println("exit:", msg)
	}
}

func main() {
	/*
		1 先调用 trace("main")
		2 得到返回的 func()
		3 把它 defer 到 main 退出时执行
	*/
	defer trace("main")()
	fmt.Println("doing something")
	/*
		enter: main
		doing something
		exit: main
	*/
}
