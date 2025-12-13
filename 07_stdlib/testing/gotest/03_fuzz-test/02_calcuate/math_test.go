package _2_calcuate

import "testing"

/*
Fuzz 非常擅长发现 panic
可以在 fuzz 里主动过滤非法输入
不一定要断言结果，只要保证程序健壮

go test -fuzz . -fuzztime=20s
使用fuzztime指定fuzz运行时间
*/
func FuzzDiv(f *testing.F) {
	f.Add(1, 2)
	f.Add(10000, 23141)

	f.Fuzz(func(t *testing.T, a, b int) {
		if b == 0 {
			return
		}
		_ = Div(a, b)
	})
}
