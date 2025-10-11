package main

import "fmt"

func main() {
	/*
		可变参数
	*/
	greeting("JimLee")
	greeting("BruceLee", "JamesBond", "Skrillex")
	greeting() // 空传递

}

func greeting(names ...string) {
	fmt.Println("names 切片内容", names)
	for _, name := range names {
		fmt.Println("hello!", name)
	}

}
