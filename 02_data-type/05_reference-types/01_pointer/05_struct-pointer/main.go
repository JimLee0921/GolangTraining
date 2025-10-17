package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) GrowUp() {
	u.Age++
}

func main() {

	u := User{
		Name: "Jim",
		Age:  22,
	}
	u.GrowUp()         // 自动取地址：编译器自动转换为 (&u).GrowUp()
	fmt.Println(u.Age) // 23
}
