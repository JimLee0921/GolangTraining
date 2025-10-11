package main

import "fmt"

// Stringable 自定义约束，必须实现了 String() 方法
type Stringable interface {
	String() string
}

func PrintString[T Stringable](v T) {
	fmt.Println(v.String())
}

type User struct {
	Name string
}

func (u User) String() string {
	return "User: " + u.Name
}

func main() {
	/*
		约束可以是接口行为
		只要类型实现了该方法，就可以作为泛型参数
		类似传统接口的多态，但在编译期决定
	*/
	PrintString(User{Name: "Tom"})
}
