# `reflect.Value`

`reflect.Value` 是运行时对值本身的抽象，允许读取、修改、创建、调用方法。
如果 `reflect.Type` 是描述一类东西，那么 `reflect.Value` 就是操作那个东西的手。

reflect.Value 是反射中真正可操作的部分，支持：

- 获取值（int/string/bool/slice/map/struct/etc）
- 修改值（必须是可寻址的 addressable）
- 创建值（通过 reflect.New / reflect.MakeXXX）
- 调用方法（MethodByName）

# Value 相关函数

### 1. `reflect.ValueOf`

Value 反射世界的入口，传入任意值，返回一个 `reflect.Value`，用于操作运行时数据，里面保存：动态类型（Type）和动态值（Value）。
反射获取的是值，而不是类型，但是可以通过 `.Type` 随时获取类型。

```
func ValueOf(i any) Value
```

- ValueOf 返回的是只读副本，除非传入的是指针，否则不会改变原值，不会创建新值，只是把这个值包装成反射对象
- `ValueOf(&x)` 也不可以直接修改x，需要使用 `ValueOf(&value).Elem()` 才能获取被指向的实际值（x 本体）并进行修改

### 2. `reflect.New`

根据类型 t 创建一个新的指向零值的指针。主要用于

- 创建结构体实例（动态构造对象）
- JSON / ORM 中动态分配对象
- 用新对象填入 slice / map

```
func New(typ Type) Value
```

- 等价于 `new(T)`（返回 `*T`）
- 返回一个 Value，其类型是指向 typ 的指针

### 3. `reflect.NewAt`

很危险的函数，在指定内存地址 pointer 上创建一个 *类型是 t 的 Value。
这段内存存放的就是类型 t 的数据”，然后它会 wrap 成一个 *T

```
func NewAt(typ Type, p unsafe.Pointer) Value
```

## 动态创建容器类型函数

MakeXXX 开头的一组动态创建值函数。这些函数跟 `reflect.New` 不同的是：

- `*New` 创建的是 Type -> 指针
- MakeXXX 创建并返回一个实际的可用值（不一定是指针）

它们对应 Go 中的内建 `make()` 函数，仅用于 slice / map / chan / func（特殊）

### 1. `reflect.MakeSlice`

动态创建一个 Slice，类似于 `make([]T, len, cap)`，可用于动态创建 T 类型切片，长度可变，可修改元素。

```
func MakeSlice(typ Type, len, cap int) Value
```

### 2. `reflect.MakeMap`

动态创建 map，相当于 `make(map[K]V)`，适用于动态构造 map，在反射序列化、ORM、解码 JSON 中非常常用。

```
func MakeMap(typ Type) Value
```

### 3. `reflect.MakeMapWithSize`

与 MakeMap 相同，只是可以提供初始容量，相当于 `make(map[K]V, size)`，适用于批量写入大 map，减少扩容次数。

```
func MakeMapWithSize(typ Type, n int) Value
```

### 4. `reflect.MakeChan`

动态创建 chan，相当于 `make(chan T, buf)`，可用于泛型消息队列、调度中心、测试 mock 时动态创建 channel。

```
func MakeChan(typ Type, buffer int) Value
```

### 5. `reflect.MakeFunc`

比较特殊，可以动态生成一个函数值（闭包），并可以自定义函数体。会返回一个类型为 typ 的函数对象，fn 会作为函数逻辑替代执行。

```
func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value
```

## 其它操作函数

这些函数多数用于 切片操作、Channel 选择、指针/间接引用、动态创建零值。

### 1. `reflect.Append`

向一个 slice Value 追加元素，类似于 `append(slice, elem1, elem2...)`，仅用于 Slice 类型 Value

```
func Append(s Value, x ...Value) Value
```

### 2. `reflect.AppendSlice`

将一个 slice 追加到另一个 slice，类似于 `append(s, t...)`，适合批量追加，注意两个 slice 的元素类型必须一致，否则 panic

```
func AppendSlice(s, t Value) Value
```

### 3. `reflect.Indirect`

返回 v 指向的值，如果传入的是指针 Value，则返回其 `Elem()`，如果 v 是空指针，则 Indirect 函数返回零值，若不是指针，则原样返回
v 的值。

```
func Indirect(v Value) Value
```

### 4. `reflect.Select`

高级函数，常用于channel调度。

类似原生 `select { case ... }` 语句，用于 多个 channel 的选择，会阻塞，直到至少有一个 case 可以执行。
返回所选 case 的索引，如果该 case 是接收操作，则返回接收到的值以及一个布尔值，该布尔值指示该值是否对应于通道上的发送操作（而不是由于通道已关闭而接收到的零值）。
`SELECT` 语句最多支持 65536 个 case。

- Send：发送
- Recv：接收
- Default：非阻塞

```
func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)
```

### 5. `reflect.SliceAt`

高级函数，但是和 NewAt 一样需要 unsafe，用于在一个底层数组的指定下标位置创建 Slice 视图。

```
func SliceAt(typ Type, p unsafe.Pointer, n int) Value
```

SliceAt 返回一个值，该值表示一个切片，其底层数据从 p 开始，长度和容量等于 n，就像不安全切片一样。

### 6. `reflect.Zero`

返回类型 T 的零值 Value，主要用于

- 动态创建默认值
- 清空字段
- 实现 DeepCopy / Reset / Pool

```
func Zero(typ Type) Value
```

# 常用方法

Value 的方法整体可分为 7 大类：

1. 取值 / 设值族
2. 类型与元信息
3. Slice / Array 专属方法
4. Map 专属方法
5. Chan / Concurrency 方法
6. Struct 字段读写与标签(Tag)
7. Func 调用与动态执行

## 取值 / 设值 方法

## 常见 CanXXX 方法

这一层不做操作，主要获取 Value 是否可以进行某些操作。

### CanAddr / CanSet

这两个是理解反射修改的核心

- CanAddr 返回改该 Value 是否可以取地址
- CanSet 返回该 Value是否可以被修改（`Setxxx`方法）

能修改并不代表肯定可以寻找地址，但通常需要先可以寻找地址才能进行修改

```
func (v Value) CanAddr() bool
func (v Value) CanSet() bool
```

### 其它 CanXXX 方法

```
func (v Value) CanInterface() bool
func (v Value) CanConvert(t Type) bool
func (v Value) CanInt() bool
func (v Value) CanComplex() bool
func (v Value) CanFloat() bool
func (v Value) CanUint() bool
```

- `CanInterface`：是否能 `.Interface()` 导出为普通 Go 类型
- `CanConvert`：报告值 v 是否可以转换为类型 t。如果 `v.CanConvert(t)` 返回 true，则 `v.Convert(t)` 不会引发 panic
- `CanInt()` / `CanUint()` / `CanFloat()` / `CanComplex()`：是否可以用相应 Getter 读取（避免 panic）

## Meta 元信息 Getter 查询方法

Struct 原信息单独放在 Struct 讲解

### Kind

基本分类，用于在反射时判断走哪种分支。
Kind 返回 v 的 Kind 属性。如果 v 的值是零（`Value.IsValid`返回 false），则 Kind 返回 Invalid。

```
func (v Value) Kind() Kind
```

> Kind 是粗分类，相当于：Int / Uint / Float / Bool / String / Struct / Slice / Map / Array / Ptr / Interface / Chan /
> Func ...

### Type

返回 v 的完整类型信息（含包名、字段、方法）

```
func (v Value) Type() Type
```

> Kind 是基础类型颗粒度，主要用于分支判断，而 Type 是全类型描述结构，主要用于字段/方法分析

### 状态类查询

主要包含 空值、nil、合法性等判断

```
func (v Value) IsNil() bool
func (v Value) IsValid() bool
func (v Value) IsZero() bool
```

- IsZero：判断 v 是否为空值，常用于做 JSON/ORM 可选字段判断
- IsNil：仅适用于判断引用类型（slice / map / chan / pointer / func / interface）是否为空
- IsValid：用于判断 `v` 是否代表一个值。当 `FieldByName` 找不到字段，`MapIndex(key)` 时 key 不存在或者数组 Slice 月结，都会返回
  `IsValid()=false` 的 Value 而不是 panic。常用于反射框架编写。

## Value 常用 Getter 方法

这些 Getter 方法主要用于取值

### 引用 Getter 方法 `Elem`

Elem 方法返回获取指针、interface 中实际存放的值，而 `reflect.Indirect` 函数是 Elem 的使用简化版。

```
func (v Value) Elem() Value
```

- 如果 v 的 Kind 不是 Interface 或 Pointer ，则会引发 panic
- 如果 v 为 nil，则返回零值

> `reflect.Indirect` 函数等价于 `if ptr → Elem() else → v`，优势是无需判断是否是指针，一步到位，ORM/JSON 场景很常用。

### 通用 Getter 方法 `Interface`

`Interface` 方法把 `reflect.Value` 还原成一个普通的 any，是所有 Getter 里最万能的。

```
func (v Value) Interface() (i any)
```

返回 v 的当前值，以 `interface{}` 的形式返回。它等价于：`var i interface{} = (v 的底层值)`

**常用于**

- 做一个通用的处理函数，所有字段都用 `Interface()` 拿出来再 `fmt.Printf`、转成 JSON 等
- `Interface()` + 类型断言：`val.(int)`、`val.(string)` 等

配套的是`CanInterface()`，有些值（比如私有字段）不能直接 `Interface()`，否则 panic。

### 标量 Getter

强类型版 Getter，要求当前 Value 的种类（Kind）是对应或可转换的类型，否则 panic，安全写法一般会配合 `CanInt()`/`CanUint()`/
`CanFloat()`/`CanComplex()`

```
func (v Value) Int() int64
func (v Value) Float() float64
func (v Value) String() string
func (v Value) Bool() bool
func (v Value) Uint() uint64
func (v Value) Complex() complex128
```

- `Int()`：适用于有符号整数类型（int、int8…），返回 int64
- `Uint()`：适用于无符号整数类型，返回 uint64
- `Float()`：适用于 float32/float64，返回 float64
- `Complex()`：适用于 complex64/complex128，返回 complex128
- `Bool()`：适用于布尔值，返回 bool，如果 v 的 kind 不是 Bool 类型也会引发 panic
- `String()`：适用于字符串，返回字符串 v 的底层值，以字符串形式返回

> String 方法比较特殊，由于 Go 语言的 `String` 方法约定，`String` 方法是一个特例。与其他 getter 方法不同，如果 `v` 的 `Kind`
> 不是`String` ，它不会引发 panic 。相反，它会返回一个形如 `<T value>` 的字符串，其中 `T` 是 `v` 的类型。`fmt` 包对 `Value`
> 类型进行了特殊处理。它不会隐式调用 `String` 方法，而是打印出它们所包含的具体值。

### 特殊 Getter

`Bytes()` 和 `Pointer()` 方法使用不多，比较特殊

Bytes 返回 v 的底层值。如果 v 的底层值不是字节切片或可寻址的字节数组，也就是 `v.Kind()` 必须为 Slice，否则该方法会引发
panic。反射里处理 `[]byte` 类型字段（比如数据库 BLOB、网络包等）。

Pointer 方法返回 v 的值，类型为 uintptr。

- 如果 v 的 Kind 不是 Chan、Func、Map、Pointer、Slice、String 或 UnsafePointer 会引发 panic
- 如果 v 的 Kind 类型为 Func ，则返回的指针是指向底层代码的指针，但这并不一定足以唯一地标识单个函数。 唯一能保证的是，当且仅当
  v 为 nil func 值时，结果才为零
- 如果 v 的 Kind 类型为 Slice，则返回的指针指向切片的第一个元素。如果切片为 nil，则返回值为 0。如果切片为空但非 nil，则返回值为非零值
- 如果 v 的 Kind 是 String，则返回的指针指向字符串底层字节的第一个元素

```
func (v Value) Bytes() []byte
func (v Value) Pointer() uintptr
```

> Pointer 本质是暴露内部地址，用错了非常容易出大问题，一般配合 unsafe 做底层 hack
> 主要用于自己管理内存、和 C 交互等场景，在正常业务逻辑里，尽量少用。

### 容器 Getter

这些主要用于容器结构一些信息的获取

```
func (v Value) Index(i int) Value
func (v Value) MapIndex(key Value) Value
func (v Value) Len() int
func (v Value) Cap() int
```

- Index(i)： 返回 v 的第 i 个元素，如果 `v.Kind()` 不是 `slice/array/string` 或下标越界会触发 panic
- MapIndex(key)：返回映射 v 中与键 key 关联的值，`v.Kind()` 必须是 Map，未找到返回对应 Value 类型的零值
- Len()：返回 v 的长度，如果 `v.Kind()` 不是 `Array/Chan/Map/Slice/String` 或指向 Array 的指针会引发 panic
- Cap()： 返回 v 的容量，如果 `v.Kind()` 不是 `Array/Chan/Slice` 或指向 Array 的指针会引发 panic

## 基础 Setter

注意反射要修改值，必须传指针，再 `Elem()` 才能 Set

### 通用 Set

Set 方法用于将另一个 Value 替换当前值。Set 方法不能自动转类型，所以类型必须完全一致 否则 panic

```
func (v Value) Set(x Value)
```

Set 方法将 x 赋值给 v。如果 `Value.CanSet` 返回 false，则会引发 panic。
x 的值必须是 v 的类型，并且不能派生自未导出的字段。

### 基本类型专用 Setter

```
func (v Value) SetInt(x int64)
func (v Value) SetFloat(x float64)
func (v Value) SetBool(x bool)
func (v Value) SetString(x string)
func (v Value) SetUint(x uint64)
func (v Value) SetZero()
func (v Value) SetComplex(x complex128)
```

- SetInt：将 v 的底层值设置为 x，如果 `v.Kind()` 不是 `Int/Int8/Int16/Int32/Int64`，或者 `Value.CanSet` 返回 false，则会引发
  panic
- SetFloat：将 v 的底层值设置为 x。如果 `v.Kind()` 不是 `Float32/Float64`，或者 `Value.CanSet` 返回 false，则会引发 panic
- SetComplex：SetComplex 将 v 的底层值设置为 x，如果 `v.Kind()` 不是 `Complex64/Complex128`，或者 `Value.CanSet` 返回
  false，则会引发 panic
- SetString：将 v 的底层值设置为 x。如果 `v.Kind()` 不是String，或者 `Value.CanSet` 返回 false，则会引发 panic
- SetUint： 将 v 的底层值设置为 x。如果 `v.Kind()` 不是 `Uint/Uintptr/Uint8/Uint16/Uint32/Uint64`，或者 `Value.CanSet` 返回
  false，则会引发 panic
- SetZero：将 v 的值设置为 v 类型的零值。如果 `Value.CanSet` 返回 false，则会引发 panic

### 其它基础 Setter

这两个很少在业务中进行使用，但在底层操作 / 网络缓冲 / C 交互时非常强力。

```
func (v Value) SetBytes(x []byte)
func (v Value) SetPointer(x unsafe.Pointer)
```

- SetPointer：修改指针指向的地址，将 `unsafe.Pointer` 类型的值 v 设置为 x。如果 `v.Kind()` 类型不是 UnsafePointer ，或者
  `Value.CanSet` 返回 false，则该方法会引发 panic
- SetBytes：用于设置变量 v 的底层值。如果 v 的底层值不是字节切片，或者 `Value.CanSet` 返回 false，则该方法会引发 panic。

