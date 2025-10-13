package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// NewUser 构造函数（工厂函数）
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	u := NewUser("Tom", 20)
	fmt.Println(u) // &{Tom 20}
}
