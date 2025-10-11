package main

import "fmt"

func resetSlice(s []int) []int {
	// 用返回值写回
	s = append(s, 100)
	return s
}

func main() {
	s := []int{1, 2, 3}
	// 重新接收
	s = resetSlice(s)
	fmt.Println(s) // [1 2 3 100]
}
