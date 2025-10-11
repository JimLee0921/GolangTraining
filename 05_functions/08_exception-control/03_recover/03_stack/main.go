package main

import "fmt"

func main() {
	/*
		B 触发 panic
		B 的 defer 执行
		panic 向上传递到 A
		A 的 defer 执行
		传到 main 的 defer
		main 的 defer 调用 recover → 捕获，停止传播
	*/
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("main recovered:", r)
		}
	}()
	A()
	/*
		B defer
		A defer
		main recovered: something wrong in B
	*/
}
func A() {
	defer fmt.Println("A defer")
	B()
}

func B() {
	defer fmt.Println("B defer")
	panic("something wrong in B")
}
