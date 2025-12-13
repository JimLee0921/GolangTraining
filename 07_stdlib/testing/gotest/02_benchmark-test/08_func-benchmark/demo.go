package main

import (
	"fmt"
	"testing"
)

/*
使用 testing.Benchmark 顶层函数来直接运行一个 benchmark 基准测试并获取 BenchmarkResult 测试结果
*/

func Add(a, b int) int {
	return a + b
}

func main() {
	result := testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs() // 展示更多结果信息
		for b.Loop() {
			_ = Add(1, 2)
		}
	})
	fmt.Println("raw", result)
	fmt.Println("ns/op:", result.NsPerOp())
	fmt.Println("allocs/op:", result.AllocsPerOp())
	fmt.Println("bytes/op:", result.AllocedBytesPerOp())
	fmt.Println("mem string:", result.MemString())
}
