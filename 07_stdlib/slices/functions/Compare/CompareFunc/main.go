package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
)

func main() {
	numbers := []int{0, 43, 8}
	strings := []string{"0", "0", "8"}
	result := slices.CompareFunc(numbers, strings, func(i int, s string) int {
		sn, err := strconv.Atoi(s)
		if err != nil {
			return 1
		}
		return cmp.Compare(i, sn)
	})
	fmt.Println(result)
}
