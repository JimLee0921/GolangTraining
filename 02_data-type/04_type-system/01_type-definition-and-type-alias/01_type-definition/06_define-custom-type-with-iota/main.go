package main

import "fmt"

type Status int

/*
定义了一个枚举型常量组
每个常量是 Status 类型
这是一种 Go 中最常见的类型 + iota组合使用方式
*/
const (
	Pending Status = iota
	Running
	Complete
	Failed
)

func main() {
	fmt.Println(Pending)
	fmt.Println(Running)
	fmt.Println(Complete)
	fmt.Println(Failed)
}
