package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, -10, 8}
	maxValue := slices.Max(numbers)
	fmt.Println(maxValue)
	maxValue = 1000
	fmt.Println(numbers)
}
