package demo

import (
	"testing"
	"time"
)

// TestSlowOperation 如果 go test 开启 -short 模式，直接跳过测试，并会输出分支内容，主要需要加上 -v 参数才能看到打印内容
func TestSlowOperation(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped slow test in short mode")
	}

	time.Sleep(2 * time.Second)
}
