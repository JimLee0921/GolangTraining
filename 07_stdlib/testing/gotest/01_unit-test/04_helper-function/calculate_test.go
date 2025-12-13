package math

import "testing"

/*
帮助函数 helpers：把创建子测试的逻辑进行抽取并使用 t.Helper() 更精准的定位错误栈
没添加 t.Helper()
*/
type calculateCase struct {
	A, B, Want int
}

func createDivTestCase(t *testing.T, c calculateCase) {
	t.Helper() // 告诉测试框架这是个辅助函数，出现错误时错误位置应该指向调用者，而不是这个 helper 函数本身
	got := Div(c.A, c.B)
	if got != c.Want {
		t.Errorf("Div(%d, %d) = %d; want %d", c.A, c.B, got, c.Want)
	}
}

func TestDiv(t *testing.T) {
	createDivTestCase(t, calculateCase{10, 5, 2})
	createDivTestCase(t, calculateCase{4, 3, 1})
	createDivTestCase(t, calculateCase{5, 5, 2}) // 错误示例
}
