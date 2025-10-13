package main

import (
	"fmt"
)

type Big struct {
	Data [1 << 20]byte // 1MB
}

// MakeBigPtr 返回指针：避免结构体拷贝
func MakeBigPtr() *Big {
	b := Big{}
	b.Data[0] = 42
	fmt.Printf("MakeBigPtr b addr: %p\n", &b)
	return &b
}

// MakeBigVal 返回值：会拷贝整个结构体（1MB）
func MakeBigVal() Big {
	b := Big{}
	b.Data[0] = 42
	fmt.Printf("MakeBigVal b addr: %p\n", &b)
	return b
}

func main() {
	// 返回指针版本
	p := MakeBigPtr()
	fmt.Printf("main p addr: %p\n", p)
	fmt.Println("p.Data[0] =", p.Data[0])

	fmt.Println("----------")

	// 返回值版本
	v := MakeBigVal()
	fmt.Printf("main v addr: %p\n", &v)
	fmt.Println("v.Data[0] =", v.Data[0])
}

/*
使用指针地址相同，避免拷贝
	MakeBigPtr b addr: 0xc000300000
	main p addr: 0xc000300000
	p.Data[0] = 42
	----------
	MakeBigVal b addr: 0xc000500000
	main v addr: 0xc000400000
	v.Data[0] = 42
*/
