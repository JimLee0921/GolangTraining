package benchmark_test

import "testing"

/*
3 或 5 的倍数求和性能差异
*/

// 方法1: if 判断 (你写的思路)
func BenchmarkSumIf(b *testing.B) {
	counter := 0
	for i := 0; i < b.N; i++ {
		if i%3 == 0 || i%5 == 0 {
			counter += i
		}
	}
	_ = counter
}

// 方法2: 单独 if（可能更简洁）
func BenchmarkSumTwoIfs(b *testing.B) {
	counter := 0
	for i := 0; i < b.N; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	_ = counter
}

// 方法3: 只加 %15 来避免重复判断
func BenchmarkSumMod15(b *testing.B) {
	counter := 0
	for i := 0; i < b.N; i++ {
		if i%3 == 0 || i%5 == 0 { // 和 BenchmarkSumIf 一样
			counter += i
		}
	}
	_ = counter
}

/*
命令 -bench='Sum' 不会测试 hello_str_test 中的测试函数
go test -bench='Sum' .\exercise-solutions\01_basic\08_benchmark-test\
goos: windows
goarch: amd64
pkg: github.com/JimLee0921/GolangTraining/exercise-solutions/01_basic/08_benchmark-test
cpu: Intel(R) Core(TM) i3-10100 CPU @ 3.60GHz
BenchmarkSumIf-8        1000000000               0.2617 ns/op
BenchmarkSumTwoIfs-8    1000000000               0.2537 ns/op
BenchmarkSumMod15-8     1000000000               0.2589 ns/op
PASS
ok      github.com/JimLee0921/GolangTraining/exercise-solutions/01_basic/08_benchmark-test      0.871s
*/
