package main

import "fmt"

func main() {
	/*
		函数基本类型定义结构
		小 TIPS：
			Parameters（形参，参数）函数定义时用的名字。是函数声明里的一部分，描述函数需要哪些输入
			Arguments（实参，实参值）调用时传进去的实际值
		func 函数名(参数名 参数类型, 参数名2 参数类型2, ...) 返回类型 {
			// 函数体
		}
		在 Go 里，如果一个函数声明时没有返回值，那么调用这个函数时它不会返回任何东西
			没有默认返回值（不会像某些语言默认返回 null、nil 或者 0）
			如果在函数里写了 return，只能单纯地结束函数执行，不能带返回值
			如果函数签名里没有返回值，也不能在调用处接收返回结果
			返回值也可以多个：(返回类型1, 返回类型2)
	*/
	hello("Go") // 只执行，不返回值
	// result := hello("Go") // 编译错误，hello 没有返回值
}

func hello(name string) {
	fmt.Println("Hello,", name)
	// 这里只能写 return，不带值
}
