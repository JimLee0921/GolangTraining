package main

import (
	"fmt"
	"iter"
)

// RangeOneToN 定义 Seq 迭代器，返回一个 Seq 函数
func RangeOneToN(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			// 在调用方 range 循环里可以提前 break 这里就会返回 false 提前结束迭代器
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	// 使用 range 内部把循环体变成一个 yield 回调
	for v := range RangeOneToN(10) {
		fmt.Println(v)
	}

	// 使用 Pull
	next, stop := iter.Pull(RangeOneToN(100))
	defer stop()
	for {
		v, ok := next()
		if !ok {
			break // 等价于 range 结束
		}
		fmt.Println(v)
	}

}
