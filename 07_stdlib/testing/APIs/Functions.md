# testing 顶层函数

开发适合大部分基本用不到，这里主要了解

## AllocsPerRun

测试某段代码平均每场执行分配了多少从内存

```
func AllocsPerRun(runs int, f func()) (avg float64)
```

- benchmark 的轻量版本，
- 用于单元测试中验证有没有多于分配，快速回归检查

## Short

判断当前是否使用了 `-short` 模式，`go test -short` 指定的

```
func Short() bool
```

## Verbose

判断当前是否开启了 verbose 模式，`go test -v` 指定的

```
func Verbose() bool
```

## Coverage

返回当前包的覆盖率（0.0 ~ 1.0），只有开启了 `-cover` 时才有意义，否则返回 0

```
func Coverage() float64
```

## CoverMode

返回当前覆盖率模式：`set/count/atomic`，`go test -covermode` 传入的

```
func CoverMode() string
```

## Testing

判断当前代码是否运行在 testing 环境中，使用 `go test`、`fuzz/benchmark`、Example 执行期间为 true

```
func Testing() bool
```

## Init

更不可能用得到，用于初始化 testing 包的内部状态，`go test` 生成的 `test main` 进行调用

```
func Init()
```

## Main

testing 框架的总入口，`go test` 自动生成的 main，用不到

```
func Main(
    matchString func(pat, str string) (bool, error),
    tests []InternalTest,
    benchmarks []InternalBenchmark,
    examples []InternalExample,
)
```

## RunTests/RunBenchmarks/RunExamples

这三个是执行 `TestXxx`/`Benchmark`/`Example` 的调度器，testing 内部引擎，不是 API，用不到

```
func RunBenchmarks(matchString func(pat, str string) (bool, error), ...)
func RunExamples(matchString func(pat, str string) (bool, error), examples []InternalExample) (ok bool)
func RunTests(matchString func(pat, str string) (bool, error), tests []InternalTest) (ok bool)
```

## RegisterCover

给 testing 框架注册一个覆盖率统计器，根本用不到

```
func RegisterCover(c Cover)
```