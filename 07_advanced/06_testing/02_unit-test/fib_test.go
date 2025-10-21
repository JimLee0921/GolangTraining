package main

import "testing"

// 表格驱动测试（Go 推荐的写法）
func TestFib(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{10, 55},
	}

	// 遍历执行测试
	for _, tt := range tests {
		got := Fib(tt.n)
		if got != tt.want {
			t.Errorf("Fib(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}
