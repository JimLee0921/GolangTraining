# reflect.Kind

在 reflect 包里，Kind 是一个枚举常量类型，代表 Go 语言类型的基础种类／类别，描述的是一个类型所属的大类。

```
type Kind uint
```

所有通过 `reflect.TypeOf(...)` / `reflect.ValueOf(...)` 拿到的 Type 或 Value，都可以通过 `.Kind()` 方法，获得这个类型/值的
Kind。

```
const (
	Invalid Kind = iota // 非法/未定义
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)
```

不论用自定义类型 (type Foo int)、结构体、slice、map、channel、函数、接口、指针等最后都能用 Kind 将其归一化到一类基础分类。

## 示例

定义一个类型：`type MyInt int`

- 对一个 MyInt 值调用 Type() -> 得到的 Type 是 `main.MyInt`
- 但是 Kind() -> 得到的是 `reflect.Int` ，底层种类仍是 int

## Kind 与 Type

| 属性        | Type                                    | Kind                                   |
|-----------|-----------------------------------------|----------------------------------------|
| 表示对象      | 语言中具体定义的类型 (可能有别名、struct、指针、map、slice…) | 底层类别 (int, slice, map, struct, ptr, …) |
| 是否稳定 / 唯一 | 唯一（包括别名／自定义）                            | 有限（属于 reflect.Kind 常量集中）               |
| 用途        | 查看字段 / 方法 / tag / 元信息                   | 快速判断类别 → 决定要用哪类反射 API                  |

- Kind 可以看作大类／族群(比如 字符串类 / 整型类 / slice类 / map类)
- Type 则是具体型号／SKU(比如 `main.MyInt` / `[]string` / `map[string]int` / `*MyStruct` 等)

## 重要性

go 中使用 reflect 反射前必须先看 Kind。

在写 reflect 代码时，一般的套路是：

1. 拿到一个 `reflect.Value v = reflect.ValueOf(...)`
2. `switch v.Kind()`，根据 Kind 来判断面对的是什么类型类别
3. 对应类别调用不同方法/操作

```
switch v.Kind() {
case reflect.Int, reflect.Float64, reflect.String:
    // 基础类型 → 用 v.Int()/v.Float()/v.String()
case reflect.Slice:
    // 切片 → 用 v.Len(), v.Index(i)…
case reflect.Struct:
    // 结构体 → 用 v.NumField(), v.Field(i)… / TypeOf + FieldByName + Tag…
case reflect.Ptr:
    // 指针 → v.Elem() 取实际值，再 recurse
case reflect.Map:
    // 映射 → v.MapKeys(), v.MapIndex(key)…
default:
    // 其他，或 panic
}
```

> 这种结构是 reflect 编程的基础，Kind 是第一层判断