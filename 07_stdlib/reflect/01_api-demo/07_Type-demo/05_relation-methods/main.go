package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func ImplementsDemo() {
	var w io.Writer
	tWriter := reflect.TypeOf(&w).Elem() // io.Writer
	t := reflect.TypeOf(os.Stdout)
	fmt.Println(t.Implements(tWriter)) // true *os.File 实现了 io.Writer
}

func AssignableToDemo() {
	type A struct {
		X int
	}
	type B A
	tA := reflect.TypeOf(A{})
	tB := reflect.TypeOf(B{})
	// tB 底层虽然是 A 但是类型不同，不能直接赋值
	fmt.Println(tB.AssignableTo(tA)) // false
	fmt.Println(tA.AssignableTo(tB)) // false
}

func ConvertibleToDemo() {
	type MyInt int
	tMy := reflect.TypeOf(MyInt(1))
	tInt := reflect.TypeOf(10)
	// 是否可以进行转换
	fmt.Println(tMy.ConvertibleTo(tInt)) // true
	fmt.Println(tInt.ConvertibleTo(tMy)) // true
}

func ComparableDemo() {
	fmt.Println(reflect.TypeOf(1).Comparable())                 // true
	fmt.Println(reflect.TypeOf("x").Comparable())               // true
	fmt.Println(reflect.TypeOf([]int{}).Comparable())           // false，slice 不可比较
	fmt.Println(reflect.TypeOf(map[string]int{}).Comparable())  // false
	fmt.Println(reflect.TypeOf(struct{ X int }{}).Comparable()) // true（如果全部字段可比较）
}

func main() {
	ImplementsDemo()
	AssignableToDemo()
	ConvertibleToDemo()
	ComparableDemo()
}
