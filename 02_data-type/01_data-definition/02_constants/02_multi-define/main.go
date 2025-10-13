package main

import "fmt"

const (
	Pi       = 3.15
	language = "GO"
)

// main 一次定义多个常量
func main() {
	/*
		const(...)被称为常量声明块
	*/
	fmt.Println(Pi, language)

}
