package main

import (
	"fmt"
)

func main() {
	/*
		多个参数定义：参数名 参数类型, 参数名2 参数类型2, ...
		相邻的相同类型参数可以合并书写
		参数名, 参数名2, 参数类型, 参数名3 参数类型2, ...
	*/
	introduce("JimLee", 20, "CA")
	greeting("JimLee", "man", 20)
}

func introduce(name string, age int, city string) {
	fmt.Printf("hello! i'm %s and i'm %d years old. my city is %s\n", name, age, city)
}

func greeting(name, sex string, age int) {
	fmt.Println("i am", name, "i am", age, "years old, i am a", sex)
}
