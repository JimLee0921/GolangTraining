package main

import "fmt"

// Box 通用容器结构体，可以使得 value 设置为任意类型
type Box[T any] struct {
	value T
}

func main() {
	intBox := Box[int]{value: 10}
	strBox := Box[string]{value: "hello"}

	fmt.Println(intBox.value) // 10
	fmt.Println(strBox.value) // hello
}
