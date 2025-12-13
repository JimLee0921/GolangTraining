package process

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

/*
使用 SetBytes、ReportAllocs 和 TempDir进行 IO / 吞吐量 基准测试
*/
func BenchmarkWriteFile(b *testing.B) {
	const blockSize = 4 << 10 // 4KB
	data := make([]byte, blockSize)
	for i := range data {
		data[i] = byte(i)
	}

	dir := b.TempDir()    // 每个 benchmark 创建一个独立的临时目录
	b.SetBytes(blockSize) // 每次操作写入 4KB
	b.ReportAllocs()      // 看看每次写文件产生多少分配
	b.ResetTimer()        // setup 阶段不计入耗时

	for i := 0; i < b.N; i++ {
		path := filepath.Join(dir, strconv.Itoa(i))
		if err := os.WriteFile(path, data, 0o644); err != nil {
			b.Fatal(err)
		}
	}
}
