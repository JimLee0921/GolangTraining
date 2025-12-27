package main

import (
	"fmt"
	"slices"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}

	type People []Person

	people := People{
		{"Gopher", 13},
		{"Alice", 20},
		{"Bob", 5},
		{"Vera", 24},
		{"Zac", 15},
	}

	for c := range slices.Chunk(people, 2) {
		fmt.Println(c)
	}
}
