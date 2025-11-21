package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}
type B interface {
}

func TypeOfDemo() {
	fmt.Println(reflect.TypeOf("Hello")) // string
	fmt.Println(reflect.TypeOf([]int{})) // []int
	fmt.Println(reflect.TypeOf(&User{})) // *main.User
	fmt.Println(reflect.TypeOf(User{}))  // main.User
	fmt.Println(reflect.TypeOf(nil))     // <nil>
}

func main() {
	TypeOfDemo()
}
