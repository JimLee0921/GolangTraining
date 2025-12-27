package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func PrintStructInfo(s any) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if t.Kind() != reflect.Struct {
		fmt.Println("not struct type")
		return
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("field: %-10s type: %-10s value: %-10v\n", field.Name, field.Type, value.Interface())
	}
}

func main() {
	PrintStructInfo(User{"Alice", 25})
}
