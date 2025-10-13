package main

import "fmt"

const (
	_ = iota      // 把 0 丢弃
	b = iota * 10 // 1 * 10
	c             // 2 * 10
	d             // 3 * 10
)

// main iota 跳过某些值
func main() {
	fmt.Println(b, c, d) // 10 20 30
}
