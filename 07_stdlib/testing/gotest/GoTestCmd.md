# go test

> 这部分主要是看着文档跟 AI 一起学习的，因为直接按照使用经验总结的话比较不够全面，其次就是确实有些东西较难理解吧

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
-shuffle off,on,N
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

### Fuzz 专用 flags

这几个 flags 只在 fuzz 测试中生效，只影响 `FuzzXxx`，对普通 `Test / Benchmark` 不生效，与 `testing.F` 直接对应

```
-fuzz regexp
-fuzztime t
-fuzzminimizetime t
```

### 输出展示相关

这几个 flags 决定输出给人看还是给机器看，不影响测试逻辑，不影响是否通过，只影响能看到什么、怎么存

```
-v
-json
-fullpath
-outputdir directory
```

### Profiling / Tracing（性能分析）

这几个 flags 生成 `pprof / trace` 文件，给工具分析用

```
-benchmem

-cpuprofile cpu.out
-memprofile mem.out
-memprofilerate n

-blockprofile block.out
-blockprofilerate n

-mutexprofile mutex.out
-mutexprofilefraction n

-trace trace.out
```

## 补充

### Windows powershell 问题

`go test -bench .` 正常应该是等价于 `go test -bench=.`
但是 Windows PowerShell 里 . 有特殊含义（表示当前目录）所以第二个命令可能不会执行 bench
可以这样写：-bench='.'

Go 的 flag 解析器会把 -bench 当成一个参数，后面必须跟 正则表达式
所以需要先指定路径，再加 -bench或者先加 -bench，再指定路径（注意 -bench 后面要有值）
go test -v ./06_testing/03_benchmark-test -bench=.
go test -v -bench=. ./06_testing/03_benchmark-test
benchmark 中有个 b.N 属性，如果该用例能够在 1s 内完成，b.N 的值便会增加，再次执行。b.N 的值大概以 1, 2, 3, 5, 10, 20, 30,
50, 100 这样的序列递增，越到后面，增加得越快

### go test 正则表达式参数

go test 中的参数可以传入正则表达式时这个正则表达式不是严格意义上的正则，而是种包含关系。

比如有如下三个性能测试：

```

func BenchmarkMakeSliceWithoutAlloc(b *testing.B)
func BenchmarkMakeSliceWithPreAlloc(b *testing.B)
func BenchmarkSetBytes(b *testing.B)

```

使用参数`-bench=Slice`，那么前两个测试因为都包含`Slice`，所以都会被执行，第三个测试则不会执行。

对于包含子测试的场景下，匹配是按层匹配的。举一个包含子测试的例子：

```

func BenchmarkSub(b *testing.B) {
b.Run("A=1", benchSub1)
b.Run("A=2", benchSub2)
b.Run("B=1", benchSub3)
}

```

测试函数命名规则中，子测试的名字需要以父测试名字做为前缀并以`/`连接，上面的例子实际上是包含4个测试：

```

Sub
Sub/A=1
Sub/A=2
Sub/B=1

```

如果想执行三个子测试，那么使用参数`-bench Sub`。如果只想执行`Sub/A=1`，则使用参数`-bench Sub/A=1`。如果想执行`Sub/A=1`和
`Sub/A=2`，则使用参数`-bench Sub/A=`
