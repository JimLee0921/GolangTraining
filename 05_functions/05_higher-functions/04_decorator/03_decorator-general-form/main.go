package main

import "fmt"

// Decorate 一个 通用（泛型）装饰器，在不修改原函数逻辑的前提下，给它前置（before）和后置（after）其它逻辑
func Decorate[T any](fn func(T) T, before func(T), after func(T, T)) func(T) T {
	return func(input T) T {
		before(input)
		output := fn(input)
		after(input, output)
		return output
	}
}

func main() {
	addOne := func(x int) int { return x + 1 }
	logBefore := func(x int) { fmt.Println("Input:", x) }
	logAfter := func(x, y int) { fmt.Println("Output:", y) }

	decorated := Decorate(addOne, logBefore, logAfter)
	decorated(10)
}
