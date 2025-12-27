package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"apple", "banana", "cherry", "date", "fig", "grape"}

	x := "banana"
	i := sort.SearchStrings(a, x)
	fmt.Printf("found %s at index %d in %v\n", x, i, a)

	x = "coconut"
	i = sort.SearchStrings(a, x)
	fmt.Printf("%s not found, can be inserted at index %d in %v\n", x, i, a)

}
