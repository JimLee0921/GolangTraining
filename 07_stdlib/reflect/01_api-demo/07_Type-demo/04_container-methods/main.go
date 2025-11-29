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
	fmt.Println(reflect.TypeOf(&User{}).Kind())        // ptr
	fmt.Println(reflect.TypeOf(&User{}).Elem().Kind()) // struct

	fmt.Println(reflect.TypeOf([]string{}).Elem()) // string

	fmt.Println(reflect.TypeOf(map[string]int{}).Elem()) // int
	fmt.Println(reflect.TypeOf(map[string]int{}).Key())  // string

	fmt.Println(reflect.TypeOf([10]int{}).Len()) // 10

	fmt.Println(reflect.TypeOf(make(chan<- int)).Kind())    // chan
	fmt.Println(reflect.TypeOf(make(chan<- int)).ChanDir()) // chan<-
	fmt.Println(reflect.TypeOf(make(chan<- int)).Elem())    // int
}
