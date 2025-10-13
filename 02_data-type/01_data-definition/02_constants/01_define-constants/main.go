package main

import "fmt"

const Pi = 3.14159

// main 定义常量
func main() {
	/*
		常量 在编译时就确定值，运行时不可更改
		使用关键字 const 定义
			常量的值必须是 编译期可确定的：数字、字符串、布尔值、常量表达式，不能是 slice，map，struct等
			var 定义的是变量，运行时可以改变
			const 定义的是常量，运行时不能修改
			常量可以有显式类型，也可以没有（无类型常量，untyped constant）
				无类型常量编译器会推迟确定类型，可以赋值给不同类型的变量
				有类型常量只能复制给指定类型的变量
			不能使用变量为常量赋值，因为常量必须在编译时就确定

	*/
	// 无类型常量

	const b = 4
	var x int = b
	var y float64 = b
	fmt.Println(x, y) // 编译器会推迟确定类型，可以赋值给不同类型的变量
	// 有类型常量
	const a int = 10 // 只能赋值给 int 类型的变量
}
