package main

import (
	"bytes"
	"sync"
	"testing"
)

var bufferPool = sync.Pool{New: func() any {
	return &bytes.Buffer{}
}}

var data = make([]byte, 10000)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}

/*
BenchmarkBufferWithPool-8       11560292               110.6 ns/op             0 B/op          0 allocs/op
BenchmarkBuffer-8                1139952              1052 ns/op           10240 B/op          1 allocs/op

这个例子创建了一个 bytes.Buffer 对象池，而且每次只执行一个简单的 Write 操作，存粹的内存搬运工，耗时几乎可以忽略
而内存分配和回收的耗时占比较多，因此对程序整体的性能影响更大
*/
