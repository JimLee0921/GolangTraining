package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, -10, 8}
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))

	grow := slices.Grow(numbers, 2)
	fmt.Println(grow)
	fmt.Println(cap(grow)) // 扩容策略不保证这里是6
	fmt.Println(len(grow))
}
