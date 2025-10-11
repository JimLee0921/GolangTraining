package main

import "fmt"

func PrintGeneric[T any](v T) {
	fmt.Printf("Type: %T, Value: %v\n", v, v)
}

func main() {
	/*
		编译器会生成两个版本：
			func PrintGeneric_int(v int) { ... }
			func PrintGeneric_string(v string) { ... }
		在编译时就确定了类型，不需要 interface 装箱，也不需要反射
	*/
	PrintGeneric(123)
	PrintGeneric("hi")
}
