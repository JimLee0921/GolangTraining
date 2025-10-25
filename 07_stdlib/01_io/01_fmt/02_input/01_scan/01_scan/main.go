package main

import "fmt"

func main() {
	// fmt.Scan 按空白分隔读取
	var name string
	var age int

	fmt.Print("enter your name and age(Use spaces to separate): ")
	n, err := fmt.Scan(&name, &age)
	fmt.Println("successful read count: ", n, "err: ", err)
	fmt.Printf("name: %s, age: %d", name, age)
}
