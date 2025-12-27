package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, -10, 8}
	hasNegative := slices.ContainsFunc(numbers, func(i int) bool {
		return i < 0
	})
	fmt.Println("has negative number:", hasNegative)

	hasOdd := slices.ContainsFunc(numbers, func(i int) bool {
		return i%2 != 0
	})
	fmt.Println("Has odd number:", hasOdd)

}
