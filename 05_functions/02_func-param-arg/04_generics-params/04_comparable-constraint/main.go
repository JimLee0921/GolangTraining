package main

import (
	"fmt"
	"math/rand"
)

func Equal[T comparable](a, b T) bool {
	return a == b
}

func main() {
	/*
		comparable 约束：类型必须支持 == 和 != 运算符
	*/
	a := rand.Intn(4)
	b := rand.Intn(4)
	if Equal(a, b) {
		fmt.Println("a == b")
	} else {
		fmt.Println("a != b")
	}
}
