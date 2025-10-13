package main

import "fmt"

// main 另外一种函数表达式写法
func main() {
	/*
		函数还可以作为返回值
	*/
	greeter := makeGreeter()
	fmt.Println(greeter())
	fmt.Printf("%T\n", greeter()) // string

}

// makeGreeter 返回值类型为 func() string
func makeGreeter() func() string {
	return func() string {
		return "Hello World"
	}
}
