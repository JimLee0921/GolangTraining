package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("=== struct field reflect ===")
	u := User{
		Name: "Jim",
		Age:  20,
	}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i)
		fmt.Printf("[%d] field=%s type=%s value=%v Tag(json)=%s\n", i, f.Name, f.Type, val.Interface(), f.Tag.Get("json"))
	}
}
