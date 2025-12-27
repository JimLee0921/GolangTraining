package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 1, 2, 4}
	fmt.Println(slices.Contains(numbers, 1))
	fmt.Println(slices.Contains(numbers, 10))
}
