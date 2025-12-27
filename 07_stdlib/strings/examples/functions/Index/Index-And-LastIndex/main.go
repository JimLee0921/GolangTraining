package main

import (
	"fmt"
	"strings"
)

// IndexDemo Index 返回字符串中指定子串第一次出现的位置
func IndexDemo() {
	fmt.Println(strings.Index("Chicken", "ken")) // 4
	fmt.Println(strings.Index("Chicken", "Cen")) // -1
	fmt.Println(strings.Index("Chicken", ""))    // 0
	fmt.Println(strings.Index("Chicken", "chi")) // -1
	fmt.Println(strings.Index("", ""))           // 0
}

// LastIndexDemo LastIndex 规则和 Index 相反，从右往左查找
func LastIndexDemo() {
	fmt.Println(strings.LastIndex("Chicken", "ken")) // 4
	fmt.Println(strings.LastIndex("Chicken", "Cen")) // -1
	fmt.Println(strings.LastIndex("Chicken", ""))    // 7
	fmt.Println(strings.LastIndex("Chicken", "chi")) // -1
	fmt.Println(strings.LastIndex("", ""))           // 0
}

func main() {
	IndexDemo()
	LastIndexDemo()
}
