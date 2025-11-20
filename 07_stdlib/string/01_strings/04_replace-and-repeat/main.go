package main

import (
	"fmt"
	"strings"
)

func ReplaceDemo() {
	s := "Hello world world"
	fmt.Println(strings.Replace(s, "world", "haha", 1))  // Hello haha world
	fmt.Println(strings.Replace(s, "world", "haha", 0))  // Hello world world
	fmt.Println(strings.Replace(s, "world", "haha", -1)) // Hello haha haha
}

func ReplaceAllDemo() {
	s := "Hello world world"
	fmt.Println(strings.ReplaceAll(s, "world", "haha")) // Hello haha haha
}

func RepeatDemo() {
	line := strings.Repeat("Hello ", 5) // Hello Hello Hello Hello Hello
	fmt.Println(line)
}
func main() {
	ReplaceDemo()
	ReplaceAllDemo()
	RepeatDemo()
}
