# `testing.TB`

`testing.TB` 是 testing 包中的一个核心接口（interface），testing 框架的最低通用抽象，所有测试行为的统一基类。
`testing.T` 单元测试、`testing.B` 基准测试、`testing.F` 模糊测试都实现了 `testing.TB` 接口。

## 接口方法定义

`testing.T`/`testing.B`/`testing.F`，实现接口都是指针接收者，这里讲解完 TB 接口方法后在学习 struct T/B/F 中就一笔带过了

```
type TB interface {
	Attr(key, value string)
	Cleanup(func())
	Error(args ...any)
	Errorf(format string, args ...any)
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Helper()
	Log(args ...any)
	Logf(format string, args ...any)
	Name() string
	Setenv(key, value string)
	Chdir(dir string)
	Skip(args ...any)
	SkipNow()
	Skipf(format string, args ...any)
	Skipped() bool
	TempDir() string
	Context() context.Context
	Output() io.Writer
	// contains filtered or unexported methods
}
```

## 元信息相关方法

### Attr

Attr 用于给当前测试（包括子测试）打标签（attribute），是无序的

```
Attr(key, value string)
```

- 可以理解为测试的 metadata: `module=payment`、`severity=critical`、`type=integration` 等
- 这些属性会记录在 testing 内部结构中，未来可以被测试过滤器，报告工具，日志系统调用

### Name

返回当前测试（或 benchmark、fuzz 子测试）的完整名称路径

```
Name() string
```

- 顶层测试：TestAdd
- 子测试：TestAdd/Case1
- 嵌套子测试：TestAdd/Case1/SubCaseA
- 基准测试：BenchmarkAdd-8 之类

> 用于调试日志前缀（特别是并行子测试中）和临时文件 / 临时目录命名前缀（确保可读性）

## 生命周期与请理相关方法

### Cleanup

为当前测试（或子测试）注册一个结束时自动执行的回调函数，无论测试正常结束或遇到 Fatal/Skip 都会执行回调函数。

```
Cleanup(f func())
```

执行顺序：

- 同一测试内的多个 Cleanup 按 LIFO 顺序 执行（后注册先执行）
- 子测试的 Cleanup 会在父测试的 Cleanup 之前执行

常见用途：

- 清理临时目录或文件
- 关闭网络连接、HTTP 服务器、数据库等资源
- 恢复修改过的状态（环境变量、配置等）

> 适用于 T/B/F

## 日志相关方法

### Log/Logf

Log 和 Logf 主要用于打印日志，与 `fmt.Println/fmt.Printf` 类似，其输出行为与测试模式相关：

- 默认情况下，日志仅在测试失败时展示
- 使用普通单元测试T `go test -v`（verbose 模式）时始终打印
- 对于基准测试B，始终会打印文本，以避免性能依赖于 `-v` 标志的值

```
Log(args ...any)
Logf(format string, args ...any)
```

### Output

Output 返回一个 `io.Writer`，所有写入这个 Writer 的内容，会以合适的缩进和格式记录到当前测试的输出中

```
Output() io.Writer
```

在 testing 实现里，这个 Writer 是一个包装器（indenter），会：

- 把每一行前面加上缩进
- 把数据追加到当前测试缓冲的 output 中
- 最后由父测试 / 顶层统一 flush 输出

**适用场景**

1. 把外部日志库的输出重定向到测试日志中
2. 配合 `log/slog` 或自建 logger，把结构化日志写入测试输出

## 失败控制相关方法

### Error/Errorf

Error 和 Errorf 两个方法会标志失败信息，但是测试会继续执行，底层调用 Fail 方法，适用于测试能够继续执行的非致命错误

- Error 等价于 Log 方法 + Fail 方法
- Errorf 等价于 Logf 方法 + Fail 方法

```
Error(args ...any)
Errorf(format string, args ...any)
```

### Fatal/Fatalf

打印失败日志并立即终止测试（调用 FailNow 方法）

- Fatal 等价于 Log 方法 + FailNow 方法
- Fatalf 等价于 Logf 方法 + FailNow 方法

```
Fatal(args ...any)
Fatalf(format string, args ...any)
```

> Fatal 系列常用于初始化失败或无法继续的场景，例如参数不合法、环境缺失等

### Fail/FailNow/Failed

Fail 和 FailNow 为较低级 API，主要为 Error 和 Fatal 系列提供辅助。

- Fail 标记错误但继续执行
- FailNow 标记错误并立即终止整个测试
    - 内部通过 `runtime.Goexit` 终止当前 goroutine
    - 在 `T/B/F` 中使用均安全且行为一致
- Failed 返回当前测试是否已标记为失败

```
Fail()
FailNow()
Failed() bool
```

> 这些方法为 Error 和 Fatal 提供底层语义支撑

## 测试跳过相关方法

### Skip/Skipf

测试跳过相关，Skip 和 Skipf 为立即终止当前测试并标记测试状态为 skipped 而不是失败

- Skip 等价于 Log 方法 + SkipNow 方法
- Skipf 等价于 Logf 方法 + SkipNow 方法

```
Skip(args ...any)
Skipf(format string, args ...any)
```

**使用场景**

- 某些平台无法运行的测试（如 Windows 不支持）
- 依赖外部环境（如 Docker、数据库）未满足时跳过
- 长时间测试在 go test -short 下跳过

### SkipNow/Skipped

- Skipped 方法返回测试是否被跳过
- SkipNow 会将测试标记为跳过并立即退出测试，底层和 FailNow 一样调用 `runtime.Goexit` 执行

```
SkipNow()
Skipped() bool
```

> SkipNow 和 FatalNow 类似，但不会将测试视为失败

## 环境与目录相关方法

### Chdir

Chdir 将当前测试的工作目录切换到 dir，并在测试结束时自动恢复原工作目录

```
Chdir(dir string)
```

- 行为作用于当前测试上下文，而不是全局进程
- 结合并行测试是安全的（内部做了隔离），避免 `os.Chdir` 带来的数据竞争
- 底层调用的还是 `os.Chdir`，使用 `os.Chdir` 是全局修改，多个并行测试会互相踩，使用 `tb.Chdir` 则是 tb 级别的、安全的目录切换。

### TempDir

创建一个当前测试专用的临时目录，并在测试结束后自动删除

```
TempDir() string
```

- 每个测试 / 子测试有独立目录，不会互相污染
- 自动通过 Cleanup 注册删除操作，无需手动清理
- 对并行测试安全
- 是文件系统测试的推荐方式

### SetEnv

安全的修改环境变量，`testing.TB` 接口提供的环境变量隔离工具，主要用于解决 `os.Setenv` 在测试中可能遇到的问题：

- 全局共享：影响其它测试的环境，尤其是并行测试时
- 无法自动回复：可能污染进程状态
- 请理复杂：需要自己编写 defer 或 Cleanup

> 把 Setenv 放进 TB 接口，可以让单元测试、benchmark、fuzz 测试都能安全地修改环境变量

```
Setenv(key, value string)
```

**核心行为**

调用 Setenv 使用键值对进行环境变量赋值后：

1. 读取当前环境变量中的旧值（如果存在）
2. 立即调用 `os.Setenv(key, value)`
3. 自动注册 `t.Cleanup()`，在测试结束后会自动调用进行恢复环境变量旧值的恢复或删除

> 底层 Setenv 会为每个测试注册自己的 Cleanup，不需要手动写 defer os.Setenv(...) 或 Cleanup，不需要担心覆盖全局变量，不需要担心并行测试会造成互相污染

## 上下文相关方法

### Context

返回一个与当前测试生命周期绑定的 `context.Context` 上下文

```
Context() context.Context
```

- 测试结束后该 content 会被自动 cancel
- 如果设置了 `go test -timeout` 或测试本身有 deadline， 那么 context 也会带上对应的 Deadline

## 创建辅助函数相关方法

### Helper

`tb.Helper()` 用于标记一个函数为测试辅助函数（helper）

```
Helper()
```

当测试失败时，testing 框架会跳过被标记为 helper 的函数栈帧，将错误位置指向实际调用 helper 的测试代码位置，而不是 helper 内部。

作用：

- 提升错误定位的准确性
- 保持测试报错输出更清晰可读
- 为构建 assert / require 等辅助库提供基础能力