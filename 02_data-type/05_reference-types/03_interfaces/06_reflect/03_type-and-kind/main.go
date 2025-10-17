package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("=== type and kind ===")

	arr := []int{1, 2, 3}
	t := reflect.TypeOf(arr)

	fmt.Println("Type:", t)        // Type: []int
	fmt.Println("Kind:", t.Kind()) // Kind: slice

}
