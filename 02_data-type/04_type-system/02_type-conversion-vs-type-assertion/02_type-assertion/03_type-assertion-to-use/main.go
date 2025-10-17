package main

import "fmt"

// main 使用断言
func main() {
	var val any = 5 // 等价于 var val interface{} = 5
	// 1. 使用 %T 打印接口里的动态类型
	fmt.Printf("值: %v - 类型: %T\n", val, val) // 值: 5 - 类型: int

	// 2. 接口值不能直接使用
	//fmt.Println(val + 10) // 会抛出错误，虽然上面打印 val 的动态类型为 int，但是编译器只看到左边是 interface{}，不允许和 int 直接相加

	// 3. 用类型断言取出再运算（val.(int) 是 类型断言，从接口里取出 int。没有 ok 的断言写法，若失败会 panic；更稳妥的是使用 ok）
	//fmt.Println(val.(int) + 5) // 10

	// 4. 使用 ok 接收断言返回值后再进行操作
	if n, ok := val.(int); ok {
		fmt.Println(n + 10)
	} else {
		fmt.Printf("not int type, can't calculate!")
	}
}
