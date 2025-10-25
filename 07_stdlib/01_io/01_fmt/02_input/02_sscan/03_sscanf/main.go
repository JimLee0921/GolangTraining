package main

import "fmt"

func main() {
	var name string
	var age int
	input := "Name=Jim Age=12"

	n, err := fmt.Sscanf(input, "Name=%s Age=%d", &name, &age)

	fmt.Println("read params count: ", n, "err:", err)
	fmt.Println(name, age)
}
