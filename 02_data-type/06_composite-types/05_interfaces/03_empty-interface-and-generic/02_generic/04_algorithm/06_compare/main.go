package main

import "fmt"

type Ordered interface {
	~int | ~float64 | ~string
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	/*
		泛型约束可以定义有序类型（Ordered）
		支持任意可比较的基础类型
		常见于排序、比较算法
	*/
	fmt.Println(Max(3, 5))       // 5
	fmt.Println(Max(2.3, 4.8))   // 4.8
	fmt.Println(Max("go", "hi")) // hi
}
