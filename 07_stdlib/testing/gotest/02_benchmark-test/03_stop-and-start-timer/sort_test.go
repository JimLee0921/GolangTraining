package sort

import "testing"

/*
循环内多次需要准备数据的耗时操作使用 StopTimer 和 StartTimer 进行跳过计入

排序基准测试执行 b.N 次很慢
go test -bench .
只执行一次
go test -bench . -benchtime=1x
*/
func BenchmarkSort(b *testing.B) {
	base := makeData()

	b.ReportAllocs() // 开启更详细的报告
	b.ResetTimer()   // 一次性数据生成不计入耗时
	for b.Loop() {
		b.StopTimer()             // 暂停耗时
		data := prepareData(base) // 每次都需要准备新数据，也不计入耗时
		b.StartTimer()            // 开启耗时

		Sort(data)
	}
}
