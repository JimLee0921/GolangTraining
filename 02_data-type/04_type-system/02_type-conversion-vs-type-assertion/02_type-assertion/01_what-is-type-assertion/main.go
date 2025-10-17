package main

import "fmt"

// main 类型断言
func main() {
	/*

	 */
	//name := "Sydney" // name 为 string 而不是接口类型 不能进行断言，会报错
	var name any = "Simon" // 等价于 var name interface{} = "Simon"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
