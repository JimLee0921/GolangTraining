package main

import "fmt"

// 自定义 type
type contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	/*
		x.(type) 只能用于 类型 switch
		它会检查 x 在运行时是什么类型
		前提是 x 是 接口类型（这里是 interface{}，表示可以接受任何类型）
	*/
	switch x.(type) {
	// 根据 x 的实际类型进入不同分支
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

// main switch还可以进行类型判断（接口断言）
func main() {
	SwitchOnType(7)                            // 传入 int -> 打印 "int"
	SwitchOnType("McLeod")                     // 传入 string -> 打印 "string"
	var t = contact{"Good to see you,", "Tim"} // 结构体实例化
	SwitchOnType(t)                            // 传入 contact -> 打印 "contact"
	SwitchOnType(t.greeting)                   // t.greeting 是 string -> 打印 "string"
	SwitchOnType(t.name)                       // t.name 也是 string -> 打印 "string"

}
