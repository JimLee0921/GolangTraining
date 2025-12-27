package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Gopher", 13},
		{"Alice", 20},
		{"Bob", 24},
		{"Alice", 55},
	}
	slices.SortStableFunc(people, func(a, b Person) int {
		return strings.Compare(a.Name, b.Name)
	})
	fmt.Println(people)
}
