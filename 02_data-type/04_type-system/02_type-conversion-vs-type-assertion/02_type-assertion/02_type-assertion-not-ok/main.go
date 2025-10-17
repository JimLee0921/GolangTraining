package main

import "fmt"

// mian 类型断言失败
func main() {
	var name any = 666
	str, ok := name.(string)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("%q value is not a string\n", str) // 断言失败返回断言类型的空值
	}

	var val any = 7
	fmt.Printf("%v: %T\n", val, val)
	fmt.Printf("%v: %T\n", val.(int), val.(int))
	//fmt.Println(val + 3)

}
