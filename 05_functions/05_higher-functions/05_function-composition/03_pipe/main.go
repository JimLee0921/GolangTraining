package main

import "fmt"

func Pipe[T any](funcs ...func(T) T) func(T) T {
	return func(x T) T {
		result := x
		for _, fn := range funcs { // 正向执行
			result = fn(result)
		}
		return result
	}
}
func add1(x int) int   { return x + 1 }
func double(x int) int { return x * 2 }
func square(x int) int { return x * x }
func main() {
	/*
		Compose: 从右到左执行（数学函数风格）
		Pipe: 从左到右执行（数据流风格）
	*/
	pipeline := Pipe(add1, double, square)
	fmt.Println(pipeline(2)) // ((2+1)*2)^2 = 36

}
