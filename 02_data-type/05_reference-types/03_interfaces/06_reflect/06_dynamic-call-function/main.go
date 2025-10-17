package main

import (
	"fmt"
	"reflect"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	v := reflect.ValueOf(Add)
	args := []reflect.Value{
		reflect.ValueOf(2),
		reflect.ValueOf(3),
	}
	result := v.Call(args)
	fmt.Println(result[0].Int()) // 5
}
