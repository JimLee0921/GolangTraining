# 单元测试 (Unit Test)

一个包中的 `_test.go` 文件会被 `go test` 自动发现并执行，他们会被编译为额外的测试二进制（不会包含在发布程序里）。

go 单元测试是为了验证代码中的最小可测试单元（通常是函数或方法）是否按预期工作，它的主要作用是提高代码质量、简化调试过程、确保代码重构的正确性，并在开发早期就发现和修复错误。
通过编写测试用例，开发者可以确保代码在各种情况下的正确性，从而提高软件的可靠性和可维护性。

## 语法要求

单元测试的函数必须满足 `func TestXxx(t *testing.T)` 格式，常用于验证函数逻辑是否正确

**要求：**

1. `Test` 开头
2. `Xxx` 是待测试的方法名，任意名称，但是必须首字母大写开头，否则无法被识别
3. 参数必须是 `*testing.T`
4. 返回值必须为 `void`

> Go 的测试框架在运行每个测试函数时，会自动创建并传入一个 `*testing.T` 对象

## *testing.T 参数

测试用的参数有且只有一个，在这里是 t *testing.T，是测试上下文对象（testing context），用来在测试函数中报告失败、打印日志、控制子测试等

**常用方法**

| 方法                                               | 说明                               |
|--------------------------------------------------|----------------------------------|
| `t.Log(args...)` / `t.Logf(format, args...)`     | 打印测试日志，仅在 `go test -v` 时显示       |
| `t.Error(args...)` / `t.Errorf(format, args...)` | 记录错误并标记测试为失败（但继续执行）              |
| `t.Fatal(args...)` / `t.Fatalf(format, args...)` | 记录错误并立即终止当前测试函数                  |
| `t.Fail()`                                       | 手动将测试标记为失败（不中断执行）                |
| `t.FailNow()`                                    | 标记失败并立即终止当前测试函数（相当于 `t.Fatal()`） |
| `t.Skip(args...)` / `t.Skipf(format, args...)`   | 跳过该测试                            |
| `t.Helper()`                                     | 标记当前函数为辅助函数，在错误日志中跳过此函数堆栈帧       |
| `t.Cleanup(func())`                              | 注册一个在测试结束时自动执行的清理函数              |
| `t.Parallel()`                                   | 允许当前测试与其他测试并行运行                  |
| `t.Run(name, func(t *testing.T))`                | 创建子测试（table-driven tests 常用）     |
| `t.TempDir()`                                    | 创建临时目录，会在测试结束后自动删除               |

## t.Error / t.Fatal 区别

Go 的测试库中经常使用两类断言方式

1. `t.Error` / `t.Errorf`：标记测试失败但继续执行，适用于不是致命的情况，就是只是做错误打印日志
2. `t.Fatal` / `t.Fatalf`：标记失败并立即中断当前测试函数，适用于无法继续执行的系统级别错误

## 基础结构示例

最基本的 test case，单元测试，如果有一个函数：

```
func Add(a, b int) int {
    return a + b
}
```

测试应该写成：

```
package mathx

import "testing"

func TestAdd(t *testing.T) {
    got := Add(1, 2)
    want := 3
    
    if got != want {
        t.Errorf("Add(1, 2) = %d; want %d", got, want)
    }
}
```

**注意事项**

1. 文件名必须是 `xxx_test.go`
2. 测试文件和测试函数所在属于一个 package （推荐做法）
3. 错误信息建议包含 `got / want`，方便快速定位