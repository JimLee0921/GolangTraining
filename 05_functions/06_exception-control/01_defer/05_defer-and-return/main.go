package main

import (
	"fmt"
)

func f() (result int) {
	defer func() {
		result++
	}()
	return 10
}

func f2() int {
	x := 10
	defer func() {
		x++
	}()
	return x
}
func main() {
	fmt.Println(f())  // 11
	fmt.Println(f2()) // 10
	/*
		return 10 先把返回值 result = 10
		调用 defer
		defer 内把 result 加 1
		最终返回 11
		注意：如果返回值不是命名返回值（例如 return 10 没声明 result），defer 修改不了它
	*/
}
