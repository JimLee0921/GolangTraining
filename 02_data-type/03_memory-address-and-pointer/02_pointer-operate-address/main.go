package main

import "fmt"

// main 操作内存地址
func main() {
	/*
		变量可以用 & 取地址，用指针 * （解引用运算符）操作地址
		p 是一个指针，指向一个 int 类型的值
		用 *p 解引用，就能拿到 a 的值
		Go 没有真正的引用传递，即使传的是地址，本质上还是把“地址值”复制了一份传进去
		但通过这个地址，可以间接修改原数据，同时避免拷贝大块数据，更高效
	*/
	x := 42
	p := &x // p 是 *int 类型，指向 x

	fmt.Println("x value:", x)
	fmt.Println("point p:", p)
	fmt.Println("get x value by point:", *p)

	*p = 100 // 修改指针指向的值
	/*
		this is useful
		we can pass a memory address instead of a bunch of values (we can pass a reference)
		and then we can still change the value of whatever is stored at that memory address
		this makes our programs more performant
		we don't have to pass around large amounts of data
		we only have to pass around addresses
		everything is PASS BY VALUE in go, btw
		when we pass a memory address, we are passing a value
	*/
	fmt.Println("changed x value:", x)
	/*
		x value: 42
		point p: 0xc00000a088
		get x value by point: 42
		changed x value: 100
	*/
}
