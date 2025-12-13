package math

import "testing"

/*
子测试 go test -run TestMul/pos 运行指定子测试（子测试中可以继续嵌套子测试）
*/
func TestSub(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		got := Sub(10, 4)
		want := 6
		if got != want {
			t.Errorf("Sub(10, 4) = %d; want %d", got, want)
		}
	})
	t.Run("neg", func(t *testing.T) {
		got := Sub(10, 20)
		want := -10
		if got != want {
			t.Errorf("Sub(10, 20) = %d; want %d", got, want)
		}
	})
}
