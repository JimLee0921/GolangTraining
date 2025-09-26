package main

import "fmt"

func main() {
	/*
		调用 greeting 方法，需要传入一个字符串参数，greeting没有返回值
	*/
	greeting("JimLee")
	greeting("JamesBond")
}

func greeting(name string) {
	fmt.Println("hello! ", name)
}
