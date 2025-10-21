package main

import (
	"fmt"
	"testing"
)

// TestAdd 单元测试 —— 验证逻辑正确性
func TestAdd(t *testing.T) {
	t.Log("start TestAdd...")
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2,3) = %d; want %d", got, want)
	}
}

// BenchmarkAdd 基准测试 —— 测性能
func BenchmarkAdd(b *testing.B) {
	b.Log("start BenchmarkAdd...")
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

// ExampleAdd 示例测试 —— 给用户示例，同时校验输出
func ExampleAdd() {
	fmt.Println(Add(2, 3))
	// Output:
	// 5
}
