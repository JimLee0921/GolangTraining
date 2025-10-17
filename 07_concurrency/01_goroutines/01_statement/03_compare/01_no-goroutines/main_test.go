package main

import "testing"

// BenchmarkFooBar 基准测试 foo + bar
func BenchmarkFooBar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Foo()
		Bar()
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/JimLee0921/GolangTraining/07_goroutines/01_basic/01_no-goroutines
cpu: Intel(R) Core(TM) i3-10100 CPU @ 3.60GHz
BenchmarkFooBar-8              1        9042287000 ns/op
PASS
ok      github.com/JimLee0921/GolangTraining/07_goroutines/01_basic/01_no-goroutines    9.051s
同步执行耗时九秒多
*/
