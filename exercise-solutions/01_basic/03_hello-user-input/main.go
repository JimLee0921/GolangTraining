package main

import "fmt"

func main() {
	var name string
	fmt.Println("please enter your name: ")
	// 不能有空格，只读取空格前的
	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}
	fmt.Println("Hello!", name)
}
