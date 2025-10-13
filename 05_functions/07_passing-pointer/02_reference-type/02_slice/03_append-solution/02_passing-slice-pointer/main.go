package main

import "fmt"

func resetSlice(s *[]int) {
	// 传递指针使用指针进行扩容
	*s = append(*s, 100)
}

func main() {
	// 传递 *[]int（指针) 可以自动扩容
	s := []int{1, 2, 3}
	resetSlice(&s)
	fmt.Println(s) // [1 2 3 100]
}
