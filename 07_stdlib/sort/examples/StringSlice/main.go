package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"SKU10", "SKU20", "SKU1"}

	s := sort.StringSlice(strs)
	s.Sort()

	fmt.Println(strs)
	idx := s.Search("SKU1")
	fmt.Println(idx)
}
