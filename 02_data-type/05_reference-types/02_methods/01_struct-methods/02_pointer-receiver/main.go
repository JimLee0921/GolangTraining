package main

import "fmt"

type User struct {
	Name      string
	Telephone string
}

/*
想在方法内部修改对象的字段，就要用指针接收者
u *User 表示接收者是指针
调用时可以直接用 u.GrowUp()
Go 会自动取地址（即 (&u).GrowUp()）
修改会作用于原始变量
*/
func (u *User) changeTelephone(phoneNumber string) {
	u.Telephone = phoneNumber
}

func main() {

	u := User{
		Name: "JimLee",
	}
	u.changeTelephone("1999999999")
	fmt.Println(u)
}
