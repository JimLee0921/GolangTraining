# 顶级函数

这四个是 reflect 包中的四个顶层函数级别 API，其它与 Type 和 Value 相关的函数在各自章节再学习

## `reflect.Copy`

Copy 函数会将 src 中的内容复制到 dst 中，直到 dst 被填满或 src 中的内容被复制完毕。
类似 `copy()` 内置函数，但操作对象是 Value 类型，只能作用于 slice / array。

```
func Copy(dst, src Value) int
```

- 返回复制的元素数量
- dst 和 src 的类型必须都为 Slice 或 Array 并且元素类型必须相同
- 数量由较短者决定，类似于 built-in copy
- 如果 dst 是Array，且 `Value.CanSet` 返回 false，则函数会引发 panic
- 并不是深拷贝，只是元素移动引用/值

## `reflect.DeepEqual`

DeepEqual 函数报告 x 和 y 是否深度相等，几乎能比较所有类型 （struct/map/slice/pointer/interface/基本类型 等）

主要用于比较两个值是否递归意义上相等，常用于调试，测试，不适用于生产环境，是 reflect 版本的 `==`，但是功能更加强大。

```
func DeepEqual(x, y any) bool
```

**判断规则如下：**

| 类型 / 情况                     | DeepEqual 判定相等的条件 / 行为                                                                                                         |
|-----------------------------|--------------------------------------------------------------------------------------------------------------------------------|
| **基础类型**（数字、布尔、字符串、channel） | 直接用 Go 的 `==` 运算符进行比较                                                                                                          |
| **Array（数组）**               | 数组长度必须一致，且每个对应元素也必须 DeepEqual                                                                                                  |
| **Struct（结构体）**             | 对应的所有字段（包括导出与非导出字段）必须都 DeepEqual                                                                                               |
| **Pointer（指针）**             | 两个 pointer 值若用 Go 的 `==` 相等 (即指向同一个地址） 就视为 equal，否则比较它们指向的值（递归 DeepEqual）                                                      |
| **Interface（接口）**           | 先检查接口中持有的具体值 (concrete value) 的类型和内容 — 如果具体值 deep-equal，则接口值 equal                                                             |
| **Map（映射）**                 | 要求两者都为 nil 或都非 nil；键数 (len) 要一致；要么是同一个 map 对象，要么对于每个 key (按 go 的相等性判断) 对应的 value 都 DeepEqual                                   |
| **Slice（切片）**               | 要求两者都为 nil 或都非 nil；长度一样；要么底层数组相同 (同一片, `&x[0] == &y[0]`)，要么逐元素 DeepEqual。注意 —— nil slice 与 non-nil 但 length 为 0 的 slice 被认为不相等 |
| **Func（函数）**                | 仅当两个函数值都为 `nil` 时，DeepEqual 才为 true；否则，不论它们是否行为相同，都被认为不相等                                                                      |

## `reflect.Swapper`

Swapper 返回一个函数，该函数交换所提供切片中的元素。如果提供的接口不是切片，Swapper 程序会崩溃。

完全用于 排序/洗牌/泛型 slice 操作，sort.Slice 内部就靠 Swapper 实现交换。

```
func Swapper(slice any) func(i, j int)
```

| 要点                        | 说明           |
|---------------------------|--------------|
| 无需转换泛型                    | 任何 slice 都支持 |
| 底层通过 `Value.Index.Set` 实现 | 等于动态反射交换     |
| 常用于 sort.Slice            | 交换器避免类型断言开销  |

## `reflect.TypeAssert`

Go1.20+ 新特性，用于从 `reflect.Value` 做泛型安全断言

```
func TypeAssert[T any](v Value) (T, bool)
```

老版本使用 `x := v.Interface().(int)` 可能会触发 panic，而 TypeAssert 会返回 (值, 是否成功)

| 优势          | 对比Interface.(T) |
|-------------|-----------------|
| 返回ok避免panic | 安全              |
| 支持泛型类型解构    | 更现代             |
| 无需接口断言→直拿T  | 速度更快，语义更清晰      |
