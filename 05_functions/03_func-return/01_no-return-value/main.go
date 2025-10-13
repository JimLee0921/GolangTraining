package main

import "fmt"

func main() {
	/*
		Go 中函数如果不需要返回值，直接省略返回值类型即可
		如果函数没有返回值，那么在调用时不能进行接收，否则会编译错误
		即使没有返回值，函数中仍然可以使用 return 来 提前结束函数执行
		常用于只做动作、不需要反馈的场景
	*/
	greeting("Skrillex", true)
	greeting("Tom", false)
	//res := greeting("JimLee", true)  编译错误，不能定义变量接收无返回值的函数

}

// greeting 方法不返回任何内容
func greeting(name string, print bool) {
	if !print {
		return // 直接结束
	}
	fmt.Println("Hello " + name)

}
