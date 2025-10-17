package main

import "fmt"

func main() {
	/*
		空接口（interface{}）指的是：一个不包含任何方法的接口类型
		空接口没有任何方法要求，所以任何类型都自动实现它，所有类型都是 interface{} 的子集
		从 Go 1.18（2022 年）开始，Go 官方引入了一个内置别名：type any = interface{}
		any 和 interface{} 完全等价
			底层库：用 interface{}
			业务代码 / 泛型：用 any，更清晰地表达任意类型的意思

		常用于：
			泛型容器
			JSON
			日志打印
			动态类型值
	*/
	// v 是一个空接口，能存放任何类型，编译器允许，因为所有类型都实现了空接口
	var v interface{} // 等价于 var v any
	v = 42
	fmt.Println(v)
	v = "hello"
	fmt.Println(v)
	v = []int{1, 2, 3}
	fmt.Println(v)
}
