package parse

import "testing"

/*
Fuzz 非常适合 []byte
不要求结果正确，只要求 健壮
所有外部输入都值得 fuzz
*/
func FuzzParseUint32(f *testing.F) {
	f.Add([]byte{0, 0, 0, 1})
	f.Add([]byte{})
	f.Add([]byte{1, 2, 3})

	f.Fuzz(func(t *testing.T, b []byte) {
		_, _ = ParseUint32(b)
	})
}
