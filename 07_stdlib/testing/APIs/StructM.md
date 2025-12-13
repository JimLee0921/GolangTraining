# `testing.M`

`testing.M` 是整个 `go test` 进程级别的总调度器/生命周期管理器，
主要控制 `go test` 的 `初始化 -> 运行所有测试 -> 退出码`，唯一直接接触的地方就是使用顶层函数 `TestMain(m *testing.M)`

## 定义

就是传递给 TestMain 函数以运行实际测试的类型，都不用自己实现，主要了解

```
type M struct {
    deps       testDeps
    tests      []InternalTest
    benchmarks []InternalBenchmark
    fuzzTargets []InternalFuzzTarget
    examples   []InternalExample
}
```

deps 对应的是 `testing` 运行时依赖的一组钩子/能力，由 `cmd/go` 生成的 test main 提供实现，用于做一些并非纯测试逻辑的事情：

- 覆盖率相关支持（`-cover`、`-coverprofile`）
- 与 fuzz 引擎、seed corpus、崩溃用例持久化等配套能力
- 可能还包括某些运行环境/指标收集的桥接

> 几个 InternalXxx 就是对应的那些测试类型，参考 [StructInternalXxxx.md](StructInternalXxxx.md)

管理测试的完整生命周期：

- 测试开始前（进程级别）
- 所有测试执行
- 所有测试结束
- 在所有结束后通过 `m.Run()` 方法获取退出码，所有测试通过就是0，任意失败就是非0

> 不是单个 Test/Benchmark 的层级，而是全局级别

## 相关函数/方法

### 1. `testing.MainStart`

属于 `testing` 的内部入口函数，用于把 `go test` 发现到的测试项（Test/Benchmark/Fuzz/Example）交给 testing 运行框架，构造并返回一个
`testing.M`

```
func MainStart(deps testDeps, tests []InternalTest, benchmarks []InternalBenchmark, fuzzTargets []InternalFuzzTarget, examples []InternalExample) *M
```

无需关注调用，是 `go test` 编译阶段自动生成的 `_testmain.go` 中构造出的

### 2. Run 方法

`m.Run` 是测试进程级别的著执行函数，用于解析并应用 `go test` 的 flags，按照规则挑选并运行测试项，最后返回退出码

```
func (m *M) Run() (code int)
```

- 根据掺入的 flags 进行对应测试策略
- 汇总所有的测试结果并返回退出码：
    - 0：全部通过
    - 非0：有失败、panic、或某些条件导致的终止错误

## `go test` 执行概述

go 的 testing 并不靠反射运行测试，在执行 `go test` 命令后，执行逻辑如下：

1. `cmd/go` 扫描包
2. 找到所有符合命名规范的测试函数
3. 自动生成一个 `_testmain.go`
4. 在 `_testmain.go` 中进行各类测试函数的构造：
    ```
    var tests = []testing.InternalTest{...}
    var benchmarks = []testing.InternalBenchmark{...}
    var fuzzTargets = []testing.InternalFuzzTarget{...}
    var examples = []testing.InternalExample{...}
    ```
5. `go test` 自动调用执行：`testing.MainStart(...).Run()`

> 如果用户自定义了 TestMain 函数参考下面的逻辑行为

## `TestMain`

`TestMain` 是用户可以插入自定义的进程级别测试入口钩子，允许在 `m.Run` 前后自定义一些比如全局初始化请理等操作，
如果测试文件中存在 `TestMain` 函数，`go test` 就会优先调用用户自定义的 `TestMain`，而不是使用默认的。

### 签名

硬性要求！！！

```
func TestMain(m *testing.M)
```

- 函数名必须是 `TestMain`
- 参数必须是 `*testing.M `
- 不能有返回值
- 必须显式调用 `os.Exit`

### 执行逻辑

用户定义了 `TestMain` 函数的话，`MainStart/Run` 完整执行链路如下：

1. 用户执行 `go test`
2. `cmd/go` 扫描包并生成 `_testmain.go`
3. 检查是否有 `TestMain`：
    - 没有：执行 `main(){os.Exit(m.Run())}`
    - 有：执行 `main{TestMain(m)}`

- 如果不定义 `TestMain` 使用的默认行为就是直接 `os.Exit(m.Run())`，不会进行全局 `setup/teardown`，每个 Test
  需要自己做初始化和清理工作，适合做简单测试
- 使用 `TestMain` 可以接管声明周期，可以理解为钩子函数
    ```
    func TestMain(m *testing.M) {
    setup()        // 全局初始化
    code := m.Run()
    teardown()     // 全局清理
    os.Exit(code)
    }
    ```

### 适用场景

- 初始化一次数据库 / Redis
- 启动 mock server
- 设置全局环境变量
- 初始化日志系统
- 准备大型测试数据
- 测试结束后统一清理资源

> TestMain 是进程级别的一次性逻辑，不是测试用例级别的！

### 注意事项

- flags 解析行为发生在 TestMain之前
- `TestMain` 不能阻止 flags 生效，flags 由 `m.Run()` 内部执行，`TestMain` 只是做一层包装
- `TestMain` 中不能使用 `*testing.T`/`*testing.B` 等测试用例相关操作
- 不能跳过 `m.Run()`，否则测试不会被执行
- 不能定义多个 `TestMain`