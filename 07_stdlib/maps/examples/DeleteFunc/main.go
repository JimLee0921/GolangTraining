package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}
	maps.DeleteFunc(m, func(s string, i int) bool {
		return i%2 != 0 // 删除值为奇数的键值对
	})
	fmt.Println(m)
}
