package main

import "fmt"

func main() {
	/*
		数组是值类型：Value Type 赋值或传参会 复制整个数组
		所以可以把一个同类型、同长度的数组整体赋给另一个
		但是数组赋值时是 复制，不是引用，a 和 b 各自独立
		数组定义完成后重新赋值不能用不同长度的数组
		其中在 Go 里，数组长度是类型的一部分，所以多维数组的每一个维度长度必须相同
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

	arrayA = [3]int{66, 77, 88}
	//arrayA = [1]int{2} 这里会报错
	fmt.Printf("arrayA: %v\n", arrayA)

}
