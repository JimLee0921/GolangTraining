package main

import (
	"fmt"
	"reflect"
)

type Speaker interface {
	Speak()
}

type Dog struct{ Name string }

func (d Dog) Speak() { fmt.Println("Woof!") }

func main() {
	var s Speaker // 静态类型 Speaker
	d := Dog{"Buddy"}
	s = d

	fmt.Println(reflect.TypeOf(s)) // main.Dog	s 的 动态类型 是 Dog
	fmt.Println(s)                 // {Buddy}	动态值 是一个 Dog 实例 {Buddy}

	// 空接口
	/*
		动态类型是 *int
		动态值是 nil
		所以接口整体不等于 nil
	*/
	var x any = (*int)(nil)
	fmt.Println(x == nil)
	fmt.Printf("Type: %v, Value: %v\n", reflect.TypeOf(x), reflect.ValueOf(x))
}
