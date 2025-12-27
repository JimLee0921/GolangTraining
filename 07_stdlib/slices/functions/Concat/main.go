package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{0, 1, 2, 3}
	s2 := []int{4, 5, 6}
	s3 := []int{7, 8, 9}
	res := slices.Concat(s1, s2, s3)
	fmt.Println(res)
}
