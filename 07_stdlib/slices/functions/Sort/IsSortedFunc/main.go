package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	names := []string{"alice", "Bob", "VERA"}
	isSortedInsensitive := slices.IsSortedFunc(names, func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})

	fmt.Println(isSortedInsensitive)
	fmt.Println(slices.IsSorted(names))
}
