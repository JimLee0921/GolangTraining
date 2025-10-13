package main

import "fmt"

func main() {
	// 1. 显式指定类型
	var a int8 = 100
	var b uint16 = 50000
	var c int = 42 // 平台相关（一般为 int64）

	// 2. 类型推断
	d := 1234 // 自动推断为 int
	e := -567 // 自动推断为 int
	fmt.Printf("a: value: %d, type: %T\n", a, a)
	fmt.Printf("b: value: %d, type: %T\n", b, b)
	fmt.Printf("c: value: %d, type: %T\n", c, c)
	fmt.Printf("d: value: %d, type: %T\n", d, d)
	fmt.Printf("e: value: %d, type: %T\n", e, e)
}

/*
a: value: 100, type: int
b: value: 50000, type: uint16
c: value: 42, type: int
d: value: 1234, type: int
e: value: -567, type: int
*/
