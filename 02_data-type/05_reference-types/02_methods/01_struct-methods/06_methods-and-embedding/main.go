package main

import "fmt"

type Human struct {
	Name string
}

type User struct {
	Human
}

func (h *Human) Hello() {
	fmt.Println("Hello, im", h.Name)
}

func (h *Human) Eating() {
	fmt.Println(h.Name, "is eating...")
}

func (u *User) Eating() {
	fmt.Println(u.Name, "is eating meat...")
}

func main() {
	/*
		如果一个结构体嵌入另一个结构体
		它会自动继承嵌入类型的方法
		如果这个结构体有自己的方法导致方法遮蔽，可以手动指定方法
	*/
	u := User{Human{Name: "JamesBond"}}
	u.Hello()        // Hello, im JamesBond
	u.Eating()       // `JamesBond is eating meat...` 默认是自己的 eating 方法
	u.Human.Eating() // JamesBond is eating...
}
