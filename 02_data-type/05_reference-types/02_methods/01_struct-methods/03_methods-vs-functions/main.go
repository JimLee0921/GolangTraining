package main

import "fmt"

type User struct {
	Name string
	Age  int
}

/*
方法语法更自然、更接近面向对象
方法可以直接参与接口实现
普通函数不能挂载到类型上
*/
func growUpFunction(u *User) {
	u.Age++
}

func (u *User) growUpMethod() {
	u.Age++
}

func main() {
	/*
		1. 结构体方法和函数调用方式不同
		2. 结构体方法自动传址相当于 (&u).growUpMethod()
	*/
	u := User{
		Name: "Jim",
		Age:  20,
	}
	growUpFunction(&u)
	fmt.Println(u) // {Jim 21}
	u.growUpMethod()
	fmt.Println(u) // {Jim 22}
}
