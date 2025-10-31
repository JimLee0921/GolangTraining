package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 实现 fmt.Stringer 接口
func (u User) String() string {
	return fmt.Sprintf("user[%s], age[%d]", u.Name, u.Age)
}

func main() {
	user := User{
		Name: "JimLee",
		Age:  17,
	}
	fmt.Println(user)         // 自动调用 String()
	fmt.Printf("%v\n", user)  // 自动调用 String()
	fmt.Printf("%#v\n", user) // %#v 会打印 Go 源码表示，不会用 String()
	// 注意最后一行：%#v 打印的是结构体源码格式，不会调用 String()
}
