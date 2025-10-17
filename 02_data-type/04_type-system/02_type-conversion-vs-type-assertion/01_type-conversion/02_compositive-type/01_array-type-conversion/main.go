package main

import "fmt"

type MyArr [3]int

func main() {
	/*
		数组的类型包括：元素类型 + 数组长度
		必须完全相同才能转换
	*/
	var a1 [3]int = [3]int{1, 2, 3}
	var a2 [3]int = a1 // 完全相同类型
	// var a3 [4]int = a1     // 长度不同不能转换
	fmt.Println(a1, a2)
	// 自定义类型数组
	var base [3]int = [3]int{1, 2, 3}

	var a MyArr = MyArr(base) // 显式转换（底层相同）
	fmt.Println(a)
}
