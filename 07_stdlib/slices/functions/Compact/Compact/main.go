package main

import (
	"fmt"
	"slices"
)

func main() {
	seq := []int{0, 1, 1, 2, 3, 5, 5, 5, 8}
	newSeq := slices.Compact(seq)
	// 会影响原数组，真正使用直接原 seq 进行接收
	fmt.Println(seq)
	fmt.Println(newSeq)
}
