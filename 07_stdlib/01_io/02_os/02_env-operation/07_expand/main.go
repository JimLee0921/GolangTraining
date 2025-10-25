package main

import (
	"fmt"
	"os"
)

func main() {
	text := "Hello, ${USER}! Today is $DAY."

	// 自定义规则
	mapping := func(key string) string {
		switch key {
		case "USER":
			return "Jim"
		case "DAY":
			return "Friday"
		default:
			return "?"
		}
	}
	newText := os.Expand(text, mapping)
	fmt.Println(newText) // Hello, Jim! Today is Friday.
}
