package main

import "fmt"

func main() {
	/*
		指针类型包含：
			指向的类型
			如果底层类型一致，可以互转
	*/

	type MyInt int
	var x MyInt = 5
	var p *MyInt = &x
	var q *int = (*int)(p) // 合法
	fmt.Println(q)
	// 但不同底层类型则 不行。
}
