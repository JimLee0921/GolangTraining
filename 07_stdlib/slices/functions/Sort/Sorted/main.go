package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 5, 55, 2, 13}
	seq := slices.Values(s)

	s2 := slices.Sorted(seq)
	fmt.Println(s2)
}
