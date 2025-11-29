package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Base struct{ ID int }
type User struct {
	Base // 匿名字段
	Name string
}

func main() {
	t := reflect.TypeOf(User{})
	fmt.Println(t.NumField()) // 2 获取字段数量

	f1 := t.Field(0) // 按下标获取字段信息，使用 StructField 机哪 StructField-demo
	fmt.Println(f1)

	f2, ok := t.FieldByName("NoField") // 找不到返回 false
	if ok {
		fmt.Println(f2)
	} else {
		fmt.Println("not found NoField")
	}

	f3, ok := t.FieldByNameFunc(func(name string) bool {
		return strings.HasPrefix(name, "ID") // 查找以 ID 为开头的字段
	})
	if ok {
		fmt.Println(f3)
	} else {
		fmt.Println("not found IDXXX field")
	}

	fmt.Println(f3.Index)          // [0 0]
	f4 := t.FieldByIndex(f3.Index) // 用于匿名嵌套字段，传入 []int，可以递归地访问深层 struct 字段
	fmt.Println(f4.Name)           // ID
}
