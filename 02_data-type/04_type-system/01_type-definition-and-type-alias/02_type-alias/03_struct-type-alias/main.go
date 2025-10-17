package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Member = User // Member 只是 User 的别名

func main() {
	/*
		Member 与 User 在编译器内部完全是同一类型
		它们可以互相赋值、传参、比较
	*/
	u := User{"Tom", 20}
	m := Member{"Jerry", 22}

	fmt.Printf("%T\n", u) // 输出：main.User
	fmt.Printf("%T\n", m) // 输出：main.User（不是 main.Member）
}
