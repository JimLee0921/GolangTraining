package main

import "fmt"

func main() {
	/*
		更多类型断言见 type-system/type-conversion-vs-type-assertion 章节
		如果类型断言失败会直接 panic
	*/
	var v interface{} = "123"

	n := v.(int)       // 类型断言
	fmt.Println(n + 1) // 124
}
