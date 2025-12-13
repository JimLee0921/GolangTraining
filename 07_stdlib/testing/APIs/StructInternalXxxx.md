# Go testing 内核层

四个 testing 模块内部使用的 `go test` 的测试描述 AST

```
type InternalTest       // 普通单元测试
type InternalBenchmark  // 基准测试
type InternalFuzzTarget // 模糊测试
type InternalExample    // 文档测试
```

这四个 struct 是：

- testing 包内部用来描述测试项的源数据结构
- 是 `go test` 命令实现的一部分，在生成并运行测试二进制时使用的中间表示
- 普通用户基本用不到，主要了解
- 只在 `testing.MainStart()` 中使用
- `go test` 不是通过反射来找 `TestXxx`/`BenchmarkXxx`/`FuzzXxx`/`ExampleXxx`，而是通过这些结构驱动执行

> 开发时从来不会手动定义这些结构体，是在运行 `go test` 时自动生成并传递给 `testing.MainStart`
> 的，真正 `go test` 调用步骤参考 [StructM.md](StructM.md)

## `testing.InternalTest`

表示 TestXXX 也就是单元测试的运行描述：

```
type InternalTest struct {
	Name string
	F    func(*T)
}
```

- 存储测试名称
- 存储测试函数指针
- 支持 `-run` 正则匹配
- 支持子测试调度

表示每一个 `func TestXxx(t *testing.T)`，在 `go test` 内部，每一个单元测试会被转换为

```
InternalTest{
    Name: "TestXxx",
    F:    TestXxx,
}
```

## `testing/InternalBenchmark`

表示 BenchmarkXxx 也就是基准测试的运行描述：

```
type InternalBenchmark struct {
	Name string
	F    func(b *B)
}
```

- `-bench` 正则匹配
- 控制 benchmark 调度
- 管理 `b.N` 生命周期
- 输出统计数据

表示每一个 `func BenchmarkXxx(b *testing.B)`，在 `go test` 内部，每个基准测试会被转换为：

```
InternalBenchmark{
    Name: "BenchmarkXxx",
    F:   BenchmarkXxx,
}
```

## `testing.InternalFuzzTarget`

表示 FuzzXxx 也就是模糊测试的运行描述：

```
type InternalFuzzTarget struct {
	Name string
	Fn   func(f *F)
}
```

- `go test -fuzz=...` 进行目标匹配
- 初始化 Seed 种子
- 管理 `testdata/fuzz`
- 控制 fuzz engine 的生命周期

表示每个 `func FuzzXxx(f *testing.F)`，在 `go test` 内部每个模糊测试会被转换为：

```
InternalFuzzTarget{
    Name: "FuzzXxx",
    Fn  : FuzzXxx,
}
```

## `testing.InternalExample`

表示 ExampleXxx 也就是文档测试的运行描述：

```
type InternalExample struct {
	Name      string
	F         func()
	Output    string
	Unordered bool
}
```

- 用于文档测试（doctest）
- `go test` 会执行 Example
- 比较 stdout 与 `// Output:` 助手

表示每个 Example 测试函数，比如：

```
func ExampleAdd(){
    fmt.Println(Add(1, 2))
    // Output: 3
}
```

在 `go test` 内部会被转换为：

```
InternalExample{
    Name: "ExampleAdd",
    F   : ExampleAdd,
    Output: "3\n",
}
```

> Go 文档里的示例代码是可执行测试