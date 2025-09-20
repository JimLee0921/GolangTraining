package main

import "fmt"

/*
x 在函数外 声明，是一个包级变量，作用域是整个 main 包
这里不能使用 x := 42
:= 不能用于包作用域
只有在函数内部才能使用
*/
var x = 42

func main() {
	// y 只能在 main 函数内部中使用
	y := 666
	fmt.Printf("func main print x: %d\n", x)
	fmt.Printf("func main print y: %d\n", y)
	foo()
}

func foo() {
	// y 只能在 main 函数内部中使用 main 中也不能使用
	z := 777
	fmt.Printf("func foo print x: %d\n", x)
	fmt.Printf("func foo print z: %d\n", z)

}
