package main

import "fmt"

func main() {
	/*
		Go 不会自动进行类型提升或降级
		必须使用 T(x) 明确转换类型
	*/
	var a int16 = 100
	var b int32 = 200

	//fmt.Println(a + b)        //  编译错误：类型不匹配
	fmt.Println(int32(a) + b) //  显式转换
}
