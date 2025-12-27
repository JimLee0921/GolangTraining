package main

import (
	"fmt"
	"strings"
)

// strings.Builder
func main() {
	var b strings.Builder

	// 1. 初始化状态，Len=0，cap 可能为 0
	fmt.Printf("[init]	Len=%d Cap=%d String=%q\n", b.Len(), b.Cap(), b.String())

	// 2. 预分配，确保还能再写入至少64个字节，减少自动扩容次数
	b.Grow(64)
	fmt.Printf("[grow]	Len=%d Cap=%d\n", b.Len(), b.Cap())

	// 3. 写入：WriteString/WriteByte/WriteRune/Write
	b.WriteString("i am ") // 最常用，直接追加 string
	b.WriteString("JimLee")

	b.WriteByte('.') // 单字节写入，常用于分隔符
	b.WriteByte(' ') // 单字节写入，常用于分隔符

	nRune, _ := b.WriteRune('中') // 按 UTF-8 编码写入（'中' 占3个字节）
	fmt.Printf("[rune] wrote %d bytes for rune '中'", nRune)

	b.WriteByte(' ')
	b.Write([]byte("i am 22 years old")) // 写入字节数组

	// 4. 此时长度容量观察
	fmt.Printf("[mid]	Len=%d Cap=%d\n", b.Len(), b.Cap())

	// 5. 产出结果：String 方法应是一个 builder 的使用终点
	s := b.String()
	fmt.Printf("[finial] %s\n", s)

	// 6. 想要复用需要先使用 Reset
	b.Reset()
	fmt.Printf("[reset]  Len=%d Cap=%d String=%q\n", b.Len(), b.Cap(), b.String())

	b.WriteString("大傻逼")
	s2 := b.String()
	fmt.Printf("[finial] %s\n", s2)

}
