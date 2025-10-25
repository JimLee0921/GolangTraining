package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

/*
使用 t.Cleanup() 进行定义
*/
func TestFile(t *testing.T) {
	// setup：创建临时目录
	dir := t.TempDir()
	file := filepath.Join(dir, "data.txt")
	os.WriteFile(file, []byte("hello"), 0644)

	// teardown: 自动清理（即使测试失败也会执行）
	t.Cleanup(func() {
		fmt.Println("teardown cache: ", dir)
		os.RemoveAll(dir)
	})

	// 逻辑测试
	data, _ := os.ReadFile(file)
	if string(data) != "hello" {
		t.Fatalf("expected hello, got %s", data)
	}
}

/*
2. 自定义封装 setup 和 teardown 函数
*/
func setup(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	t.Log("Setup done:", dir)
	return dir
}

func teardown(t *testing.T, dir string) {
	t.Helper()
	os.RemoveAll(dir)
	t.Log("Teardown done:", dir)
}

func TestExample(t *testing.T) {
	dir := setup(t)
	// 使用 defer 确保请理逻辑执行
	defer teardown(t, dir)

	// 测试逻辑
}

/*
3. 整个包的所有测试前后做准备或清理工作

如果测试文件中包含函数 TestMain，那么生成的测试将调用 TestMain(m)，而不是直接运行测试
调用 m.Run() 触发所有测试用例的执行，并使用 os.Exit() 处理返回的状态码，如果不为0，说明有用例失败
因此可以在调用 m.Run() 前后做一些额外的准备(setup)和回收(teardown)工作
这里的 setup 和 teardown 也可以封装为函数
*/
func TestMain(m *testing.M) {
	fmt.Println("全局 Setup：连接数据库")
	code := m.Run() // 运行所有测试
	fmt.Println("全局 Teardown：关闭数据库")
	os.Exit(code)
}
