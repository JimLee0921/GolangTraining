package demo

import (
	"runtime"
	"testing"
)

// 如果分配的 CPU 数量少于 10 就直接跳过本基准测试
func BenchmarkParallelWork(b *testing.B) {
	if runtime.NumCPU() < 10 {
		b.Skip("skipping benchmark: need at least 10 CPUs")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = i * i
	}
}
