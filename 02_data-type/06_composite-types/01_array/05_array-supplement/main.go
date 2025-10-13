package main

import "fmt"

func main() {

	/*
		数组定义完成后重新赋值不能用不同长度的数组
		并且数组长度是类型的一部分，所以多维数组的每一个维度长度必须相同
	*/
	arrayA := [3]int{66, 77, 88}
	//arrayA := [1]int{2} 这里会报错
	fmt.Printf("arrayA: %v\n", arrayA)

}
