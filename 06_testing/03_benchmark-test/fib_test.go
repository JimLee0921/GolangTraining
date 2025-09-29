package main

import "testing"

// 基准测试：递归 Fib
func BenchmarkFibRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(30) // 固定输入，模拟常见性能压力点
	}
}

/*
	go test -v -bench . ./06_testing/03_benchmark-test/
	测试结果：
		goos: windows
		goarch: amd64
		pkg: github.com/JimLee0921/GolangTraining/06_testing/03_benchmark-test
		cpu: Intel(R) Core(TM) i3-10100 CPU @ 3.60GHz
		BenchmarkFibRecursive
		BenchmarkFibRecursive-8              292           4073576 ns/op
		PASS
		ok      github.com/JimLee0921/GolangTraining/06_testing/03_benchmark-test       1.611s

	go test -v -bench . ./06_testing/03_benchmark-test/ -benchmem
		goos: windows
		goarch: amd64
		pkg: github.com/JimLee0921/GolangTraining/06_testing/03_benchmark-test
		cpu: Intel(R) Core(TM) i3-10100 CPU @ 3.60GHz
		BenchmarkFibRecursive
		BenchmarkFibRecursive-8              267           4177590 ns/op               0 B/op          0 allocs/op
		PASS
		ok      github.com/JimLee0921/GolangTraining/06_testing/03_benchmark-test       1.612s
*/
