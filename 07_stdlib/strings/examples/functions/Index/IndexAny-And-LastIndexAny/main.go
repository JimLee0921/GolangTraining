package main

import (
	"fmt"
	"strings"
)

// IndexAnyDemo IndexAny 返回字符串中 chars 中任意的第一个字符索引
func IndexAnyDemo() {
	fmt.Println(strings.IndexAny("Chicken", "wen")) // 5
	fmt.Println(strings.IndexAny("Chicken", "cen")) // 3
	fmt.Println(strings.IndexAny("Chicken", "Cea")) // 0
	fmt.Println(strings.IndexAny("Chicken", "chi")) // 1
	fmt.Println(strings.IndexAny("Chicken", "ab"))  // -1
	fmt.Println(strings.IndexAny("Chicken", ""))    // -1
	fmt.Println(strings.IndexAny("", ""))           // -1
}

// LastIndexAnyDemo 从右往左查找
func LastIndexAnyDemo() {
	fmt.Println(strings.LastIndexAny("Chicken", "wen")) // 6
	fmt.Println(strings.LastIndexAny("Chicken", "cen")) // 6
	fmt.Println(strings.LastIndexAny("Chicken", "Cea")) // 5
	fmt.Println(strings.LastIndexAny("Chicken", "chi")) // 3
	fmt.Println(strings.LastIndexAny("Chicken", "ab"))  // -1
	fmt.Println(strings.LastIndexAny("Chicken", ""))    // -1
	fmt.Println(strings.LastIndexAny("", ""))           // -1
}
func main() {
	IndexAnyDemo()
	LastIndexAnyDemo()
}
