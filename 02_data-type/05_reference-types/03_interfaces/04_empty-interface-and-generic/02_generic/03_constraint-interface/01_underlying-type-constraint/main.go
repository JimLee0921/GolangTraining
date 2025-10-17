package main

import "fmt"

// MyInt 自定义类型
type MyInt int

// Integer 运行所有底层类型为 int 的类型
type Integer interface {
	~int
}

func Double[T Integer](v T) T {
	return v * 2
}

func main() {
	/*
		~int 表示底层类型是 int
		允许别名类型（MyInt）也能使用
		泛型在定义通用库类型（例如数据库 ID、单位值类型）时非常有用
	*/
	fmt.Println(Double(10))        // int
	fmt.Println(Double(MyInt(10))) // MyInt（底层是 int）
}
