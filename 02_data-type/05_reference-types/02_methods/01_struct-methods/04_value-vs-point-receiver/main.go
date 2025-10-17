package main

import "fmt"

type User struct {
	Name string
}

func (u User) RenameByValue(newName string) {
	u.Name = newName
}

func (u *User) RenameByPointer(newName string) {
	u.Name = newName
}

func main() {
	u := User{"Jim"}
	u.RenameByValue("Bruce")   // // 值接收者方法
	fmt.Println(u.Name)        // Jim 修改失败
	u.RenameByPointer("James") // 自动取地址
	fmt.Println(u.Name)        // James 修改成功
}
