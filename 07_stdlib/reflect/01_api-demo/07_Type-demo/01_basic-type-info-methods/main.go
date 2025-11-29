package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int64
}

func inspect(x any) {
	t := reflect.TypeOf(x)

	fmt.Println("Type Name: ", t.Name())
	fmt.Println("Kind: ", t.Kind())
	fmt.Println("Pkg Path: ", t.PkgPath())
	fmt.Println("String: ", t.String())
	fmt.Println("Size: ", t.Size())
	fmt.Println("Align: ", t.Align())
	fmt.Println("FieldAlign: ", t.FieldAlign())
	fmt.Println("------------------------------")
}

func main() {
	inspect(10)
	inspect([]string{})
	inspect(User{})
	inspect(&User{})
}
