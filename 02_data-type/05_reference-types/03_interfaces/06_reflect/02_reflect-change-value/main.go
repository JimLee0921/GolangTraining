package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("=== change value ===")
	x := 3.5
	v := reflect.ValueOf(&x) // 传入指针
	/*
		Elem() 返回接口 v 包含的值或指针 v 指向的值
		如果 v 的 Kind 不是Interface或Pointer
		则会发生混乱。如果 v 为 nil ，则返回零值
	*/
	v = v.Elem() // 取出指向的值
	/*
		CanSet 返回 v 的值是否可以更改
		如果 CanSet 返回 false 进行修改会导致 panic
	*/
	if v.CanSet() {
		v.SetFloat(777)
	}
	fmt.Println(x) // 777
}
