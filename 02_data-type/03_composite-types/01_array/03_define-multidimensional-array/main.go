package main

import (
	"fmt"
)

func main() {

	// 1. 创建多维数组，默认为0或其它类型空值
	var numberMatrixA [3][2]int
	fmt.Println(numberMatrixA)

	// 2. 创建多维数组并全部初始化
	stringMatrixA := [3][2]string{
		{"hello", "bro"},
		{"you", "are"},
		{"funny", "man"},
	}
	fmt.Println(stringMatrixA)

	// 3. 创建多维数组部分初始化
	numberMatrixB := [2][3]int{
		{2: 3},
		{2, 3},
	}
	fmt.Println(numberMatrixB)

	// 4. 创建多位数组使用 ... 推断最外层数组的长度（只能推断最外层数组）
	stringMatrixB := [...][11]string{
		{10: "shab"},
		{"yes", "bro"},
	}
	fmt.Printf("%q", stringMatrixB)
}
