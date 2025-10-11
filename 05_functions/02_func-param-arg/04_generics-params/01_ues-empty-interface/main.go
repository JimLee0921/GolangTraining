package main

import "fmt"

func PrintAny(v any) {
	fmt.Printf("Type: %T, Value: %v\n", v, v)
}

func main() {
	/*
		编译器会生成一个函数：
		func PrintAny(v interface{}) { ... }
		所有类型都要在运行时装箱（boxing），用反射打印
	*/
	PrintAny(123)
	PrintAny("hi")
}
