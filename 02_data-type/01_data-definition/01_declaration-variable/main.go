package main

import "fmt"

// main 演示 Go 中几种常见的变量声明方式。
func main() {
	// 方式一：使用 var 指定类型，未赋值时会得到该类型的零值
	var count int
	var name string
	var floatOne float32
	var floatTwo float64
	var isTrue bool
	fmt.Println("未赋值的 count 默认是零值:", count)
	fmt.Println("未赋值的 name 默认是空字符串:", name)
	fmt.Println("未赋值的 floatOne 默认是零值:", floatOne)
	fmt.Println("未赋值的 floatTwo 默认是零值:", floatTwo)
	fmt.Println("未赋值的 isTrue 默认是 false:", isTrue)

	// 方式二：使用 var 同时指定类型和值，便于阅读。
	var message string = "你好，Go" // 也可以写成 var message = "你好，Go"
	fmt.Println("显式声明并初始化 message:", message)

	// 方式三：一次声明多个同类型变量。
	var width, height int = 1920, 1080
	fmt.Println("width 和 height:", width, height)

	// 方式四：短变量声明 := 只能在函数内部使用，类型由右值推断。
	ratio := 1.618
	fmt.Println("ratio 通过 := 推断为 float64:", ratio)

	// 方式五：分组声明便于整理相关变量。
	var (
		username string = "gopher"
		age      int    = 8
		active   bool
	)
	fmt.Println("分组声明的变量:", username, age, active)
}
