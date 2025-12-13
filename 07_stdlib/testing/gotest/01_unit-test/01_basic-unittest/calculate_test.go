package math

import "testing"

/*
普通单元测试

	运行 go test，该 package 下所有的测试用例都会被执行
	运行 go test -v，-v 参数会显示每个用例的测试结果，另外 -cover 参数可以查看覆盖率
	如果只想运行其中的一个测试函数用例，例如 TestAdd，可以用 -run 参数指定，该参数支持通配符 *，和部分正则表达式，例如 ^、$
*/
func TestAdd(t *testing.T) {
	got1 := Add(3, 5)
	want1 := 8
	if got1 != want1 {
		t.Errorf("Add(3, 5) = %d; want %d", got1, want1)
	}

	got2 := Add(10, 20)
	want2 := 30
	if got2 != want2 {
		t.Errorf("Add(10, 20) = %d, want %d", got2, want2)
	}
}
