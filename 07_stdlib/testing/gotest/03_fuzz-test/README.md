# 模糊测试 (Fuzz Test)

Go 1.18+ 支持，模糊测试就是让 Go 自动生成大量非人工设计的输入，持续尝试把程序搞崩。

它主要解决三类问题：

- panic / crash
- 边界条件错误（空值、极端长度、非法编码）
- 不变量被破坏（`decode(encode(x)) != x`）

> Fuzz 不是随机测试，而是反馈驱动的输入搜索

## 基本结构

```
func FuzzXxx(f *testing.F)
```

## Seed 种子输入

Fuzz 测试每次都需要先使用 `f.Add()` 方法进行种子输入，因为 Fuzz 并不是从零开始乱猜，而是：从给的 Seed 开始逐步变异，保留能触发新路径的输入

```
func FuzzReverse(f *testing.F) {
    f.Add("hello")
    f.Add("")
    f.Add("你好")
    f.Add("🙂")

    f.Fuzz(func(t *testing.T, s string) {
        rr := Reverse(Reverse(s))
        if rr != s {
            t.Fatalf("failed")
        }
    })
}
```

**经验法则**

- 至少一个正常值
- 一个空值
- 一个非 ASCII 码值
- 一个极端情况（比如长字符串）

## 运行

使用 `go test -fuzz` 后面加具体函数名或 `.` 进行模糊测试

1. 从 Seed 开始
2. 自动生成变体
3. 覆盖率反馈（coverage-guided）
4. 找到 panic/fatal
5. 如果失败保存失败用例到 testdata/fuzz/

> 这是一个长期运行工具，不是一跑就停，如果不加 `go test -fuzz` 不添加 `-fuzztime` 参数默认会一直跑下去直到遇到错误或手动停止，这一分钟包含模糊搜索阶段

## 对比单元测试 Test

双方是互补的而不是替代关系

- Test 是人写用例，覆盖能想到的情况，主要为了验证功能正确性
- Fuzz 是机器生成用例，覆盖想不到的情况是为了验证健壮性和安全性