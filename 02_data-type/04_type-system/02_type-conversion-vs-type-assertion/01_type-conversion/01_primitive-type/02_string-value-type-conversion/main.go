package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		字符串和数值之间不能直接用 T(v)，要用标准库：
		strconv库中常见转换方法
			1. strconv.Atoi(s)：ASCII to int
				把一个 十进制字符串 转换成 int。
				参数：s（必须是合法的十进制字符串）
				返回：int + error
				只能解析十进制返回值类型是 int，跟机器架构相关（32 位机是 int32，64 位机是 int64）
			2. strconv.Itoa(i)：Int to ASCII
				把一个 int 转换成 字符串
				只支持十进制输出
				只接受 int 类型，不能直接传 int64，如果需要传入 int64，应该用 strconv.FormatInt
			3. strconv.ParseInt(s string, base int, bitSize int) (i int64, err error)
				比 Atoi 更通用
				把字符串 s 解析成 int64。
				参数：
					s：待解析的字符串
					base：进制（2 到 36；如果为 0，会根据前缀自动判断："0x" -> 16，"0" -> 8，其他 -> 10）
					bitSize：整数大小（0、8、16、32、64），决定溢出检查范围。返回的结果会再转换到这个范围内
				返回：int64 + error
			4. strconv.ParseUint(s string, base int, bitSize int) (uint64, error)
				无符号整数版本的 ParseInt 把字符串 s 解析成 uint64
				参数：
					s：待解析的字符串。
					base：进制（2 到 36；如果为 0，会根据前缀自动判断："0x" -> 16，"0" -> 8，其他 -> 10）
					bitSize：整数大小（0、8、16、32、64），决定溢出检查范围。返回的结果会再转换到这个范围内
				返回：
					uint64
					error（如果字符串不是合法无符号数，或超出范围，返回错误）
			5. strconv.ParseFloat(s string, bitSize int) (float64, error)
				把字符串 s 解析成浮点数。
				参数：
					s：待解析的字符串（支持十进制、科学计数法 "1.23e4"）
					bitSize：目标浮点数大小（32 或 64）。
						32 -> 结果会转换为 float32 再转为 float64 返回
						64 -> 直接作为 float64。
				返回：
					float64
					error（解析失败或溢出时返回错误）
			6.strconv.ParseBool(s string) (bool, error)
				把字符串 s 解析成布尔值 不区分大小写
				参数：
					s：待解析的字符串
						true 值："1", "t", "T", "true", "TRUE", "True"
						false 值："0", "f", "F", "false", "FALSE", "False"
				返回：
					bool
					error（如果不是合法布尔字符串，返回错误）
	*/

	// Atoi
	x := "12"
	y := 6
	z, _ := strconv.Atoi(x)
	fmt.Println(y + z)

	// Itoa
	m := 12
	//b := "I have so many apples : " + a	// 这样会报错 int 与 string 不匹配
	n := "I have so many apples : " + strconv.Itoa(m)
	fmt.Println(n)

	// ParseInt
	//	ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)

	fmt.Println(b, f, i, u)

	//	FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:
	w := strconv.FormatBool(true)
	t := strconv.FormatFloat(3.1415, 'E', -1, 64)
	v := strconv.FormatInt(-42, 16)
	e := strconv.FormatUint(42, 16)

	fmt.Println(w, t, v, e)
}
