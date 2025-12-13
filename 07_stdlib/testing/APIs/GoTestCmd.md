# go test

> 这部分主要是看着文档跟 AI 一起学习的，因为直接按照使用经验总结的话比较不够全面，其次就是确实有些东西较难理解吧
> 文档：https://pkg.go.dev/cmd/go#hdr-Test_packages 还有 `go help testflag`

```
go test [build/test flags] [packages] [build/test flags & test binary flags]
```

直接运行 `go test` 默认会重新编译每个包，并把所有文件名匹配 `*_test.go` 的文件一起编译进去

这些额外的测试文件中可以包含：

- 测试函数（`TesxXxx`）
- 基准测试函数（`BenchmarkXxx`）
- 模糊测试（`FuzzXxx`）
- 示例函数（`ExampleXxx`）

每个包都会生成一个独立的测试二级制，比如 `go test ./pkg1 ./pkg2`，会分别为 `pkg1` 和 `pkg2` 构建一个测试程序并运行。

## 测试文件命名规则

注意文件名以 `_` 下划线或 `.` 点开头的文件都会被忽略，包括 `_test.go`、`.foo.go`、`_bar_test.go` 这种形式的文件名。

如果测试文件中声明的包名以 `_test.go` 结尾，那么这些测试文件会被当作一个单独的包来进行编译，之后再与注册十二进制链接一起运行，比如
`package mypkg_test` 这意味着：

- 这是一个外部测试
- 只能访问被测试包的导出 API
- 不能访问未导出标识符

go 工具会自动忽略命名为 `testdata` 的目录，因为需要用于存放测试所需的辅助数据，fuzz 测试中会把测试失败的信息也存入
`testdata/fuzz`

## `go vet` 检查行为

在构建测试二进制的过程中，`go test` 会对包及其测试源码运行 `go vet`，如果运行结果有问题，go test就会报告这些问题并且不会运行测试二进制。
默认只运行 `go vet` 的一部分高置信度检查项：`atomic`, `bool`, `buildtags`, `directive`, `errorsas`, `ifaceassert`,
`nilfunc`, `printf`, `stringintconv` 和 `tests`，可以通过 `go doc cmd/vet` 查看这些检查项，同时也可以使用 `-vet=off` 关闭
`go vet`或使用 `-ver=all` 进行开启。

所有测试输出和汇总信息都会被打印到 `go test` 命令的标准输出（stdout），即使测试代码把内容写到了标准错误（stderr），同时
`go test` 自己的标准错误（stderr）只用于输出构建测试时的错误。

在测试运行的环境中，`go test` 会把 `$GOROOT/bin` 放在 `$PATH` 的最前面。

## 两种运行模式

`go test` 有两种运行模式：

### 本地目录模式

local directory mode，本地目录模式，当 `go test`在没有指定任何包参数时触发（开发较常用），例如 `go test` 或 `go test -v`，
在这种模式下，`go test` 会编译当前目录下的包源码和测试文件，然后运行生成的测试二进制。

- 只测试当前目录这个包，也就是当前目录下必须有 `.go` 文件且这些 `.go` 文件属于同一个包
- 不涉及子目录，子目录不会被自动递归
- 不涉及包列表，因为就只测试当前目录这个包

注意这种模式下测试缓存是被禁用的，也就是每次都会重新运行测试，不会命中 `test cache`

在测试结束后，`go test` 会打印一行汇总信息，显示 测试状态（ok 或 FAIL) / 包名 / 耗时：`ok example.com/mypkg 0.123s`

### 包列表模式

package list mode，第二种模式为包列表模式，当 `go test` 使用了明确的包参数时会触发，例如：`go test math`、`go test ./...`
或 `go test .`（也属于包列表模式）

- `go test .` 表示测试当前目录对应的包，`.` 表示一个包路径参数，所以触发的时包列表模式
- `go test ./...` 表示从当前目录开始，递归测试所有子目录中的 Go 包

在这种模式下，`go test` 会对命令行中列出的每一个包分别进行编译和测试，也就是 `go test ./pkg1 ./pkg2` 会依次对 pkg1 和 pkg2
进行测试

- 如果某个包的测试通过，`go test` 之会打印最终的 ok 汇总行，而不会打印测试内部的输出
- 如果某个包的测试失败，`go test` 会打印该包的完整测试输出结果，包括 `t.Log`、`fmt.Println`、`panic` 等

注意 -bench 或 -v 标准会改变输出行为，如果使用这种标志，即使包测试通过，`go test` 也会打印完整输出

- `-bench` 需要输出 benchmark 结果
- `-v` 也就是 verbose 模式需要输出详细日志

在所有列出的包的测试都完成并输出结果之后，如果有任何一个包测试失败，`go test` 会在最后打印一个总体的 FAIL 状态。
也就是说：可能前面看到很多 ok，只要有一个包失败，最终结果就是 FAIL

## 测试缓存

仅在包列表模式下， `go test` 才会缓存成功的包测试结果，以避免不必要的重复执行

```
go test          # 本地目录模式 不走缓存
go test .        # 包列表模式 走缓存
go test ./...    # 包列表模式 走缓存
```

在某次测试的结果可以从缓存中恢复时，`go test` 不会再次运行二进制，而是直接重新显示之前的输出并且在汇总行中用 `(cached)`
替代原本显示的耗时。
比如：`ok example.com/mypkg (cached)` 这表示测试没有重新运行，结果直接来自缓存

本次运行使用的是相同的测试二进制，并且命令行中的所有 flag 都来自一个可缓存 flag 的受限集合时才算命中缓存：

1. 测试二进制必须相同
    - 包源码不变
    - 依赖不变
    - 构建方式不变
2. 只能使用可缓存的测试 flag：`-benchtime`/`-coverprofile`/`-cpu`/`-failfast`/`-fullpath`/`-list`/`-outputdir`/
   `-parallel`/`-run`/`-short`/`-skip`/`-timeout`/`-v`，如果命令中出现不在这个列表中的
   flag，缓存也不会使用，所以不想走缓存就加一个不再这个列表中的flag

> 最常用的不使用缓存是加一个 `-count=1`，最推荐和标准的写法


缓存的测试结果被视为执行时间为 0，因此，一个成功的包测试结果会被缓存并复用，而不受 `-timeout` 设置的影响。
即使设置了很小的 `-timeout`，只要缓存命中，测试就会瞬间成功，不会触发 `-timeout`

## `go test` 命令控制选项

### `-args`

`-args` 会把命令行中剩余的参数（`-args` 之后的内容），原封不动，不做解析的传递给测试二进制（包括 `-c` 、 `-json` 、 `-run`
等所有flag），因为这个 flag 会吃掉命令行剩余的所有参数，所以包列表如果存在名也必须写在 `-args` 之前，也就是 `-args`
通常写在命令最后面

`go test ./pkg -args foo bar baz` 等价于 `pkg.test foo bar baz`，`go test` 不再解析 `foo bar baz`，会完全交给测试程序自己处理（比如
`os.Args`）

### `-c`

`-c` flag 为只编译不运行，会把测试二进制编译为 `pkg.test`（位于当前目录下），但不运行它，其中 `pkg` 是包导入路径的最后一段，可以配合
`-o` flag 来修改生成的文件名或目标目录。

`go test -c ./mypkg` 会生成一个文件 `mypkg.test`，而不会执行测试

### `-exec xprog`

用指定程序运行测试二进制，使用 `xprog` 来运行测试二进制，行为和 `go run` 相同。

`go test -exec qemu-arm ./mypkg` 测试二进制不会直接运行，而是会执行 `qemu-arm mypkg.test`

### `-json`

把测试输出转为适合自动化处理的 JSON 格式，同时也会以 JSON 格式输出构建阶段的输出
`go test -json ./...` 输出的不是给人看的文本，而是：

```
{"Action":"run","Package":"example.com/pkg"}
{"Action":"pass","Package":"example.com/pkg","Elapsed":0.12}
```

### `-o file`

用于把测试二进制编译到指定的文件中，测试仍然会运行（除非同时制定了 `-c` 和 `-i`），如果 `file` 以斜杠结尾或指向一个已存在的目录，
那么测试二进制就会被写入该目录下的 `pkg.test` 文件，测试结束不会删除该程序。

`go test -o /tmp/mytest ./mypkg` 会生成测试二进制到 `/tmp/mytest`，默认测试仍然会执行（与 `-c` 的区别）

## 常用运行实例命令

- 运行模块下所有测试实例：go test <module name>/<package name> 用来运行某个 package 内的所有测试用例
- 运行当前 package 内的用例：`go test .`
- 运行子 package 内的用例： go test example/<package name> 或 go test ./<package name>
- 递归测试当前目录下的所有的 package：`go test ./...` 或 go test example/...

> 参考：https://my.oschina.net/renhc/blog/3016178

## 控制测试的参数

`go test` 命令既接受作用于 `go test` 命令本身的 flags，也接受作用于最终生成的测试二进制的 flags

- go 命令层决定如何构建，如何组织包测试
- 测试二进制层由 `testing` 框架解释，也就是如何运行测试示例

其中有一些 flags 用于控制性能剖析（profiling），并写出可供 `go tool pprof` 使用的执行 profile 文件，使用 `go tool pprof -h`
获取更多信息， 这里的 `-cpuprofile` / `-memprofile` / `-blockprofile` / `-mutexprofile` 这类参数输出的文件是给 `pprof`
用的，不是给人看的，`pprof` 的 `--alloc_space`、`--alloc_objects` 和 `--show_bytes` 选项可以用于控制这些信息（profile
数据）以何种方式展示。

> 这里意思是 pprof 工具自己的展示参数，不属于 `go test` 的 flag，意思是 profile 文件生成之后怎么看由 pprof 控制

### 决定测试类型 flags

下面这几个参数觉得了哪些 unit / benchmark / fuzz /example 会被选中，后面都是接正则表达式，都不改变包，只影响 `testing.M`
调度哪些函数

```
-bench regexp
-run regexp
-skip regexp
-list regexp
-fuzz regexp
```

#### `-run regexp`

选择 `TestXxx` / `ExampleXxx` / `FuzzXxx`，不会选中 `Benchmark`，正则匹配的对象是测试的标识符也就是测试函数名。
`func TestAdd(t *testing.T) {}` 对应标识符就是 `TestAdd`

`/` 是有特殊含义的，比如 `-run X/Y` 会被拆为 `[X] [Y]` 用于匹配 父测试/子测试，比如 `go test -run TestAdd/Negative`：

1. 先运行 `TestAdd`（先运行父测试是为了找到子测试，就是 `TestAdd` 本身没有断言也会运行一次）
2. 在 `TestAdd` 中寻找子测试 `Negative`
3. 只真正执行匹配的子测试

```
go test -run TestAdd           # 单个测试
go test -run Add               # 一类测试
go test -run TestAdd/Negative  # 子测试
go test -run .                 # 所有测试
```

#### `-bench regexp`

只选择基准测试 `BenchmarkXxx`，默认行为 `go test` 不会跑任何 benchmark，必须显式指定 `go test -bench .`

`/` 的含义和 `-run` 中 `/` 含义一样

```
func BenchmarkAdd(b *testing.B) {
    b.Run("Small", func(b *testing.B) {})
}
```

使用 `go test -bench BenchmarkAdd/Small`：

1. `BenchmarkAdd` 会先以 `b.N=1` 运行
2. 找到子 benchmark `Small` 测试
3. 真正运行 `Small`

```
go test -bench .               # 所有 benchmark
go test -bench Add             # 名字包含 Add
go test -bench BenchmarkAdd    # 单个 benchmark
```

#### `-fuzz regexp`

行为只选择 fuzz 测试 `FuzzXxx`，只能匹配一个包或匹配一个 `FuzzXxx`，所以不允许使用 `go test -fuzz .`

> test 执行顺序为：Test -> Benchmark -> Example -> Fuzz（真正 fuzzing）

```
go test -fuzz FuzzParse
go test ./mypkg -fuzz FuzzParse -fuzztime=10s
```

#### `-skip regexp`

反向选择器，运行所有不匹配正则的 `Test/Benchmark/Fuzz/Example`，匹配规则和 `-run` / `-bench` 一样， `/` 分层规则和匹配对象都一样

```
go test -skip Integration   // 跳过所有名字中包含 `Integration` 的测试
```

#### `-list regexp`

只列出测试示例列表，不运行，列出对象包括 `Test/Benchmark/Fuzz/Example`，有一个很重要的限制是不会列出子测试和子
benchmark基准测试（刚好对应上为什么在运行子测试前需要先运行父测试来寻找对应的子测试）

```
go test -list .
go test -list TestAdd
```

### 控制运行方式 flags

这一类 flags 决定了运行方式、并发、次数、顺序、超时等，不改变测试内容，改变执行节奏/并发/重复，与 `testing.T/B/F` 行为强相关

```
-benchtime t
-count n
-cpu list
-parallel n
-failfast
-timeout d
-shuffle off/on/N
-short
```

#### `-benchtime t`

控制每个 benchmark 运行多长时间（或运行多少次），只影响 `BenchmarkXxx`，默认是 1s 也就是一秒，Go 会自动调节 `b.N` 直到耗时
`>=t`

```
-benchtime 2s      # 按时间
-benchtime 100x    # 固定跑 100 次
```

```
go test -bench . -benchtime 3s
go test -bench BenchmarkAdd -benchtime 100x
```

#### `-count n`

把一个测试整体重复运行 n 次，作用于 `Test/Benchmark/Fuzz seed`，默认是 1，`Example` 永远只会跑一次，不作用于真正的 fuzzing(
`-fuzz`)，注意可以和 `-cpu` 进行组合，比如 `-count 2 -cpu 1,2` 实际运行次数就是 `2次 * 每个 GOMAXPROCS 值`

```
go test . -count=1  // 这是显式关闭测试缓存最标准的写法
```

#### `-cpu 1,2,4`

在不同的 GOMAXPROCS 值下重复执行测试，默认情况下 Go 会设置：`runtime.GOMAXPROCS(x)`，然后完整跑一边测试，默认只用当前
GOMAXPROCS

```
go test -cpu 1,2,4
go test -bench . -cpu 1,4
```

#### `-parallel n`

限制 `t.Parallel()` 能同时运行的最大测试数，只影响调用了 `t.Parallel()` 的测试，默认值就是 GOMAXPROCS，注意还可以写为 `-p`

- `-parallel` 为单个包内测试并发
- `-p` 为多个包之间的并发

> 在 fuzzing 阶段 `-parallel` = 同时运行的 fuzz 子进程数，有是否调用 `t.Parallel()`无关

#### `-failfast`

一旦有测试失败了，就不再进行新的测试，已经在跑的测试不会被中断，但是不会再进行新的测试调度

```
go test -failfast   // 主要用于 CI 和冒烟测试
```

#### `-timeout d`

限制整个测试二进制的最大运行时间，如果超时会触发 panic，默认为 10m 也就是十分钟，`-timeout=0` 表示关闭超时，永不超时，适用于所有
`Test / Benchmark / Example / Fuzz seed replay`

#### `-shuffle off/on/N`

打乱测试和 benchmark 的执行顺序，主要用于暴露顺序依赖 bug，找出隐藏的共享状态问题，三个可选值：

- off：默认
- on：用系统时间作为随机种子
- N：固定 seed（可复现）

```
go test -shuffle=on     // 开启
go test -shuffle=12345  // 可复现
```

#### `-short`

给测试一个现在是短模式的信号，本身不跳过任何测试，需要在测试代码中使用 `testing.Short` 进行判断

### 覆盖率相关

这几个决定是否插桩、如何统计、统计哪些包，影响编译阶段（插桩），会生成覆盖率数据

```
-cover
-covermode set,count,atomic
-coverpkg pattern1,pattern2,...
-coverprofile cover.out
```

#### `-cover`

开启覆盖率统计，并在测试输出中显示一个百分比。`go test -cover`示例输出：

```
ok   example.com/mypkg   0.123s   coverage: 78.3% of statements
```

- 只统计当前测试包
- 不生成测试文件
- 只是给一个大概数字

#### `-covermode`

决定覆盖率统计如何计算，决定粒度和准确性，一共有三种模式

1. `set` 默认值，只关心有没有执行过，不关心执行测试，最快，默认模式
2. `count` 统计每条语句执行次数，比 set 更精细，但在并发场景下不安全
3. `atomic` 使用原子操作机械能统计，并发安全，但是性能会明显下降

> 如果启用了 `-race`，默认 covermode 会自动变成 atomic

#### `-coverpkg`

指定哪些包的代码需要被统计覆盖率，比如 `go test ./service -cover` 默认只统计 service 包本身，即使 `service` 调用了 `repo`，
`repo` 的代码也被执行了，也不会算覆盖率，需要使用 `go test ./service -coverpkg=./repo,./model`，意思就是给 `repo` 和
`model` 进行插桩，覆盖率统计包含这些包。

#### `-coverprofile cover.out`

把覆盖率数据写成 profile 文件供工具分析，`go test -coverprofile=cover.out` 回自动启用 `-cover` 并生成 `cover.out`

### Fuzz 专用 flags

这几个 flags 只在 fuzz 测试中生效，只影响 `FuzzXxx`，对普通 `Test / Benchmark` 不生效，与 `testing.F` 直接对应

```
-fuzz regexp
-fuzztime t
-fuzzminimizetime t
```

> `-fuzz` 在上面这里不在讲解

#### `-fuzztime t`

控制真正的 fuzzing（变异+执行目标函数）这一阶段运行多长时间，Go 会不断调用 `FuzzXxx` 函数，自动调整执行次数直到类型耗时
`>=t`

```
-fuzztime 10s
-fuzztime 2m
-fuzztime 1h30s
-fuzztime 1000x
```

- t 为标准的 time.Duration
- 默认情况下不会停止，只能手动中断或遇到 failure
- t 也可以设为 1000x 表示不按照时间，按照固定次数

```
# 默认不会停止
go test -fuzz FuzzParse

# 本地快速试跑
go test -fuzz FuzzParse -fuzztime=10s

# CI 稳定 fuzz
go test -fuzz FuzzParse -fuzztime=2m

# 调试：只跑固定次数
go test -fuzz FuzzParse -fuzztime=100x
```

#### `-fuzzminimizetime t`

控制最小化失败输入（minimization）阶段，每一次尝试最多跑多久，当 fuzz 发现一个失败输入后，Go 会：

1. 把触发失败输入保存下来
2. 反复尝试缩小这个输入
3. 目标是更短更简单，但仍然能复现失败

这个过程就叫做 minimization

默认值为 60s 也就是一分钟，同样支持 `100x` 这样次数输入，也就是每一轮最小化只跑 100 次 fuzz 目标函数

```
# 默认就够用
go test -fuzz FuzzParse

# 想更快结束 minimization
go test -fuzz FuzzParse -fuzzminimizetime=10s

# CI 中控制上限
go test -fuzz FuzzParse -fuzztime=1m -fuzzminimizetime=10s
```

### 输出展示相关

这几个 flags 决定输出给人看还是给机器看，不影响测试逻辑，不影响是否通过，只影响能看到什么、怎么存

```
-v
-json
-fullpath
-outputdir directory
```

> json 上面也已经讲过了

#### `-v`

verbose 也就是详细模式，让每个测试的执行过程都可见，每个 `Test/Benchmark/Example` 开始结束都会打印，即使测试成功 `t.Log`/
`t.Lofg` 也会输出，如果没有添加 `-v` flag，`t.Log("hello")` 如果测试成功就不会显示，失败才会显示。

```
go test -v
go test -run TestAdd -v
go test ./... -v
```

#### `-fullpath`

在错误信息中显示完整的文件路径而不是相对路径，比如默认是 `now/now_test.go:12`，使用 `-fullpath` 后会变成
`C:/demo/GolangTraining/now/now_test.go:12`

#### `-outputdir directory`

指定 `profiling / trace / cover` 等输出文件的统一存放目录，其实放在 Profiling / Tracing（性能分析）讲解应该更好。

影响到文件包括：

```
-cpuprofile
-memprofile
-blockprofile
-mutexprofile
-trace
-coverprofile
```

默认是输出到当前运行 `go test` 的目录

`go test -bench . -cpuprofile cpu.out -outputdir ./profiles` 结果是 `./profiles/cpu.out`

### Profiling / Tracing（性能分析）

这几个 flags 生成 `pprof / trace` 文件，给工具分析用，
这一组 flag 的唯一目的是让测试在运行过程中顺便采集性能数据，并输出给专业工具分析。

- 不会改变测试是否通过
- 不会改变测试逻辑
- 主要给工具（pprof / trace）用，不是给人直接看的
- 大多数都会生成一个文件

```
# Benchmark
-benchmem

# CPU / 内存分配（pprof）
-cpuprofile cpu.out
-memprofile mem.out
-memprofilerate n

# 阻塞 / 锁竞争（pprof）
-blockprofile block.out
-blockprofilerate n
-mutexprofile mutex.out
-mutexprofilefraction n

# 覆盖率（profile 文件，上面覆盖率也讲过了）
coverprofile

# 执行轨迹（trace）
-trace trace.out
```

#### `-benchmem`

在 benchmark 结果中，额外打“内存分配统计。 当跑 benchmark 时：

```
go test -bench . -benchmem
```

输出会多出两列：

```
BenchmarkAdd-8    1000000    120 ns/op    16 B/op    1 allocs/op
```

- `B/op`：每次操作分配多少字节
- `allocs/op`：每次操作分配几次

> 不统计 C / C.malloc 的内存分配

#### `-cpuprofile cpu.out`

采集 CPU 使用情况的采样 profile，测试结束前写入文件，等价于：`pprof.StartCPUProfile`

```
go test -bench . -cpuprofile cpu.out
go tool pprof cpu.out
```

#### `-memprofile mem.out`

采集内存分配 profile（谁分配了内存）。测试全部通过后才写，统计的是分配来源

#### `-memprofilerate n`

控制采样精度 vs 性能开销，默认是采样，`n` 越小，越精确，越慢。极端模式：

```
-memprofilerate=1   // 记录每一次分配
```

#### `-blockprofile block.out`

采集 goroutine 阻塞（channel / select / sleep 等）的 profile。用于检查卡在哪，查 channel 阻塞

#### `-blockprofilerate n`

控制阻塞事件的采样频率。默认：`n = 1`（记录全部），值越大，采样越少，性能越好

#### `-mutexprofile mutex.out`

采集 mutex 锁竞争的 profile。用于检查哪个锁竞争最严重

#### `-mutexprofilefraction n`

对 mutex 持有栈做抽样。`n=1` 为全部记录，`n=10` 为采集 1/10

#### `-trace trace.out`

采集 Go 程序完整执行轨迹。这是信息量最大、也最重要的一种。可以观察：

- goroutine 生命周期
- 调度器行为
- GC
- syscalls
- 网络、channel、锁

```
go test -trace trace.out
go tool trace trace.out
```

- 文件可能很大
- 性能开销明显
- 不适合长期跑

## 补充

### Windows powershell 问题

`go test -bench .` 正常应该是等价于 `go test -bench=.` 但是 `Windows PowerShell` 里 `.`
有特殊含义（表示当前目录）所以第二个命令可能存在问题，可以这样写：`-bench='.'`

> `go test run .` 同理

### go test 正则表达式参数

> go test 里所有正则表达式类 flag，使用的都是 Go 标准库的 regexp（RE2）语法：https://pkg.go.dev/regexp/syntax