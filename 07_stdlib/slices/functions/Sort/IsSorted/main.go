package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(slices.IsSorted([]string{"Annie", "JimLee", "Vera"}))
	fmt.Println(slices.IsSorted([]int{1, 55, 2}))
}
