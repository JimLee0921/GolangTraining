package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 5, 55, 2, 13}
	seq := slices.Values(s)

	s2 := slices.SortedFunc(seq, func(i int, i2 int) int {
		return cmp.Compare(i2, i) // 逆序
	})
	fmt.Println(s2)
}
