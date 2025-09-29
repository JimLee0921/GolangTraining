package main

import "fmt"

func main() {
	/*
		类型断言：x.(T) —— 只对接口变量生效，用来取出里面的具体值
		类型转换：T(v) —— 在具体类型之间转换，不能把接口直接转成具体类型
	*/
	rem := 7.24
	fmt.Printf("%T\n", rem)      // float64
	fmt.Printf("%T\n", int(rem)) // int

	var val interface{} = 7
	fmt.Printf("%T\n", val) // int
	//fmt.Printf("%T\n", int(val)) // 错误：不能直接转换
	fmt.Printf("%T\n", val.(int)) // 正确：类型断言
}
