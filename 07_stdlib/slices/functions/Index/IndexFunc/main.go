package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, -10, 8}
	i := slices.IndexFunc(numbers, func(i int) bool {
		// 查找第一个负数下标
		return i < 0
	})

	fmt.Println("the first negative at index:", i)
}
