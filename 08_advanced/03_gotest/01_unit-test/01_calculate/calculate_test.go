package main

import "testing"

/*
 1. 普通测试文件
    运行 go test，该 package 下所有的测试用例都会被执行
    运行 go test -v，-v 参数会显示每个用例的测试结果，另外 -cover 参数可以查看覆盖率
    如果只想运行其中的一个用例，例如 TestAdd，可以用 -run 参数指定，该参数支持通配符 *，和部分正则表达式，例如 ^、$
*/
func TestAdd(t *testing.T) {
	if res := Add(1, 2); res != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", res)
	}
	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}
}

/*
 2. 子测试
    go test -run TestMul/pos 运行指定子测试
*/
func TestSub(t *testing.T) {
	// 使用 t.Run() 创建子测试 subtests
	t.Run("pos", func(t *testing.T) {
		if Sub(3, 1) != 2 {
			t.Fatal("fail")
		}
	})
	t.Run("neg", func(t *testing.T) {
		if Sub(2, 3) != -1 {
			t.Fatal("fail")
		}
	})
}

/*
3. 多个子测试推荐写法
对于多个子测试的场景，更推荐如下的写法(table-driven tests)

	所有用例的数据组织在切片 cases 中，看起来就像一张表，借助循环创建子测试。这样写的好处有：
		新增用例非常简单，只需给 cases 新增一条测试数据即可
		测试代码可读性好，直观地能够看到每个子测试的参数和期待的返回值
		用例失败时，报错信息的格式比较统一，测试报告易于阅读
		如果数据量较大，或是一些二进制数据，推荐使用相对路径从文件中读取
*/
func TestMul(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, -3, -6},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if res := Mul(c.A, c.B); res != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got", c.A, c.B, c.Expected, res)
			}
		})
	}
}

/*
4. 帮助函数 helpers：把创建子测试的逻辑进行抽取并使用 t.Helper() 更精准的定位错误栈
没添加 t.Helper()
*/
type calculateCase struct {
	A, B, Expected int
}

func createDivTestCas(t *testing.T, c *calculateCase) {
	t.Helper() // 告诉测试框架忽略本函数的栈帧
	if res := Div(c.A, c.B); res != c.Expected {
		t.Fatalf("%d / %d expected %d, but %d got", c.A, c.B, c.Expected, res)
	}
}

func TestDiv(t *testing.T) {
	createDivTestCas(t, &calculateCase{6, 3, 2})
	createDivTestCas(t, &calculateCase{10, 3, 3})
	createDivTestCas(t, &calculateCase{3, 1, 1}) // 错误示例，运行时会报错
}
