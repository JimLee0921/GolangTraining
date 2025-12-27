package main

import (
	"fmt"
	"strings"
)

// Count 统计子串出现次数（不重叠）
func main() {
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", ""))    // 5
	fmt.Println(strings.Count("aaaa", "aa"))  // 2
	fmt.Println(strings.Count("aaaa", "aaa")) // 1
}
