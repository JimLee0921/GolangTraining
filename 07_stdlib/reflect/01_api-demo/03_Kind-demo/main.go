package main

import (
	"fmt"
	"reflect"
)

type MyInt int

func main() {
	var a MyInt
	v := reflect.ValueOf(a)
	fmt.Println("Type: ", v.Type()) // Type:  main.MyInt
	fmt.Println("Kind: ", v.Kind()) // Kind:  int

	s := []string{"a", "b"}
	vs := reflect.ValueOf(s)
	fmt.Println("Slice Kind: ", vs.Kind()) // Slice Kind:  slice

	m := map[string]int{"x": 1}
	vm := reflect.ValueOf(m)
	fmt.Println("Map Kind: ", vm.Kind()) // Map Kind:  map
}
