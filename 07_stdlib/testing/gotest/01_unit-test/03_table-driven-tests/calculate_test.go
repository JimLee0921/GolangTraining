package math

import "testing"

/*
多个子测试推荐写法
对于多个子测试的场景，更推荐如下表驱动进行子测试的生成(table-driven tests)
*/
func TestMul(t *testing.T) {
	cases := []struct {
		Name       string
		A, B, want int
	}{
		{"pos", 2, 5, 10},
		{"neg", 2, -2, -4},
		{"zero", 20, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := Mul(c.A, c.B)
			if got != c.want {
				t.Errorf("Mul(%d, %d) = %d; want %d", c.A, c.B, got, c.want)
			}
		})
	}
}
