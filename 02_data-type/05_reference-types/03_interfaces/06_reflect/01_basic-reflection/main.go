package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("=== basic reflection ===")
	var x float64 = 3.5

	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())

	fmt.Println("Value:", v)
	fmt.Println("Value.Float():", v.Float())

}
