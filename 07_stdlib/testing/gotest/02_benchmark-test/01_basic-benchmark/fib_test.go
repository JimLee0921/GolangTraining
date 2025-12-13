package fib

import "testing"

/*
最基础的基准测试
go test -bench . 运行当前包下所有基准测试
go test -bench BenchmarkFib_N 指定运行某一个基准测试
go test -bench ^BenchmarkFib_N$ 可以进行正则匹配
*/

// 老版本写法，使用 b.N 进行循环
func BenchmarkFib_N(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}

// 新版本写法，使用 b.Loop() 更推荐使用
func BenchmarkFib_Loop(b *testing.B) {
	for b.Loop() {
		Fib(10)
	}
}

//// 基准测试：递归 Fib
//func BenchmarkFib(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Fib(30) // 固定输入，模拟常见性能压力点
//	}
//}
//
//func BenchmarkHello(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		fmt.Sprintf("Hello")
//	}
//}
//
//// 如果有前置某些耗时操作可以使用 ResetTimer 重置计时器
//func BenchmarkHelloResetTimer(b *testing.B) {
//	// 耗时操作
//	time.Sleep(time.Second)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		fmt.Sprintf("Hello")
//	}
//}
//
//// 并行测试
//func BenchmarkHelloRunParallel(b *testing.B) {
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			fmt.Sprintf("Hello")
//		}
//	})
//}
