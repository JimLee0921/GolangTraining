package main

import "fmt"

func main() {
	/*
		如果 接口内部类型不可比较会直接报错 panic
	*/
	var a interface{} = []int{1, 2, 3}
	var b interface{} = []int{1, 2, 3}
	fmt.Println(a == b) // panic: runtime error: comparing uncomparable type []int
}
