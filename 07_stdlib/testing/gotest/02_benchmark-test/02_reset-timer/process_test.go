package demo

import "testing"

/*
基准测试配合 ResetTimer 不统计数据准备阶段的耗时
go test -bench .
*/
func BenchmarkProcess(b *testing.B) {
	// 一次性准备的数据，不计入总耗时
	data := makeHugeData()
	b.ReportAllocs() // 开启分配统计，可以配合 ResetTimer 使用
	b.ResetTimer()   // 上面准备数据的 setup 阶段不计入耗时
	for b.Loop() {
		_ = Process(data)
	}
}
