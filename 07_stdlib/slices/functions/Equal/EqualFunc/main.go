package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	numbers := []int{0, 42, 8}
	strings := []string{"000", "42", "0o10"}

	equal := slices.EqualFunc(numbers, strings, func(i int, s string) bool {
		sn, err := strconv.ParseInt(s, 0, 64)
		if err != nil {
			return false
		}
		return i == int(sn)
	})
	fmt.Println(equal)
}
