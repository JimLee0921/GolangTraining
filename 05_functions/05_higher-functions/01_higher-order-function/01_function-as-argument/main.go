package main

import "fmt"

// doTwice 函数作为参数
func doTwice(fn func()) {
	fn()
	fn()
}

func operate(fn func(int, int) int, a, b int) int {
	return fn(a, b)
}

func add(x, y int) int { return x + y }
func mul(x, y int) int { return x * y }
func sayHello() {
	fmt.Println("Hello!")

}

func main() {
	doTwice(sayHello)
	fmt.Println(operate(add, 2, 3)) // 5
	fmt.Println(operate(mul, 2, 3)) // 6
}
