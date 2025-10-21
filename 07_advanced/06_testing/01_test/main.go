package main

func Add(a, b int) int { return a + b }

/*
	Go 的测试规则是：
		测试文件必须以 *_test.go 结尾。并且测试文件中不能携带 func main()
		测试函数必须是：
			func TestXxx(t *testing.T)
			func BenchmarkXxx(b *testing.B)
			func ExampleXxx()
	运行实例：go test <module name>/<package name> 用来运行某个 package 内的所有测试用例
		运行当前 package 内的用例：go test example 或 go test .
		运行子 package 内的用例： go test example/<package name> 或 go test ./<package name>
		如果想递归测试当前目录下的所有的 package：go test ./... 或 go test example/...
	在 Go 里，官方内置的测试体系（全部基于 testing 包）主要就三种：
	1. 单元测试 (Unit Test)
		函数名以 Test 开头，校验功能正确性
		作用：验证函数逻辑是否正确
		运行单元测试：go test
	2. 基准测试 (Benchmark Test)
		函数名以 Benchmark 开头，测试性能效率
		作用：测试代码性能（运行速度、耗时、内存分配等）
		运行基准测试：`go test -bench .` 或 `go test -bench=.`
	3. 示例测试 (Example Test)
		函数名以 Example 开头
		作用：
			给用户提供代码示例，使用示例 + 自动校验输出
			自动校验输出是否和注释里的 // Output: 匹配
		运行示例测试：go test
	参数：
		-v：verbose 模式，显示每个测试函数的执行情况
		.Log：在 TestXxx 和 BenchmarkXxx 中可以使用 t/b.Log，然后使用 verbose 模式时就会打印日志
		-bench：go test 命令默认不运行 benchmark 用例的，如果想运行 benchmark 用例，则需要加上 -bench 参数
			-bench 参数支持传入一个正则表达式，匹配到的用例才会得到执行，例如，只运行以 Fib 结尾的 benchmark 用例：go test -bench='Fib$' .
		-run：可以指定一个正则表达式，只运行名字匹配正则的测试函数（TestXxx），参考 -bench 参数
		-benchmem：配合 -bench 使用，报告内存分配次数和字节数
		-count=N：运行同一个测试 N 次（默认 1）
		-timeout=DURATION：设置单个包的测试超时时间（默认 10m）
		-failfast：一旦有一个测试失败，立即停止整个测试
	遇到的问题：
		1. `go test -bench .` 正常应该是等价于 `go test -bench=.`
		但是 Windows PowerShell 里 . 有特殊含义（表示当前目录）所以第二个命令可能不会执行 bench
		可以这样写：-bench='.'
		2. Go 的 flag 解析器会把 -bench 当成一个参数，后面必须跟 正则表达式
		所以需要先指定路径，再加 -bench或者先加 -bench，再指定路径（注意 -bench 后面要有值）
		go test -v ./06_testing/03_benchmark-test -bench=.
		go test -v -bench=. ./06_testing/03_benchmark-test
	benchmark 中有个 b.N 属性，如果该用例能够在 1s 内完成，b.N 的值便会增加，再次执行。b.N 的值大概以 1, 2, 3, 5, 10, 20, 30, 50, 100 这样的序列递增，越到后面，增加得越快
*/
