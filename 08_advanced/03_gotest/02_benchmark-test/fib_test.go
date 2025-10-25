package main

import (
	"fmt"
	"testing"
	"time"
)

// 基准测试：递归 Fib
func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(30) // 固定输入，模拟常见性能压力点
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Hello")
	}
}

// 如果有前置某些耗时操作可以使用 ResetTimer 重置计时器
func BenchmarkHelloResetTimer(b *testing.B) {
	// 耗时操作
	time.Sleep(time.Second)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Hello")
	}
}

// 并行测试
func BenchmarkHelloRunParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fmt.Sprintf("Hello")
		}
	})
}
