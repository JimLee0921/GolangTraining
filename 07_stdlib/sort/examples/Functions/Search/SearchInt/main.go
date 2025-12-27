package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 23, 35, 64, 242, 1241}
	x := 23
	i := sort.SearchInts(a, x)
	if a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)

	} else {
		fmt.Printf("%d not found, can be inserted at index %d in %v\n", x, i, a)
	}

	x = 3
	i = sort.SearchInts(a, x)
	if a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)

	} else {
		fmt.Printf("%d not found, can be inserted at index %d in %v\n", x, i, a)
	}
}
