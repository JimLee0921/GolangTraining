package main

import "fmt"

func main() {
	// 遇到换行符立即停止读取（不会跨行读取）
	var name string
	var age int

	fmt.Print("enter your name and age(Use spaces to separate): ")
	n, err := fmt.Scanln(&name, &age)
	fmt.Println("successful read count: ", n, "err: ", err)
	fmt.Printf("name: %s, age: %d", name, age)
}
