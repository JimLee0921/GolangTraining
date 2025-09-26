package main

import "fmt"

func main() {
	/*
		Go 支持 可变参数（类似 Python 的 *args 或 Java 的 ...）
		在 Go 里，如果函数的最后一个参数写成 ...T（三个点加类型），就表示这个参数可以接收 零个或多个该类型的值
		在函数体内部，这个参数会被当成一个 切片 ([]T) 来使用
		注意事项：
			在 Go 语言里，一个函数参数列表里只能有一个可变参数，并且必须放在最后
			Go 的可变参数语法是 ...T，这里的 T 必须是同一种类型
			可变参数也可以什么都不传入（也就是个空切片）
	*/
	greeting("JimLee")
	greeting("BruceLee", "JamesBond", "Skrillex")

}

func greeting(names ...string) {
	fmt.Println("names 切片内容", names)
	for _, name := range names {
		fmt.Println("hello!", name)
	}

}
