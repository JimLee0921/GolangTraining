package main

import "fmt"

// main 演示 Go 常见的基本数据类型以及它们的使用方式
func main() {
	// 布尔型：只有 true 和 false 两个取值，经常用于条件判断
	fmt.Printf("布尔值 isReady: %t，类型: %T\n", true, false)

	/*
		整数：
			有符号：int8, int16, int32, int64, int（跟平台相关，通常 32/64 位）
			无符号：uint8（等价于 byte）、uint16, uint32, uint64, uint
			特殊别名：rune（等价于 int32，常用来表示 Unicode 码点）
		浮点数：float32, float64 默认使用 float64 以获得更高精度
		复数：complex64（实部虚部是 float32）、complex128（实部虚部是 float64）
	*/
	var distance int = 384400 // 地球到月球的平均距离，单位：千米
	var color uint8 = 255     // 0~255 之间的无符号整数
	fmt.Printf("整型 distance: %v，类型: %T\n", distance, distance)
	fmt.Printf("无符号整型 color: %v，类型: %T\n", color, color)

	var temperature float32 = 36.6
	pressure := 1013.25 // 通过 := 推断为 float64
	fmt.Printf("浮点数 temperature: %v，类型: %T\n", temperature, temperature)
	fmt.Printf("浮点数 pressure: %v，类型: %T\n", pressure, pressure)

	impedance := complex(5, -3)
	fmt.Printf("复数 impedance: %v，类型: %T\n", impedance, impedance)

	// rune 与 byte：rune 表示 Unicode 码点，byte 表示单个字节。
	var letter rune = '汉'
	var ascii byte = 'A'
	fmt.Printf("rune letter: %c，Unicode: %U，类型: %T\n", letter, letter, letter)
	fmt.Printf("byte ascii: %c，数值: %d，类型: %T\n", ascii, ascii, ascii)

	/*
		字符串：由 UTF-8 字节序列组成
		可以使用 len 函数计算其字节长度
		创建后不可变：一旦创建，内部的字节序列不能被改变。只能让变量指向一个新的字符串值
	*/
	greeting := "你好，Gopher"
	greeting = "yesyesyes"
	fmt.Printf("字符串 greeting: %q，字节长度: %d，类型: %T\n", greeting, len(greeting), greeting)

}
