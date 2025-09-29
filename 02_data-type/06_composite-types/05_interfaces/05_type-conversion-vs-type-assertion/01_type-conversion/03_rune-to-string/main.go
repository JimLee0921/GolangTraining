package main

import "fmt"

// main rune和string的转换
func main() {
	/*
		rune 的本质就是 int32 存放的是 Unicode 码点
		单引号式字符字面量 双引号式字符串字面量
	*/
	var x rune = 'b' // 等同于 var x runt = 98
	var y int = 'a'  // 本质存的还是字符 a 的 Unicode 码点

	fmt.Println(x)         // 98
	fmt.Println(string(x)) // "b"
	fmt.Println(y)         // 97
	fmt.Println(string(y)) // "a" 这里是不安全的 这里是将码点转为字符 a
}
