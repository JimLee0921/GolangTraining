package main

import "fmt"

func main() {
	res := Equal("123", "123")
	if res {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

// Equal comparable 约束限定类型必须满足 == 和 != 操作符
func Equal[T comparable](a, b T) bool {
	return a == b
}
