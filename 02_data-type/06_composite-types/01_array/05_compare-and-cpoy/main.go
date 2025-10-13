package main

import "fmt"

func main() {
	/*
		数组是值类型：Value Type
		赋值或传参会 复制整个数组
		可以把一个同类型、同长度的数组整体赋给另一个
		但是数组赋值时是 复制，不是引用，a 和 b 各自独立
	*/
	arrayA := [3]int{1, 2, 3}
	arrayB := [3]int{4, 5, 6}
	arrayA = arrayB // 直接把 arrayB 的内容整体复制给 arrayA
	fmt.Printf("arrayA: %v\n", arrayA)
	fmt.Printf("arrayB: %v\n", arrayB)
	arrayB[0] = 100
	arrayA[2] = 100
	fmt.Printf("arrayA: %v\n", arrayA)
	fmt.Printf("arrayB: %v\n", arrayB)

	/*
		Go 支持 数组比较（==、!=），只要元素类型可比较
	*/
	x := [3]int{1, 2, 3}
	y := [3]int{1, 2, 3}
	z := [3]int{3, 2, 1}

	fmt.Println(x == y) // true
	fmt.Println(x == z) // false

}
