package main

import "fmt"

// MyInt 自定义类型
type MyInt int

func main() {
	/*
		即使底层类型相同（int），Go 也要求显式转换
		因为它们是不同的命名类型
	*/
	var a MyInt = 10
	var b int = int(a)     // 显示转换
	var c MyInt = MyInt(b) // 反向转换
	fmt.Printf("%v: %T\n", a, a)
	fmt.Printf("%v: %T\n", b, b)
	fmt.Printf("%v: %T\n", c, c)
}

/*
10: main.MyInt
10: int
10: main.MyInt
*/
