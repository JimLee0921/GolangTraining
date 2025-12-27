package main

import (
	"fmt"
	"slices"
)

func main() {
	seq := []int{1, 4, 1231, 2, 41, 123, 5}
	slices.Reverse(seq)
	fmt.Println(seq)
}
