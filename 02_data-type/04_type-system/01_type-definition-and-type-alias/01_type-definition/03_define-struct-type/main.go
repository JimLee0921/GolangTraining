package main

import "fmt"

// User 结构体也是通过 type 定义的
type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	u := User{
		ID:   1,
		Name: "JimLee",
		Age:  20,
	}
	fmt.Println(u)
}
