package main

import (
	"bytes"
	"fmt"
	"sync"
)

// 定义一个全局的 Buffer Pool
var bufPool = sync.Pool{New: func() any {
	// 当 Pool 中没有可用对象时会自动调用这里的 New 进行创建
	fmt.Println("create new buffer")
	return new(bytes.Buffer)
}}

// 模拟第一次请求处理
func handleRequest(id int, data string) {
	// 1. 从 Pool 中获取 buffer 对象（手动进行类型转换）
	buf := bufPool.Get().(*bytes.Buffer)

	// 2. 一定要先重置，因为可能获取到的是别人刚用完的，内容不可信
	buf.Reset()

	// 3. 确保使用完成后归还
	defer bufPool.Put(buf)

	// 使用获取到的 buf 对象
	buf.WriteString("request ")
	buf.WriteString(fmt.Sprint(id))
	buf.WriteByte(':')
	buf.WriteString(data)

	fmt.Println(buf.String())
}

func main() {
	var wg sync.WaitGroup

	// 模拟并发请求
	for i := 0; i < 100; i++ {
		// create new buffer 会 < 100 次，也就是实际创建的数量小于请求次数
		// 后续请求会复用已有的 buffer
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			handleRequest(id, "surprise, mother fxxker")
		}(i)
	}

	wg.Wait()
}
