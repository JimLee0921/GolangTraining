# `testing.F`

`testing.F` 是 Go 的模糊测试（Fuzz Testing）类型，出现在 Go 1.18+。

模糊测试的目标是为了自动生成不同的输入以触发代码中未考虑的边界情况、错误分支、panic、崩溃、死循环等问题。

它是覆盖率驱动的，这意味着 fuzz engine 会：

- 观察哪些输入能增加程序覆盖率
- 自动调整输入
- 不断生成更有效地测试用例

最终可以找出：

- panic / 崩溃
- 未处理的边界
- 多步组合逻辑的漏洞
- 字符串解析错误
- JSON/XML/协议解析缺陷
- 类型转换缺陷
- 加密/压缩算法的异常情况

## 定义

很简单的定义

```
type F struct {
	// contains filtered or unexported fields
}
```

## 测试模板

最基本的模糊测试模板如下：

```
func FuzzXxx(f *testing.F){
    // 1. 添加种子输入（必须）
    f.Add("hello")
    
    // 2. 注册 fuzz 函数
    f.Fuzz(func(t *testing.T, input string)){
        // 对 input 进行测试
    }
}
```

对比单元测试和基准测试

| 类型                           | 签名        | 作用         |
|------------------------------|-----------|------------|
| `TestXxx(t *testing.T)`      | 单次固定输入    | 只跑一次       |
| `BenchmarkXxx(b *testing.B)` | b.N 循环    | 多次固定输入     |
| `FuzzXxx(f *testing.F)`      | f.Fuzz 调度 | 自动生成无限随机输入 |

## 核心方法

当然 `testing.F` 也是实现了 `testing.TB` 接口，这里就不再讲解那些方法，属于 `testing.F` 的主要有两个方法：`Add()` 和
`Fuzz()`

### 1. Add

添加初始种子输入（seed corpus），这些输入会：

1. 在 fuzz 前先执行一次（类似于单元测试）
2. 作为 fuzz 引擎的起点

```
func (f *F) Add(args ...any)
```

> 必须至少 Add 一组输入，否则 T 型函数无法推断参数类型

### 2. Fuzz

注册 fuzzing 函数，引擎会先用 `Add()` 的输入跑一次，然后开始自动生成成千上万的随机输入

```
func (f *F) Fuzz(fn func(t *testing.T, args ...any))
```

- fn 的第一个参数永远是 `testing.T`
- 剩余参数的类型由 `Add()` 决定
- 每次 Fuzz 调用相当于一个独立的测试用例
