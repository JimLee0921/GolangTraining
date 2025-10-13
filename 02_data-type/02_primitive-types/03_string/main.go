package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/*
		字符串：由 UTF-8 字节序列组成
		可以使用 len 函数计算其字节长度
		创建后不可变：一旦创建，内部的字节序列不能被改变。只能让变量指向一个新的字符串值
	*/

	// rune 类型
	ch1 := 'A' // rune
	ch2 := '中' // rune

	fmt.Printf("%c %d\n", ch1, ch1) // A 65
	fmt.Printf("%c %d\n", ch2, ch2) // 中 20013

	// string 包常用方法
	strings.Contains("hello", "he")       // 是否包含指定字符串
	strings.HasPrefix("hello", "he")      // 是否已指定字符串开头
	strings.HasSuffix("hello", "lo")      // 是否已指定字符串结尾
	strings.ToUpper("hello")              // 全部字符转小写
	strings.ToLower("HELLO")              // 全部字符转大写
	strings.Split("a,b,c", ",")           // 字符串转为array
	strings.Join([]string{"a", "b"}, "-") // array转为字符串
	// 字符串替换 s：原字符串 old：要被替换的子串 new：替换成的新子串 n：替换次数 （-1 为全部等同于strings.ReplaceAll(s, old, new)）
	strings.Replace("go go", "go", "java", -1)

	// strconv 包常用方法
	// 字符串转数值
	// Atoi：string → int
	i, _ := strconv.Atoi("123")
	fmt.Println(i)
	// ParseInt：string → int64（可指定进制和位数）
	n, _ := strconv.ParseInt("123", 10, 64) // 十进制，64位
	fmt.Println(n)
	// ParseUint：string → uint64; ParseFloat：string → float64
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println(f)
	// 数值转字符串
	// Itoa：int → string
	s1 := strconv.Itoa(123) // "123"
	fmt.Println(s1)
	// FormatInt：int64 → string（可指定进制）
	s2 := strconv.FormatInt(123, 2) // "1111011" (二进制)
	fmt.Println(s2)
	// FormatUint：uint64 → string; FormatFloat：float64 → string（可指定格式和精度）
	s3 := strconv.FormatFloat(3.14159, 'f', 2, 64) // "3.14"
	fmt.Println(s3)
	// 布尔值转换
	// ParseBool：string → bool
	b, _ := strconv.ParseBool("true") // true
	fmt.Println(b)
	// FormatBool：bool → string
	s := strconv.FormatBool(false) // "false"
	fmt.Println(s)
}
