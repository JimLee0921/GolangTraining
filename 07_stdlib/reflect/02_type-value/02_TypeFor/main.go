package main

import (
	"fmt"
	"reflect"
)

type User struct {
}

func main() {
	// TypeFor 不需要跟 TypeOf 一样必须传入一个值
	fmt.Println(reflect.TypeFor[int]())
	fmt.Println(reflect.TypeFor[*string]())
	fmt.Println(reflect.TypeFor[[]int]())
	fmt.Println(reflect.TypeFor[map[string]User]())
}
