package main

import "fmt"

// User 定义结构体
type User struct {
	Name string
	Age  int
}

// Greet
/*
u User 称为值接收者（receiver）
让 Greet() 成为 User 类型的方法
所以可以通过 u.Greet() 调用
u 在方法内部是一个副本（值接收者）
通常不需要修改结构体数据时使用值接收者
*/
func (u User) Greet() {
	fmt.Println("Hi! my name is", u.Name)
}

func main() {

	u := User{Name: "Tom", Age: 30}
	u.Greet() // 调用方法
}
