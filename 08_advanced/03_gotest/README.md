# go testing

Go 语言的标准库 testing 进行单元测试。
在 Go 语言中，`go test` 是官方内置的测试命令，用于运行单元测试、基准测试（benchmark）以及示例测试。
`go test` 会自动查找并运行以 `_test.go` 结尾的文件中的测试函数。

Go 语言推荐测试文件和源代码文件放在一块，测试文件以 _test.go 结尾。
比如，当前 package 有 calculate.go 一个文件，想测试 calculate.go 中的 Add 和 Mul 函数，那么应该新建 calculate_test.go
作为测试文件。

```
example/
   |--calculate.go
   |--calculate_test.go
```

> 测试文件必须以 *_test.go 结尾。并且测试文件中不能携带 func main()

## 测试类型

常见的三种测试类型：

| 类型   | 函数签名                              | 用途                |
|------|-----------------------------------|-------------------|
| 单元测试 | `func TestXxx(t *testing.T)`      | 测试某个函数或功能是否正确     |
| 基准测试 | `func BenchmarkXxx(b *testing.B)` | 测试代码性能（执行速度、内存分配） |
| 示例测试 | `func ExampleXxx()`               | 提供示例代码，验证输出并生成文档  |

## 其他类型

### 子测试

子测试是一个测试函数内部动态创建的独立测试单元。单元测试和基准测试都可以通过参数 t/b 创建子测试，示例测试函数签名为
`func ExampleXxx()`，所以不能调用 `t.Run()` 创建子测试。

#### 创建子测试

- 单元测试：通过调用`t.Run(name string, f func(t *testing.T))` 来创建多个子测试
- 基准测试：b *testing.B（基准测试的上下文对象）也有同样的 Run 方法，调用`b.Run(name, func(b *testing.B))`创建子测试
- 示例测试函数签名为`func ExampleXxx()`，所以不能调用 `t.Run()` 创建子测试。

### 帮助函数 t.Helper()

对一些重复的逻辑，抽取出来作为公共的帮助函数(helpers)，可以增加测试代码的可读性和可维护性。 借助帮助函数，可以让测试用例的主逻辑看起来更清晰。

在测试中，经常会写一些重复使用的断言逻辑，例如检查结果是否相等、是否返回错误等。
如果直接调用这些函数，错误日志会指向辅助函数本身，而不是实际出错的测试代码行。
这会导致调试困难。

在测试函数中加上 `t.Helper()` 告诉测试框架：这个函数是辅助函数，不要在错误日志中显示它的调用栈，可以让报错信息更准确，有助于定位

#### 配合子测试

- 用子测试组织不同测试场景
- 用帮助函数封装断言逻辑
- 在帮助函数中调用 t.Helper() 让错误定位清晰

#### 注意事项

- 不要返回错误， 帮助函数内部直接使用 t.Error 或 t.Fatal 即可，在用例主逻辑中不会因为太多的错误处理代码，影响可读性
- t.Helper() 只会影响 通过 t.Error* / t.Fatal* / t.Fail* 等方法产生的错误位置归因（把报错行归到调用者那里）
- 对 运行时 panic 的堆栈 没有抹除效果

### 表驱动测试（Table-driven tests）

把多个测试用例（输入 + 期望输出）放在一个表（通常是切片）中，用循环去执行这些用例。
就是用一张数据表来驱动测试逻辑，而不是手动写一堆重复的测试代码。

#### 思想

表驱动（table-driven）并不是 Go 独有的术语。
它来自通用的软件工程思想：

- 把逻辑决策或测试数据放进“表格（数据结构）
- 用循环或规则去执行，而不是写死在代码里

#### go 实现

```text
cases := []struct {
    name string
    input int
    want int
}{
    {"case1", 1, 2},
    {"case2", 3, 6},
}
```

> 然后用 for _, c := range cases { ... } 来执行测试

| 优点         | 说明                              |
|------------|---------------------------------|
| **简洁可读**   | 所有测试数据集中在一个地方                   |
| **方便维护**   | 新增或修改用例只需改“表”，不需改逻辑             |
| **易于扩展**   | 可轻松增加几十个测试场景                    |
| **结合子测试**  | 可用 `t.Run(c.name, …)` 单独运行、并行执行 |
| **自动报告清晰** | 子测试名称自动出现在输出里                   |

## setup & teardown

在 Go 测试中，setup（测试前准备） 和 teardown（测试后清理） 是一种测试生命周期管理模式，
用于在每个测试或整个测试集运行前后，自动做一些初始化与清理工作。

### 使用场景

**setup**

- 打开数据库连接
- 初始化配置文件
- 创建临时目录或文件
- 启动一个模拟服务器

**teardown**

- 关闭连接
- 删除临时文件
- 恢复全局变量

有下面三种实现方式：

### t.Cleanup()

Go 1.14+ 推荐，`(*testing.T).Cleanup() / (*testing.B).Cleanup()` 是 Go 内置的清理机制，
可以注册一个在测试结束后自动执行的函数。

```
func TestSomething(t *testing.T) {
    // setup: 创建临时目录
    dir := t.TempDir()
    file := filepath.Join(dir, "data.txt")
    os.WriteFile(file, []byte("hello"), 0644)

    // teardown: 自动清理（即使测试失败也会执行）
    t.Cleanup(func() {
        fmt.Println("清理资源：", dir)
        os.RemoveAll(dir)
    })

    // 测试逻辑
    data, _ := os.ReadFile(file)
    if string(data) != "hello" {
        t.Fatalf("expected hello, got %s", data)
    }
}
```

- 无论测试成功或失败，Cleanup 都会执行
- 可以注册多个 t.Cleanup()，会按注册的逆序执行
- 自动处理并发和子测试

### 手动 setup/teardown 函数（通用）

```
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
    defer teardown(t, dir)

    // 测试逻辑
}
```

- `defer teardown(...)` 确保测试结束自动清理
- 可以在多个测试中复用
- 与 t.Helper() 配合使用，错误定位准确

### 测试集级别的 setup/teardown（全局）

如果想在整个包的所有测试前/后做准备或清理，可以实现特殊函数：

```
func TestMain(m *testing.M) {
    fmt.Println("全局 Setup：连接数据库")
    code := m.Run() // 运行所有测试
    fmt.Println("全局 Teardown：关闭数据库")
    os.Exit(code)
}
```

- TestMain 是包级入口
- setup 和 teardown 也可以封装为函数进行调用
- 调用 m.Run() 触发所有测试用例的执行，并使用 os.Exit() 处理返回的状态码，如果不为0，说明有用例失败
- 在调用 m.Run() 前后做一些额外的准备(setup)和回收(teardown)工作
- 常用于数据库、外部服务、全局初始化

### 对比

| 场景           | 推荐做法                               | 说明                                             |
|--------------|------------------------------------|------------------------------------------------|
| 单个测试前后       | `t.Cleanup()` 或 `defer teardown()` | 简单明了                                           |
| 多个测试共用 setup | 单独封装函数 + `t.Helper()`              | 代码复用                                           |
| 全包测试前后       | `TestMain(m *testing.M)`           | 类似 Python 的 `setup_module` / `teardown_module` |

在 Go 里，官方内置的测试体系（全部基于 testing 包）主要三种：


## go test

### 运行实例

- 运行模块下所有测试实例：go test <module name>/<package name> 用来运行某个 package 内的所有测试用例
- 运行当前 package 内的用例：go test example 或 go test .
- 运行子 package 内的用例： go test example/<package name> 或 go test ./<package name>
- 递归测试当前目录下的所有的 package：go test ./... 或 go test example/...

> 参考：https://my.oschina.net/renhc/blog/3016178

### 控制编译的参数

- `-args`
    - 指示 `go test` 把 `-args` 后面的参数带到测试中去。具体的测试函数会跟据此参数来控制测试流程
    - `-args` 后面可以附带多个参数，所有参数都将以字符串形式传入，每个参数做为一个 `string`，并存放到字符串切片中
- `-json`
    - `-json` 参数用于指示 `go test` 将结果输出转换成json格式，以方便自动化测试解析使用
- `-o <file>`
    - `-o` 参数指定生成的二进制可执行程序，并执行测试，测试结束不会删除该程序
    - 没有此参数时，`go test` 生成的二进制可执行程序存放到临时目录，执行结束便删除

### 控制测试的参数

- `-bench regexp`
    - go test默认不执行性能测试，使用-bench参数才可以运行，而且只运行性能测试函数
    - 其中正则表达式用于筛选所要执行的性能测试。如果要执行所有的性能测试，使用参数 `-bench .`或 `-bench=.`
    - 正则表达式不是严格意义上的正则，而是种包含关系
- `-benchtime <t>s`
    - `-benchtime` 指定每个性能测试的执行时间，如果不指定，则使用默认时间1s
    - 例如，执定每个性能测试执行 `2s`，则参数为 `go test -bench Sub/A=1 -benchtime 2s`
- `-cpu 1,2,4`
    - `-cpu` 参数提供一个 `CPU` 个数的列表，提供此列表后，那么测试将按照这个列表指定的CPU数设置 `GOMAXPROCS` 并分别测试
    - 比如 `-cpu 1,2`，那么每个测试将执行两次，一次是用1个CPU执行，一次是用2个CPU执行
- `-count n`
    - `-count` 指定每个测试执行的次数，默认执行一次
    - 如果使用 `-count` 指定执行次数的同时还指定了 `-cpu` 列表，那么测试将在每种 CPU 数量下执行 count 指定的次数
    - 示例测试不关心 `-count` 和 `-cpu` 参数，它总是执行一次
- `-failfast`
    - 默认情况下，`go test` 将会执行所有匹配到的测试，并最后打印测试结果，无论成功或失败
    - `-failfast`指定如果有测试出现失败，则立即停止测试。这在有大量的测试需要执行时，能够更快的发现问题
- `-list regexp`
    - `-list` 只是列出匹配成功的测试函数，并不真正执行。而且，不会列出子函数
- `-parallel n`
    - 指定测试的最大并发数
    - 当测试使用 `t.Parallel()` 方法将测试转为并发时，将受到最大并发数的限制
    - 默认情况下最多有 `GOMAXPROCS` 个测试并发，其他的测试只能阻塞等待
- `-run regexp`
    - 根据正则表达式执行单元测试和示例测试。正则匹配规则与 `-bench` 类似
- `-timeout DURATION`
    - 默认情况下，测试执行超过10分钟就会超时而退出
    - 可以超时时间设置为1s，由本来需要3s的测试就会因超时而退出：`go test -timeout=1s`
    - 设置超时可以按秒、按分和按时：

        - 按秒设置：-timeout xs或-timeout=xs
        - 按分设置：-timeout xm或-timeout=xm
        - 按时设置：-timeout xh或-timeout=xh
- `-v`
    - 默认情况下，测试结果只打印简单的测试结果，`-v` 参数开启 verbose 模式，可以打印详细的日志
    - 性能测试下，总是打印日志，因为日志有时会影响性能结果
    - 在 `TestXxx` 和 `BenchmarkXxx` 中可以使用 t/b.Log，然后使用 `verbose` 模式时就会打印日志
- `-benchmem`
    - 默认情况下，性能测试结果只打印运行次数、每个操作耗时
    - 使用 `-benchmem` 则可以打印每个操作分配的字节数、每个操作分配的对象数

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

### t *testing.T vs b *testing.B

| 功能   | `t *testing.T`                    | `b *testing.B`                         |
|------|-----------------------------------|----------------------------------------|
| 用途   | 功能正确性测试                           | 性能测试                                   |
| 参数   | 测试函数 `func TestXxx(t *testing.T)` | 基准测试 `func BenchmarkXxx(b *testing.B)` |
| 控制循环 | 无                                 | 有（`b.N`）                               |
| 时间统计 | 无                                 | 自动测量耗时                                 |
| 内存统计 | 无                                 | 可选（`b.ReportAllocs()` / `-benchmem`）   |
| 并行运行 | `t.Parallel()`                    | `b.RunParallel()`                      |
| 子测试  | `t.Run()`                         | `b.Run()`                              |
| 清理函数 | `t.Cleanup()`                     | `b.Cleanup()`                          |