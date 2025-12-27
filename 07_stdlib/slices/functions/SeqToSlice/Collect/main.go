package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{1, 5, 2, 31, 12, 55, 2}

	res := slices.Collect(slices.Values(s1))
	fmt.Println(res)
}
