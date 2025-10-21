package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func process(id int, text string, wg *sync.WaitGroup) {
	defer wg.Done()
	// 1. 从池中拿出一个 *bytes.BUffer
	buf := bufPool.Get().(*bytes.Buffer)

	// 2. 清空（必须，防止上一次残留）
	buf.Reset()

	// 3. 使用
	buf.WriteString(fmt.Sprintf("[worker-%d] %s", id, text))

	// 4. 读出内容（仅展示）
	fmt.Println(buf.String())

	// 5. 使用完后放回池子供其它 goroutine 使用
	bufPool.Put(buf)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go process(i, "Hello Pool", &wg)
	}
	wg.Wait()
}
