package main

import "fmt"

func PrintAny[T any](v T) {
	fmt.Println(v)
}
func main() {
	/*
	 any 表示任意类型，等价于 interface{} 空接口
	*/
	PrintAny("123")
	PrintAny(555)
}
