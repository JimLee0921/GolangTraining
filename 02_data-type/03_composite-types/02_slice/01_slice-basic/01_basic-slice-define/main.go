package main

import "fmt"

// main slice的基本创建方式
func main() {
	/*
		var zero []int			  // 仅声明类型，初始化为空值
		var zero = []int{1,2,3}	  // 初始化创建
		slice := []Type{elements} // 简写初始化创建
		slice := []Type			  // 简写跳初始化为空值
	*/
	// 方式一：仅声明类型，全是空值 nil 切片
	var zero []int
	fmt.Printf("空切片: %#v | 长度len=%d 容量cap=%d | isNil=%v\n", zero, len(zero), cap(zero), zero == nil)

	// 方式二：声明类型且初始化
	var initSlice = []int{1, 2, 3}
	fmt.Println(initSlice)

	// 方式二：声明时提供字面量
	var literal = []string{"Go", "Rust", "Python"}
	fmt.Printf("literal: %#v\n", literal)

	// 方式三：使用 := 简写并初始化
	shorthand := []float64{3.14, 1.41, 2.71}
	fmt.Printf("shorthand: %#v\n", shorthand)
}
