package main

import (
	"encoding/json"
	"sync"
	"testing"
)

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "JimLee", Age: 25})

var studentPool = sync.Pool{
	New: func() any {
		return new(Student)
	},
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}

/*
cpu: Intel(R) Core(TM) i3-10100 CPU @ 3.60GHz
BenchmarkUnmarshal-8               12825             94826 ns/op            1392 B/op          7 allocs/op
BenchmarkUnmarshalWithPool-8       13096             91663 ns/op             240 B/op          6 allocs/op

Student 结构体内存占用较小，内存分配几乎不耗时间
而标准库 json 反序列化时利用了反射，效率是比较低的，占据了大部分时间，因此两种方式最终的执行时间几乎没什么变化
但是内存占用差了一个数量级，使用了 sync.Pool 后，内存占用仅为未使用的 240/1392，对 GC 的影响就很大了
*/
