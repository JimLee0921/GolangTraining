package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	/*
		常用于结构体实例初始化等操作
	*/
	user := newUser("Jim", 20)
	fmt.Println(user) // &{Jim 20}
	user.Name = "dsb"
	fmt.Println(user)
	fmt.Println(*user)
	/*
		&{Jim 20}
		&{dsb 20}
		{dsb 20}
	*/
}

func newUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}
