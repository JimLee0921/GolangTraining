package main

import "fmt"

// main 类型断言
func main() {
	/*
		GO 中的类型断言 (type assertion) 用于从 接口类型变量 中取出它的 动态值
		在 Go 里，接口变量可以存储任何实现了接口的值，但是当需要拿出里面的具体类型时，就需要使用类型断言
		v, ok := x.(T)
			x：必须是接口类型的变量（如 interface{}(简写为 any)、自定义接口）
			T：目标类型
			v：如果断言成功，就是 x 存的那个值，并且类型是 T
			ok：布尔值，表示断言是否成功
				带ok:
					如果断言成功 -> v 是类型 T 的值，ok = true
					如果断言失败 -> v 是 T 的零值，ok = false
				不带 ok
					如果 x 不是接口类型，或者接口里存的动态值不是 T，直接 panic
					不安全，除非能百分百确定类型
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
