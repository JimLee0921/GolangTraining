package main

import (
	"fmt"
	"slices"
)

func main() {
	seq := []int{0, 1, 1, 2, 3, 5, 8}
	newSeq := slices.DeleteFunc(seq, func(i int) bool {
		return i%2 != 0 // 删除 偶数
	})
	fmt.Println(seq) // 会影响原切片，真正使用直接还使用 seq 进行接收，这里只是示例
	fmt.Println(newSeq)
}
