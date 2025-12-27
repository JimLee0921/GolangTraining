package main

import (
	"fmt"
	"sync/atomic"
)

const (
	FlagRead  uint32 = 1 << 0 // 0001
	FlagWrite uint32 = 1 << 1 // 0010
	FlagExec  uint32 = 1 << 2 // 0100
)

func main() {
	var flags atomic.Uint32

	// 打开 Read 和 Write
	flags.Or(FlagRead | FlagWrite)

	fmt.Printf("flags = %03b\n", flags.Load())

	// 关闭 Write
	flags.And(^FlagWrite)

	fmt.Printf("flags = %03b\n", flags.Load())

	// 判断是否有 Exec 权限
	hasExec := flags.Load()&FlagExec != 0
	fmt.Println("has exec:", hasExec)
}
