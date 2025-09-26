package main

import "fmt"

// main 引用数据类型 slice 作为参数传递
func main() {
	nameSlice := []string{"Jim", "Ted", "James"}
	fmt.Println(nameSlice) // [Jim Ted James]
	changeSlice(nameSlice)
	fmt.Println(nameSlice) // [JimLee Ted James]
}

func changeSlice(s []string) {
	s[0] = "JimLee"
}
