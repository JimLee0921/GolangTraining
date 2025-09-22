package main

import "fmt"

// main 切片的多种声明与初始化方式
func main() {
	/*
		切片其实就是对数组的一个“引用视图”，底层依然是数组
		和数组不同，切片的长度可以变化（动态扩展/缩减）
		切片比数组更灵活，因此实际开发中用切片远多于数组
	*/
	// 方式一：仅声明类型，全是空值 nil 切片
	var zero []int
	fmt.Printf("空切片: %#v | 长度len=%d 容量cap=%d | isNil=%v\n", zero, len(zero), cap(zero), zero == nil)

	// 方式二：声明时提供字面量
	var literal = []string{"Go", "Rust", "Python"}
	fmt.Printf("literal: %#v\n", literal)

	// 方式三：使用 := 简写并初始化
	shorthand := []float64{3.14, 1.41, 2.71}
	fmt.Printf("shorthand: %#v\n", shorthand)

	/*
		方式四：使用 make 创建切片
		语法：make([]T, len, cap)：
			len：切片的长度（初始化后能直接用下标访问的元素个数）
			cap：切片的容量（底层数组大小）
		规则：必须满足 0 <= len <= cap
			如果 len > cap，会直接编译报错
			如果 len == cap，切片刚好占满底层数组。
			如果 len < cap，切片“预留”了一些空间，方便后续 append
			省略 cap 时，cap = len
			len=0 但 cap>0 时，切片为空，但有预留容量，只有在 0 <= index < len(slice) 时，下标赋值才合法
	*/
	sized := make([]int, 3, 5)
	sized[0] = 42
	fmt.Printf("sized: %#v | len=%d cap=%d\n", sized, len(sized), cap(sized))

	// 仅分配容量，后续通过 append 填充
	buffered := make([]int, 0, 4)
	buffered = append(buffered, 10, 20, 30)
	fmt.Printf("buffered: %#v | len=%d cap=%d\n", buffered, len(buffered), cap(buffered))

	// 方式六：通过数组切片获得视图，共享底层存储
	source := [5]int{1, 2, 3, 4, 5}
	window := source[1:4]
	fmt.Printf("window: %#v | len=%d cap=%d\n", window, len(window), cap(window))

	// 方式七：构造切片的切片，用于表示二维结构
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("matrix: %#v | rows=%d\n", matrix, len(matrix))
}
