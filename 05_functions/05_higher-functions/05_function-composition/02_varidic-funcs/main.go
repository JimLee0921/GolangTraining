package main

import "fmt"

func Compose[T any](funcs ...func(T) T) func(T) T {
	return func(x T) T {
		result := x
		for i := len(funcs) - 1; i >= 0; i-- { // 从右到左执行
			result = funcs[i](result)
		}
		return result
	}
}
func add1(x int) int   { return x + 1 }
func double(x int) int { return x * 2 }
func square(x int) int { return x * x }

func main() {
	/*
		Compose 从右向左执行（数学函数式风格）
		多个函数被自动串联
		代码更可读、更声明化
	*/
	pipeline := Compose(add1, double, square)
	fmt.Println(pipeline(2)) // 先 square(2)=4 → double(4)=8 → add1(8)=9
}
