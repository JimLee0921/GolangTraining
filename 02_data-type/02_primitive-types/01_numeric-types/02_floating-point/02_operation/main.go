package main

import "fmt"

func main() {
	/*
		不支持取模 %
		运算结果类型保持与操作数相同
		/ 永远执行浮点除法（不会截断）
	*/
	x := 10.0
	y := 3.0

	fmt.Println(x + y) // 加法 -> 13
	fmt.Println(x - y) // 减法 -> 7
	fmt.Println(x * y) // 乘法 -> 30
	fmt.Println(x / y) // 除法 -> 3.3333333333333335

	/*
		0.1 与 0.2 无法精确表示
		比较时不要直接用 ==
		应使用容差（epsilon）判断
	*/
	a, b := 0.1, 0.2
	fmt.Println(a+b == 0.3) // false
	fmt.Println(a + b)      // 0.30000000000000004}

	/*
		浮点与整数不能直接运算，需显式转换
	*/
	var m int = 3
	var n float64 = 2.5
	// fmt.Println(m + n) // 类型不匹配
	fmt.Println(float64(m) + n) // 显式转换

}
