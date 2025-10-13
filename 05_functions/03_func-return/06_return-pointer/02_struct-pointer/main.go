package main

import "fmt"

type Counter struct {
	count int
}

func (c Counter) IncByValue() {
	c.count++
}

func (c *Counter) IncByPointer() {
	c.count++
}

func main() {
	c1 := Counter{}
	c1.IncByValue()
	fmt.Println(c1.count) // 0（值拷贝）

	c2 := Counter{}
	c2.IncByPointer()
	fmt.Println(c2.count) // 1（指针修改）
}
