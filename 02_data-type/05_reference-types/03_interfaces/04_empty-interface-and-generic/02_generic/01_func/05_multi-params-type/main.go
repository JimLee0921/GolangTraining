package main

import "fmt"

// Map 多个数据类型泛型，约束一样可以一起书写
func Map[T, R any](arr []T, f func(T) R) []R {
	result := make([]R, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

func main() {
	nums := []int{1, 2, 3}
	strs := Map(nums, func(n int) string {
		return fmt.Sprintf("num:%d", n)
	})
	fmt.Println(strs) // [num:1 num:2 num:3]
}
