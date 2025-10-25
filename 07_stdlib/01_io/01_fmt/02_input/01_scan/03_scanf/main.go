package main

import "fmt"

func main() {
	// 遇到换行符立即停止读取（不会跨行读取）
	var name string
	var age int

	fmt.Print("enter your name and age(Name=Tom Age=18): ")
	// 输入格式必须与模板严格匹配（包括空格、符号）
	n, err := fmt.Scanf("Name=%s Age=%d", &name, &age)
	fmt.Println("read count:", n, "err:", err)
	fmt.Printf("name: %s, age: %d", name, age)
}
