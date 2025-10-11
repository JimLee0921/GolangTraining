package main

import "fmt"

// SumOne 空接口写法
func SumOne(a, b any) any {
	switch a := a.(type) {
	case int:
		return a + b.(int)
	case float64:
		return a + b.(float64)
	default:
		panic("unsupported type")
	}
}

type Number interface {
	int | float64
}

// SumTwo 泛型写法
func SumTwo[T Number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(SumOne(3, 5))     // int
	fmt.Println(SumTwo(1.2, 3.4)) // float64

}
