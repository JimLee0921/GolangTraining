package main

import "fmt"

func compose[A any, B any, C any](
	f func(B) C,
	g func(A) B,
) func(A) C {
	return func(x A) C {
		return f(g(x))
	}
}

func add1(x int) int   { return x + 1 }
func square(x int) int { return x * x }

func main() {
	/*
		compose(f, g) 返回一个新函数
		执行顺序：g -> f
		最终结果与 f(g(x)) 等价
	*/
	add1ThenSquare := compose(square, add1)
	fmt.Println(add1ThenSquare(2)) // (2 + 1)^2 = 9
}
