package main

import "fmt"

func Map[T any, R any](src []T, f func(T) R) []R {
	result := make([]R, len(src))
	for i, v := range src {
		result[i] = f(v)
	}
	return result
}

func main() {
	/*
		[T, R any] 表示输入和输出可以是不同类型
		这在泛型算法中非常常见
		Go 没有内置 map/filter/reduce，泛型使它变得自然
	*/
	nums := []int{1, 2, 3}
	strs := Map(nums, func(n int) string {
		return fmt.Sprintf("num:%d", n)
	})
	fmt.Println(strs) // [num:1 num:2 num:3]
}
