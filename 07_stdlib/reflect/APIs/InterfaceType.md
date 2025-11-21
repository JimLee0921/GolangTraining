# `reflect.Type`

`reflect.Type` 是反射中最基础的概念，是 Go 编译器生成的类型描述信息（Metadata），反射通过它知道一个值的类型结构是什么样的。

在 Go 源码中，Type 是一个接口，遮住了内部复杂的实现（不同类型对应不同底层 `struct`，如 `rtype`、`arrayType`、`structType` 等）。

## 简化版结构

```
type Type interface {
    Align() int
    FieldAlign() int
    Method(int) Method
    MethodByName(string) (Method, bool)
    NumMethod() int
    Name() string
    PkgPath() string
    Size() uintptr
    String() string
    Kind() Kind
    ...
}
```

只需要知道 `reflect.Type` 是一个接口，真正持有信息的是底层的 `rtype` 结构。

## 相关函数

前两个是用的比较多的顶层函数用于获取一个值或type的类型，后面几个主要用于动态创建一个新的 `reflect.Type`

### 1. `reflect.TypeOf`

这是最常用的入口函数，可以根据一个值返回它的类型，最常用、也是学 reflect 时接触最多的函数。

```
func TypeOf(i any) Type
```

- 需要传一个值（如果传入 nil interface 则返回 nil）
- 返回这个值实际的类型信息
- 是反射最基础的入口

> 空接口 + 空类型” 才是真的 nil interface

**常用于**

- 查看类型
- 遍历 struct 字段
- 获取类型的 Kind
- 获取 Pointer 的 Elem 类型
- ORM、序列化框架

### 2. `reflect.TypeFor`

Go 1.20+ 新增的类型构造函数（现代泛型友好），根据 泛型 推导类型，而不是值。
重点是不需要值也不需要实例，直接通过一个类型就看也获得 Type

```
func TypeFor[T any]() Type
```

- TypeFor 返回表示类型参数 T 的 类型
- 只需要一个类型不需要实例就看也获取 Type
- 常用于反射泛型，构造类型，一定要类型而非值

**对比 TypeOf**

| 特性               | TypeOf  | TypeFor            |
|------------------|---------|--------------------|
| 是否需要一个值          | 必须要     | 不需要                |
| 是否支持泛型推导         | 不支持     | ️ 完美支持             |
| 是否安全（不会被 nil 干扰） | 可能 nil  | 完全不会               |
| Go 版本            | Go 1.0+ | Go 1.20+           |
| 使用场景             | 反射已有变量  | 反射泛型，构造类型，一定要类型而非值 |

TypeOf 需要一个值，所以在很多场景下非常不方便，例如：

- 只知道类型，不想创建实例
- 在写泛型库，而 TypeOf 做不到泛型的类型推导

这时就建议直接使用 TypeFor。

### 3. `reflect.ArrayOf`

动态创建一个固定长度数组类型，例如想创建 `[5]int` 这种数组类型，但数量是运行时才能确定：

```
func ArrayOf(length int, elem Type) Type
```

可以用它创建：`[3]string`，`[10]float64`，`[N]User` 等
