package main

import "fmt"

func main() {
	/*
		使用 any (interface{}) 存储任意数据类型
	*/
	data := map[string]any{
		"id":      101,
		"name":    "Tom",
		"balance": 99.8,
		"active":  true,
		"tags":    []string{"vip", "premium"},
	}

	for k, v := range data {
		fmt.Printf("%s => (%T) %v\n", k, v, v)
	}
	/*
		active => (bool) true
		tags => ([]string) [vip premium]
		id => (int) 101
		name => (string) Tom
		balance => (float64) 99.8
	*/
}
