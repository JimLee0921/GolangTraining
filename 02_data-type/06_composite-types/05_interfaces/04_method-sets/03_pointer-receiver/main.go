package main

import "fmt"

type coder interface {
	debug()
}

type Geeker struct {
	language string
}

func (g *Geeker) debug() {
	fmt.Printf("I'm debugging %s language\n", g.language)

}

func main() {
	/*
		指针接收者只能接收指针类型，不能接收值类型
	*/
	// 值类型
	var c1 coder = &Geeker{"Go"}
	c1.debug()
	// 指针类型

	//var c2 coder = Geeker{"Python"}	这里会报错，只要实现的方法中有一个指针接收者就必须使用指针
	//c2.debug()

}
