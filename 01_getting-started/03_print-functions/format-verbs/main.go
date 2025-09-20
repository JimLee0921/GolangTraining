package main

import "fmt"

// 定义接口
type point struct {
	X int
	Y int
}

// main 演示 fmt.Printf 系列常用的格式化动词。
func main() {
	// 准备一组不同类型的示例数据 := 是 go 中的短变量声明运算符
	n := 42
	f := 3.1415926
	s := "This is \n string"
	b := true
	p := point{X: 3, Y: 4}
	fmt.Println()

	/*
		整数动词：十进制、二进制、八进制、十六进制小写、十六进制大写
		%d 十进制 (decimal)
		%b 二进制 (binary)
		%o 八进制 (octal)
		%x 十六进制 (hexadecimal, 小写)
		%X 十六进制 (hexadecimal, 大写)
	*/
	fmt.Println("Integers:")
	fmt.Printf("%d 十进制为: %d\n", n, n)
	fmt.Printf("%d 二进制为: %b\n", n, n)
	fmt.Printf("%d 八进制为: %o\n", n, n)
	fmt.Printf("%d 十六进制小写为: %x\n", n, n)
	fmt.Printf("%d 十六进制大写为: %X\n", n, n)
	fmt.Println()

	/*
		浮点数动词：默认小数、限制精度、科学计数法
		%f 默认保留六位小数且四舍五入
		%f.xf 保留 x 位小数
		%e 科学计数法（小写 e）
		%E 科学计数法（大写 E
		%g 最简表示保留有效位数（科学计数法时小写E）
		%G 最简表示保留有效位数（科学计数法时大写E）
	*/
	fmt.Println("Floats:")
	fmt.Printf("浮点数: %f\n", f)
	fmt.Printf("浮点数保留两位小数: %.2f\n", f)
	fmt.Printf("浮点数保留有效位数: %g\n", f)
	fmt.Printf("浮点数科学计数法: %e\n", f)
	fmt.Println()

	/*
		字符串动词：原样输出与带引号形式
		%s 原样输出字符串不加引号，不转义
		%q 会给字符串加上双引号里面的特殊字符（\n、\t 等）会自动转义
	*/
	fmt.Println("Strings:")
	fmt.Printf("字符串: %s\n", s)
	fmt.Printf("字符串: %q\n", s)
	fmt.Println()

	/*
		布尔值动词：true/false
	*/
	fmt.Println("Booleans:")
	fmt.Printf("布尔值: %t\n\n", n < 6)
	fmt.Printf("布尔值: %t\n\n", b)
	fmt.Println()
	/*
		类型与结构体动词：查看类型、值以及 Go 语法表示
		%T 打印变量的类型
		%v 打印变量的值结构体只显示字段的值（万能动词 int、string、float和结构体等都可以打印）
		%+v 打印值，结构体会带上字段名
		%#v 打印 Go 语法表示
	*/
	fmt.Println("Types and structs:")
	fmt.Printf("变量类型: %T\n", p)
	fmt.Printf("变量值不携带变量名: %v\n", p)
	fmt.Printf("变量值携带变量名: %+v\n", p)
	fmt.Printf("单个变量值: %v\n", p.X)
	fmt.Printf("Go语法表示: %#v\n", p)
}
