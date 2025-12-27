package main

import (
	"fmt"
	"slices"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[:4:10]
	fmt.Println(cap(s))
	clip := slices.Clip(s)
	fmt.Println(clip)
	fmt.Println(len(clip))
	fmt.Println(cap(clip))
}
