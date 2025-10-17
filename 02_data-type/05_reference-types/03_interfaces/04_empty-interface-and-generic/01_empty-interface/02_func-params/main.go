package main

import "fmt"

// PrintAny 函数参数使用空接口（通用容器）
func PrintAny(v any) {
	fmt.Println(v)
}

func main() {
	PrintAny(123)
	PrintAny("abc")
	PrintAny([]string{"Go", "Rust", "Python"})
}
