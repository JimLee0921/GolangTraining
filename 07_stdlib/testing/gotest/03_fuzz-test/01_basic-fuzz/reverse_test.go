package reverse

import "testing"

/*
测试 Reverse 函数对一个字符串反转两次是否还相同
使用 go test -fuzz . 或 go test -fuzz=FuzzReverse
错误后会在同级目录下生成一个 testdata/ 里面有测试失败信息
*/
func FuzzReverse(f *testing.F) {
	// 添加种子
	f.Add("Hello")
	f.Add("")
	f.Add("大傻逼")
	f.Add("🙂")

	f.Fuzz(func(t *testing.T, s string) {
		r := Reverse(s)
		rr := Reverse(r)
		if rr != s {
			t.Fatalf("double reverse failed: %q -> %q -> %q", s, r, rr)
		}
	})
}
