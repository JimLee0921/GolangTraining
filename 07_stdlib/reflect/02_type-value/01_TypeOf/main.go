package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// TypeOf 所有的 Type / Value API 基本都需要 reflect.TypeOf 作为入口
	fmt.Println(reflect.TypeOf(123))              // int
	fmt.Println(reflect.TypeOf("hello"))          // string
	fmt.Println(reflect.TypeOf([]int{}))          // []int
	fmt.Println(reflect.TypeOf(map[string]int{})) // map[string]int
	fmt.Println(reflect.TypeOf(User{}))           // main.User
	fmt.Println(reflect.TypeOf(&User{}))          // *main.User
}
