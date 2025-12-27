package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []int{1, 4, 5, 66, 112, 424}
	n, found := slices.BinarySearch(names, 4)
	fmt.Println("Bruce:", n, found)

	n, found = slices.BinarySearch(names, 6)
	fmt.Println("Bill:", n, found)
}
