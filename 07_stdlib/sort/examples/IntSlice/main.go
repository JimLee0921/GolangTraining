package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 1, 2, 7, 22, 1, 23}
	s := sort.IntSlice(nums)
	s.Sort()
	fmt.Println(nums)

	idx := s.Search(5)
	fmt.Println(idx)
}
