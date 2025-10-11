package main

import "fmt"

func PrintAnything(v any) {
	fmt.Println(v)
}

func main() {
	PrintAnything(123)
	PrintAnything("hello")
	PrintAnything(true)
	PrintAnything([]int{1, 2, 3})
}
