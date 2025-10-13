package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	setFirstByValue(arr)
	fmt.Println(arr) // [1 2 3]
	setFirstByPointer(&arr)
	fmt.Println(arr) // [99 2 3]
}

func setFirstByPointer(a *[3]int) {
	a[0] = 99
}

func setFirstByValue(a [3]int) {
	a[0] = 99
}
