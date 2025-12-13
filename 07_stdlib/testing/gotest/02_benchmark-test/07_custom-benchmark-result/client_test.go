package client

import (
	"testing"
	"time"
)

/*
使用 ReportMetric + Context + Cleanup 自定义指标进行结果输出
*/
func BenchmarkClient_Do(b *testing.B) {
	client := &Client{}

	b.Cleanup(func() {
		// 连接关闭和资源清理等操作
		client.Close()
	})

	b.ReportAllocs() // 打印详细信息

	// 这里使用的 ctx 是 b.Context, 会随 benchmark 的生命周期/超时自动 cancel
	ctx := b.Context()
	var totalLatency int64

	b.ResetTimer() // 前面 setup 阶段不计入耗时

	for b.Loop() {
		start := time.Now()
		if err := client.Do(ctx); err != nil {
			b.Fatal(err)
		}
		totalLatency += time.Since(start).Nanoseconds()
	}
	avgLatency := float64(totalLatency) / float64(b.N)
	// 自定义 latency_ns 指标进行结果输出
	b.ReportMetric(avgLatency, "latency_ns")
}
