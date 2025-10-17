package main

import "fmt"

type MyInt int

type Celsius float64

type Fahrenheit float64

func main() {
	/*
		上面定义的这些都是新类型
		和原始类型不同
		即使底层相同也不能直接互换
	*/
	var a MyInt = 20
	var b int = 10
	// b = a        // 编译错误：不同类型
	b = int(a)     // 必须显式转换
	fmt.Println(b) // 20
}
