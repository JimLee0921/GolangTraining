package main

import "fmt"

type Reader interface{ Read() }

type File struct{}

func (f *File) Read() {}

func main() {
	/*
		接口与 nil 比较时，Go 比较的还是 type+value
		只有 (type, value) == (nil, nil) 也就是类型和数值都为 nil 才返回 true
	*/
	var r Reader = nil    // 动态类型和动态值都为 nil 接口完全为空（真正的 nil 接口）
	fmt.Println(r == nil) // true

	var f *File = nil // 动态类型 *File 动态值 nil 接口里有类型信息，但值部分为 nil
	r = f
	fmt.Println(r == nil) // false
}
