package benchmark_test

import (
	"fmt"
	"testing"
)

/*
字符串构造性能差异
*/

// 方法1: fmt.Sprintf
func BenchmarkHelloSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("Hello")
	}
}

// 方法2: 字符串常量
func BenchmarkHelloConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "Hello"
	}
}

// 方法3: 拼接（模拟多变量拼接场景）
func BenchmarkHelloConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "Hel" + "lo"
	}
}

/*
	测试生成 "hello" 字符串的三种方式的性能差异
	Go 的测试框架会自动控制 b.N 的值。 它一开始会尝试运行 1 次、2 次、N 次……不断调整，直到得到一个 稳定的耗时
	运行命令： go test -bench='Hello' .\exercise-solutions\01_basic\08_benchmark-test\
	如果在当前文件目录下：go test -bench='.'
	测试结果：
		BenchmarkHelloSprintf-8         26023029                49.86 ns/op
		BenchmarkHelloConst-8           1000000000               0.2673 ns/op
		BenchmarkHelloConcat-8          1000000000               0.2630 ns/op
	结果验证：
		Sprintf：每次要做格式化解析 -> 比较慢
		Const：直接取常量 -> 极快
		Concat：编译器会在编译期优化 -> 几乎和常量一样快
*/
//
