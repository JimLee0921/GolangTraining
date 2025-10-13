package main

import "fmt"

func greeter() func(string) string {
	return func(name string) string {
		return "Hello, " + name
	}
}

func main() {
	/*
		greeter 返回一个匿名函数
		生成的函数可以带着外部逻辑返回
		是构建函数工厂（factory）的基础
	*/
	greet := greeter()
	fmt.Println(greet("Tom"))
}
