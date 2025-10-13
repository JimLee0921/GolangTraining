package main

import "fmt"

func main() {
	// 字面量初始化
	c1 := 1 + 2i
	c2 := -3.5 + 4.2i

	// 使用内置函数
	c3 := complex(3.0, 4.0)

	// 显式类型声明
	var c4 complex64 = 5 + 6i
	var c5 complex128 = complex(7, 8)

	fmt.Println(c1, c2, c3, c4, c5)
}

/*
(1+2i) (-3.5+4.2i) (3+4i) (5+6i) (7+8i)
*/
