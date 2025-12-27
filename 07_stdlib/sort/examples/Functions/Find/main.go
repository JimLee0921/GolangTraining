package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	a := []string{"apple", "banana", "lemon", "mango", "pear", "strawberry"}
	for _, x := range []string{"banana", "orange"} {
		i, found := sort.Find(len(a), func(i int) int {
			return strings.Compare(x, a[i])
		})
		if found {
			fmt.Printf("found %s at index %d\n", x, i)
		} else {
			fmt.Printf("%s not found, would insert at %d\n", x, i)
		}
	}
}
