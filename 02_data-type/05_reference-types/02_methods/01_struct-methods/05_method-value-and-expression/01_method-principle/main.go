package main

import "fmt"

type Counter int

func (c *Counter) Add(n int) {
	*c += Counter(n)
}

// 上面的方法实际等价于下面这个函数，也就是说，Add 只是普通函数，但多了一个接收者参数

func Add(c *Counter, n int) {
	*c += Counter(n)
}

func main() {
	/*
		Go 的方法底层其实是 函数 + 隐式接收者参数
	*/
	// 使用上是一样的
	c := Counter(10)
	c.Add(20)
	fmt.Println(c) // 30
	Add(&c, 10)
	fmt.Println(c) // 40
}
