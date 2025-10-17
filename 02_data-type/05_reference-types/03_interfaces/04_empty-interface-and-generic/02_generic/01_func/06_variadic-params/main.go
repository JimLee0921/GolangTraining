package main

import "fmt"

// Join 的不定参数泛型合并为切片后返回
func Join[T any](elems ...T) []T {
	return elems
}

func main() {
	fmt.Println(Join(1, 2, 3))
	fmt.Println(Join("a", "b", "c"))
}
