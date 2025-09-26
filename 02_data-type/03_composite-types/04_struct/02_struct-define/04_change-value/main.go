package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	person := person{
		firstName: "Jim",
		lastName:  "Lee",
		age:       20,
	}
	fmt.Println(person)
	// 直接通过 .属性名 访问即可修改属性值（指针修改见struct-pointer）
	person.age = 10
	fmt.Println(person)
}
