package main

import "fmt"

func sum(nums ...int) int {
	// 接收可变参数
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// main 切片展开式
func main() {
	/*
		`slice...` 切片展开语法
		slice expansion，主要用在 函数调用（可变参数函数） 或 append（拼接切片） 里
	*/
	// 拼接切片
	a := []int{1, 2}
	b := []int{3, 4}

	a = append(a, b...) // 把 b 展开成 3,4 插到 a 里
	fmt.Println(a)      // [1 2 3 4]

	// 函数调用
	fmt.Println(sum(a...))

}
