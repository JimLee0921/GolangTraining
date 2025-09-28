package main

import "fmt"

type coder interface {
	code()
	debug()
}

type Geeker struct {
	language string
}

func (g Geeker) code() {
	fmt.Printf("I'm coding %s language\n", g.language)
}

func (g *Geeker) debug() {
	fmt.Printf("I'm debugging %s language\n", g.language)

}

func main() {
	/*
		Geeker 类型如果传递给指针接收器时会自动转换为指针进行传递
	*/
	// 值类型
	var g = Geeker{"Go"}
	g.code()
	(&g).debug()
	g.debug() // 等同于 &g.debug()

}
