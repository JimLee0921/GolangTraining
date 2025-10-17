package main

import "fmt"

func describe(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("string: %q, length %d\n", v, len(v))
	case int:
		fmt.Printf("int: %d, double %d\n", v, v*2)
	default:
		fmt.Printf("unknown type %T\n", v)
	}
}

func main() {
	/*
		普通类型断言：s, ok := i.(string)
			每次只能判断一个类型，要写多个就得连着写好几个 if
		类型 switch：
			switch v := i.(type) {
			case string:
				fmt.Println("string:", v)
			case int:
				fmt.Println("int:", v)
			default:
				fmt.Println("unknown")
			}
		一次性把不同类型的情况都列出来，更直观
		注意：
			1. i 必须是接口类型（例如 interface{}）
			2. i.(type) 只能用在 switch 里
			3. 每个 case 后面写的是一个类型，不是值
			4. 进入某个 case 分支时，变量 v 的类型就是该分支里的类型
	*/
	describe("hello")
	describe(42)
	describe(3.14)
}
