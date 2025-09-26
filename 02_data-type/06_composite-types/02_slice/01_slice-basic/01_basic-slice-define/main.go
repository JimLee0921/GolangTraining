package main

import "fmt"

// main slice的基本创建方式
func main() {
	/*
		切片其实就是对数组的一个引用视图，底层依然是数组
		和数组不同，切片的长度可以变化（动态扩展/缩减）
		切片比数组更灵活，因此实际开发中用切片远多于数组
		切片和数组中的元素都是可以重复的
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
