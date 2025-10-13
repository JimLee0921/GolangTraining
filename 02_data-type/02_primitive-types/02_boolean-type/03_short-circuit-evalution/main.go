package main

import "fmt"

func A() bool {
	fmt.Println("A() called")
	return false
}

func B() bool {
	fmt.Println("B() called")
	return true
}

func main() {
	fmt.Println("Case 1: A() && B()")
	// A() 为 false，&& 短路，B() 不会被调用
	if A() && B() {
		fmt.Println("A && B -> true")
	} else {
		fmt.Println("A && B -> false")
	}

	fmt.Println("\nCase 2: B() || A()")
	// B() 为 true，|| 短路，A() 不会被调用
	if B() || A() {
		fmt.Println("B || A -> true")
	} else {
		fmt.Println("B || A -> false")
	}
}
