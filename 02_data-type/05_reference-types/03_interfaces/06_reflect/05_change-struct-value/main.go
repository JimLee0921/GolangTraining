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
	fmt.Println("=== change struct fields ===")
	u := User{"Jim", 20}
	v := reflect.ValueOf(&u).Elem()

	nameField := v.FieldByName("Name")
	if nameField.CanSet() {
		nameField.SetString("Jerry")
	}

	ageField := v.FieldByName("Age")
	if ageField.CanSet() {
		ageField.SetInt(30)
	}

	fmt.Println("changed:", u) // {Jerry 30}
	fmt.Println()
}
