package main

import "fmt"

type coder interface {
	code()
}

type Geeker struct {
	language string
}

func (g Geeker) code() {
	fmt.Printf("I'm coding %s language\n", g.language)
}

// main 方法集值类型接收值和指针
func main() {
	/*
		值接收者可以接收值类型和指针类型
	*/
	// 值类型
	var c1 coder = Geeker{"Go"}
	// 指针类型
	var c2 coder = Geeker{"Python"}
	c1.code()
	c2.code()

}
