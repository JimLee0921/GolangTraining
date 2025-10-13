package main

import "fmt"

func resetSlice(s []int) {
	// s = append(s, 100) 会让 s 指向一个新的底层数组，而原切片 main 中的 s 不会被影响。
	s = append(s, 100)
	fmt.Println("inside:", s)
}

func main() {
	s := []int{1, 2, 3}
	resetSlice(s)
	fmt.Println("outside:", s)
	/*
		inside: [1 2 3 100]
		outside: [1 2 3]
	*/
}
