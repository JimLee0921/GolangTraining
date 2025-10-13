package main

import "fmt"

func MakeAdder[T int | float64](base T) func(T) T {
	return func(t T) T {
		return t + base
	}
}

func main() {
	/*
		泛型让工厂函数支持多种数值类型
		编译期类型安全
		可生成整型或浮点型加法器
	*/

	intCal := MakeAdder(10)
	floatCal := MakeAdder(8.88)

	fmt.Println(intCal(20))
	fmt.Println(floatCal(88.2))
}
