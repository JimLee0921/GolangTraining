package main

import "fmt"

func main() {
	/*
		接口比较
			只有当两个接口的 类型和值都相等 时，比较结果才为 true
			如果接口和 nil 比较，只有在 (type == nil && value == nil) 时才为 true
			如果接口中值类型不支持比较（例如 map、slice），直接 panic
	*/
	/*
		同类型同值
			a 和 b 的动态类型都是 int
			动态值也都是 10
			所以相等
	*/
	var a any = 10
	var b any = 10
	fmt.Println("a == b", a == b) // a == b true

	// 同类型不同值
	var x interface{} = 10
	var y interface{} = 20

	fmt.Println("x == y", x == y) // false

	// 不同类型同值
	var m interface{} = int(10)
	var n interface{} = int64(10)

	fmt.Println("m == n", m == n) // false
}
