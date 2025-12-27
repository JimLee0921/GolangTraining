package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Gopher", 13},
		{"Alice", 20},
		{"Bob", 5},
		{"Vera", 24},
		{"Zac", 20},
	}

	sortFunc := func(x, y Person) int {
		return cmp.Compare(x.Age, y.Age)
	}

	s := slices.SortedStableFunc(slices.Values(people), sortFunc)
	fmt.Println(s)
}
