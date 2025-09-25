package main

import "fmt"

/*
用 type 定义了一个新的类型，名字叫 foo
底层类型是 int，但是 foo 和 int 是两个不同的类型
这种方式叫 命名类型 (named type)，不是别名
*/
type foo int

func main() {
	/*
		“user defined types — we declare a new type, foo, the underlying type of foo: int”
		你可以自定义一个新类型 foo，其底层类型（underlying type）是 int。

		“conversion: int(myAge)”
		要把一个 foo 类型的值转换回 int，可以写：int(myAge)。

		“THIS CODE IS ONLY FOR EXAMPLE IT IS A BAD PRACTICE TO ALIAS TYPES one exception: if you need to attach methods to a type see the time package for an example of this godoc.org/time type Duration int64 Duration has methods attached to it”
		警告说：“把一个已有基本类型直接起别名（alias）”通常不是好习惯，但如果你要给这个类型 挂方法（attach methods），那就是一个合理的理由。举例：time.Duration 是 int64 的别名／自定义类型，但它被赋予了各种方法，比如 func (d Duration) String() string。
	*/
	var num2 foo // 声明一个变量 myAge,类型是 foo(不能直接和 int 混用，必须显式转换)
	num1 := 42
	num2 = 42
	num3 := int(num2)                  // 显式转换
	fmt.Printf("%T %v \n", num1, num1) // int 42
	fmt.Printf("%T %v \n", num2, num2) // main.foo 42
	fmt.Printf("%T %v \n", num3, num3) // int 42
	// fmt.Println(num1 + num2)           无效操作 类型不同
	fmt.Println(num1 + int(num1)) // 84

}
