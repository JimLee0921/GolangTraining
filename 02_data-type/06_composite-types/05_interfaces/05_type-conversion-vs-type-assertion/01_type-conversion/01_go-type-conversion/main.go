package main

import "fmt"

// main Go中的类型转换
func main() {
	/*
		Go 里，类型转换（type conversion）是指把一个值从某种类型显式地转成另一种类型
		这和某些编程语言中的强制类型转换类似，但 Go 的设计更安全：
			Go 中如果不进行类型转换直接运算比如 int + float 会编译失败
			Go 转换必须显式写出目标类型（Go中没有隐式转换，不会自动转换）
			只能在“兼容”或“可转换”的类型之间进行
			Go 不支持“野蛮”的指针强转（除非用 unsafe 包）
		T(v)
			T 是目标类型
			v 是要转换的值
		可转换的情况：
			1. 基本数值类型：不同数值类型之间可以转换（int、float、uint...），注意可能出现精度丢失或截断
			2. 字符与数字：Go 的 rune（int32）和 byte（uint8）本质上是整数，可以和 int 互转
			3. 字符串与 byte/rune 切片：[]byte 是 UTF-8 编码的字节序列，[]rune 是 Unicode 码点的序列
	*/
	// 数值
	var a int = 10
	var b float64 = float64(a)
	fmt.Printf("%v: %T\n", b, b) // 10: float64

	// rune
	var c rune = '你'
	fmt.Printf("%v: %T\n", c, c) // 20320: int32

	// string 和 []byte
	s := "go"
	bs := []byte(s)
	fmt.Printf("%v: %T\n", bs, bs) // [103 111]: []uint8

	// string 和 []rune
	rs := []rune("中国")
	fmt.Printf("%v: %T\n", rs, rs) // [20013 22269]: []int32
}
