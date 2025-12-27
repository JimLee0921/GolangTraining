package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{1, 5, 2, 31, 12, 55, 2}
	s2 := []int{55, 55, 11, 11}
	res := slices.AppendSeq(s2, slices.Values(s1))
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(res)
}
