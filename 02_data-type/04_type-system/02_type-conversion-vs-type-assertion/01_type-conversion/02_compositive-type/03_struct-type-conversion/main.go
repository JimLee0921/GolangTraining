package main

import (
	"fmt"
)

type A struct {
	X int
	Y string
}

type B struct {
	X int
	Y string
}
type C struct {
	Y string
	X int
} // 字段顺序不同

func main() {
	/*
		结构体的字段顺序、名称、类型 必须完全一致才可转换
		若有任何差异（字段顺序、名称、数量、类型），就不可转换
	*/
	a := A{1, "Hi"}
	b := B(a)         // 字段完全相同。可以转换
	fmt.Println(a, b) // {1 Hi} {1 Hi}
	// c := C(a) // 不允许
}
