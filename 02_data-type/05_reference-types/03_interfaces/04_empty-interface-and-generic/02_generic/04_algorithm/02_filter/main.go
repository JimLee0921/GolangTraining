package main

import "fmt"

func Filter[T any](src []T, test func(T) bool) []T {
	var result []T
	for _, v := range src {
		if test(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	/*
		泛型 Filter 不依赖具体类型
		编译期仍然保持类型安全
		适用于 string、struct、float 等任意类型
	*/
	nums := []int{1, 2, 3, 4, 5, 6}
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println(evens) // [2 4 6]
}
