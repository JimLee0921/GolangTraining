package main

import "fmt"

func add(a, b int) int { return a + b }

type BinOp func(int, int) int

func main() {
	var op BinOp = add // 合法
	fmt.Println(op(1, 2))

}
