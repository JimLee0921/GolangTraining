package main

import "fmt"

// main Go中的类型转换
func main() {
	/*
		数值类型转换（最常见）
		Go 不支持隐式类型转换，必须显式写出
		var b float64 = a  会编译错误
	*/
	var a int = 10
	var b float64 = float64(a) // int -> float64
	var c int64 = int64(a)     // int -> int64

	var f float64 = 3.9
	var i int = int(f) // float64 → int（截断，不四舍五入）
	fmt.Printf("%v: %T\n", b, b)
	fmt.Printf("a: %v type: %T\n", a, a)
	fmt.Printf("a: %v type: %T\n", b, b)
	fmt.Printf("a: %v type: %T\n", c, c)
	fmt.Printf("a: %v type: %T\n", f, f)
	fmt.Printf("a: %v type: %T\n", i, i)
}

/*
10: float64
a: 10 type: int
a: 10 type: float64
a: 10 type: int64
a: 3.9 type: float64
a: 3 type: int
*/
