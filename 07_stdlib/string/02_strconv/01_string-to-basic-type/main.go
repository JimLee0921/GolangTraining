package main

import (
	"fmt"
	"strconv"
)

func ParseBoolDemo() {
	v1, _ := strconv.ParseBool("true")
	fmt.Printf("%T: %v\n", v1, v1) // bool: true

	v2, _ := strconv.ParseBool("F")
	fmt.Printf("%T: %v\n", v2, v2) // bool: false
}

func ParseIntDemo() {
	n1, _ := strconv.ParseInt("101", 2, 64)  // 二进制 5
	n2, _ := strconv.ParseInt("FF", 16, 64)  // 十六进制 255
	n3, _ := strconv.ParseInt("-42", 10, 32) // -42
	fmt.Printf("%T: %v\n", n1, n1)           // int64: 5
	fmt.Printf("%T: %v\n", n2, n2)           // int64: 255
	fmt.Printf("%T: %v\n", n3, n3)           // int64: -42
}

func ParseUnitDemo() {
	u, _ := strconv.ParseUint("FF", 16, 64) // 255
	fmt.Printf("%T: %v\n", u, u)            // uint64: 255
}

func ParseFloatDemo() {
	f1, _ := strconv.ParseFloat("3.14", 64) // float64(3.14)
	f2, _ := strconv.ParseFloat("1e6", 64)  // 科学计数法也可以
	fmt.Printf("%T: %v\n", f1, f1)          // float64: 3.14
	fmt.Printf("%T: %v\n", f2, f2)          // float64: 1e+06

}

func AtoiItoaDemo() {
	i, _ := strconv.Atoi("123")
	s := strconv.Itoa(123)       // "123"
	fmt.Printf("%T: %v\n", i, i) // int: 123
	fmt.Printf("%T: %v\n", s, s) // string: 123

}

func main() {
	ParseBoolDemo()
	ParseIntDemo()
	ParseUnitDemo()
	ParseFloatDemo()
	AtoiItoaDemo()
}
