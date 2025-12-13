# go testing

> 文档地址：https://pkg.go.dev/testing

Go 语言的标准库 testing 进行单元测试。
在 Go 语言中，`go test` 是官方内置的测试命令，用于运行单元测试、基准测试（benchmark）、模糊测试（fuzz）以及示例测试。
`go test` 会自动查找并运行以 `_test.go` 结尾的文件中的测试函数。

Go 语言推荐测试文件和源代码文件放在一块，测试文件以 `_test.go` 结尾。
比如，当前 package 有 `calculate.go` 一个文件，想测试 `calculate.go` 中的 Add 和 Mul 函数，那么应该新建 `calculate_test.go`
作为测试文件。

```
example/
   |--calculate.go
   |--calculate_test.go
```

> 测试文件必须以 *_test.go 结尾。并且测试文件中不能携带 func main()

## 测试类型

常见的几种测试类型：

| 类型   | 函数签名                              | 用途                | 运行命令             |
|------|-----------------------------------|-------------------|------------------|
| 单元测试 | `func TestXxx(t *testing.T)`      | 测试某个函数或功能是否正确     | `go test`        |
| 基准测试 | `func BenchmarkXxx(b *testing.B)` | 测试代码性能（执行速度、内存分配） | `go test -bench` |
| 模糊测试 | `func FuzzXxx(f *testing.F)`      | 测试可能意外遇到的 bug     | `go test -fuzz`  |
| 示例测试 | `func ExampleXxx()`               | 提供示例代码，验证输出并生成文档  |                  |

## 其他类型

### 子测试

子测试是一个测试函数内部动态创建的独立测试单元。单元测试和基准测试都可以通过参数 t/b 创建子测试，示例测试函数签名为
`func ExampleXxx()`，所以示例测试函数不能调用 `t.Run()` 创建子测试。

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

- 不要返回错误， 帮助函数内部直接使用 `t.Error` 或 `t.Fatal` 即可，在用例主逻辑中不会因为太多的错误处理代码，影响可读性
- `t.Helper()` 只会影响 通过 `t.Error*` / `t.Fatal*` / `t.Fail*` 等方法产生的错误位置归因（把报错行归到调用者那里）
- 对运行时 panic 的堆栈 没有抹除效果

### 表格驱动测试

这是 Go 测试的核心风格，在 Go 的测试中，大部分专业库基本都采用表格驱动方式(table-driven tests)。

所有用例的数据组织在切片 cases 中，看起来就像一张表，借助循环创建子测试。

这样写的好处有：

1. 新增用例非常简单，只需给 cases 新增一条测试数据即可
2. 测试代码可读性好，直观地能够看到每个子测试的参数和期待的返回值
3. 用例失败时，报错信息的格式比较统一，测试报告易于阅读
4. 如果数据量较大，或是一些二进制数据，推荐使用相对路径从文件中读取

把多个测试用例（输入 + 期望输出）放在一个表（通常是切片）中，用循环去执行这些用例。
就是用一张数据表来驱动测试逻辑，而不是手动写一堆重复的测试代码。

#### 思想

表驱动（table-driven）并不是 Go 独有的术语，它来自通用的软件工程思想：

- 把逻辑决策或测试数据放进“表格（数据结构）
- 用循环或规则去执行，而不是写死在代码里

#### 示例

```
func TestAdd(t *testing.T){
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"simple", 1, 2, 3},
        {"negative", -1, 2, 1},
        {"zero", 0, 0, 0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T){
            if got != tt.want {
                t.Errorf("Add(%d %d) = %d; want %d", tt.a, tt.b, got, want)
            }
        })
    }
}
```

> 然后用 for _, c := range cases { ... } 来执行测试

| 优点         | 说明                                |
|------------|-----------------------------------|
| **简洁可读**   | 所有测试数据集中在一个地方                     |
| **方便维护**   | 新增或修改用例只需改“表”，不需改逻辑               |
| **易于扩展**   | 可轻松增加几十个测试场景                      |
| **结合子测试**  | 可用 `t.Run(c.name, ...)` 单独运行、并行执行 |
| **自动报告清晰** | 子测试名称自动出现在输出里                     |

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

### 实现方式

| 场景           | 推荐做法                               | 说明                                             |
|--------------|------------------------------------|------------------------------------------------|
| 单个测试前后       | `t.Cleanup()` 或 `defer teardown()` | 简单明了                                           |
| 多个测试共用 setup | 单独封装函数 + `t.Helper()`              | 代码复用                                           |
| 全包测试前后       | `TestMain(m *testing.M)`           | 类似 Python 的 `setup_module` / `teardown_module` |

#### t.Cleanup()

Go 1.14+ 推荐使用，`(*testing.T).Cleanup() / (*testing.B).Cleanup()` 是 Go 内置的清理机制，
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

#### 手动 setup/teardown 函数（通用）

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

#### 测试集级别的 setup/teardown（全局）

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

