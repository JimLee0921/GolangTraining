# `reflect.Value`

`reflect.Value` 是运行时对值本身的抽象，允许读取、修改、创建、调用方法。
如果 `reflect.Type` 是描述一类东西，那么 `reflect.Value` 就是操作那个东西的手。

reflect.Value 是反射中真正可操作的部分，支持：

- 获取值（int/string/bool/slice/map/struct/etc）
- 修改值（必须是可寻址的 addressable）
- 创建值（通过 reflect.New / reflect.MakeXXX）
- 调用方法（MethodByName）

# Value 相关函数 / 方法

## 获取 & 构造 Value / 零值

1. 顶层函数（从类型或普通值得到 Value）

    - ValueOf：从普通值得到 Value，一切反射入口，nil 返回 `Value{}`（IsValid=false）
    - Indirect：如果是指针则返回 `Elem()`，否则原样返回，等于自动解指针版 Elem

2. 按 Type 动态构造 Value

    - New：等价于 `reflect.ValueOf(new(T))`，返回一个指向类型 typ 零值的 Value（kind 为 Ptr）
    - Zero：构造类型 typ 的零值 Value（非指针）
    - NewAt：用给定内存地址 p，伪装成类型 typ 的指针 Value（高危，需保证内存布局）

### `reflect.ValueOf`

Value 反射世界的入口，传入任意值，返回一个 `reflect.Value`，用于操作运行时数据，里面保存：动态类型（Type）和动态值（Value）。
反射获取的是值，而不是类型，但是可以通过 `.Type` 随时获取类型。

```
func ValueOf(i any) Value
```

- ValueOf 返回的是只读副本，除非传入的是指针，否则不会改变原值，不会创建新值，只是把这个值包装成反射对象
- `ValueOf(&x)` 也不可以直接修改x，需要使用 `ValueOf(&value).Elem()` 才能获取被指向的实际值（x 本体）并进行修改

### `reflect.Indirect`

Value 自动解引用工具，安全包装版 Elem。

```
func Indirect(v Value) Value
```

如果是指针，则返回 `Elem()`，不是指针则原样返回。 如果 v 是空指针，则 Indirect 函数返回零值，等价于：

```
if v.Kind() == Ptr && !v.IsNil():
    return v.Elem()
else:
    return v
```

**和 `Elem()` 的区别**

| `Elem()`   | `Indirect()`      |
|------------|-------------------|
| 非指针会 panic | 非指针原样返回           |
| 常用于明确知道是指针 | 常用于不确定是不是指针       |
| 不处理 nil 安全 | 遇 nil 会返回 invalid |

### `reflect.New`

根据 Type 创建该类型的指针 Value，底层类似于 `reflect.ValueOf(new(T))`

```
func New(typ Type) Value
```

根据类型 t 创建一个新的指向零值的指针（`Kind()` 为 Ptr，指向 typ 的空值）

**注意事项**

1. 永远返回指针类型的 Value
    - 如果 typ 是 int，则结果是 `*int`
    - 如果 typ 是 struct，则结果是 `*struct`
2. 内部自动分配内存并填充为类型零值
    - int -> 0
    - struct -> 字段全部为零值
    - slice/map/pointer -> nil
3. 常用于动态构建对象用于编码/反序列化
    - ORM/JSON/XML/YAML 底层常用它 new 一个空对象再填值

### `reflect.Zero`

构造 非指针 的零值 Value。

```
func Zero(typ Type) Value
```

- 返回类型 T 的零值 Value，`Kind()` 与 `typ.Kind()` 一致
- Zero 是直接返回类型 T 的零值，而 New 是返回 `*T`，实际值在 Elem 中
- Zero 不可直接 Set（除非拿到可寻址副本）
- 只提供值，不提供地址，不适合构建可写模型

### `reflect.NewAt`

最底层、最需要慎重使用的构造型方法。使用既存内存地址 p，将它作为 typ 的指针包装成 Value

```
func NewAt(typ Type, p unsafe.Pointer) Value
```

很危险的函数，要求 p 指向的内存布局完全符合 typ

**注意事项**

1. 不会分配内存。传入的内存必须已经存在，否则悬空指针，也就是未定义行为
2. 常见用途：
    - 反射访问任意内存区域（FFI、共享内存、C结构体）
    - 实现 unsafe 动态类型映射
    - Zero-copy 序列化框架 / 性能机制代码常用
3. p 必须指向大小与对齐规则满足 typ 的对象，否则GC/逃逸/数据损坏皆可能

## 类型信息 / 基本元信息

这些方法/函数主要用于查询类型、结构、有效性、零值状态、可否转回 interface、是否能进行某些方法操作（Canxxx系列）等

- `(v Value) Type() Type`：返回静态类型（`reflect.Type`）
- `(v Value) Kind() Kind`：返回底层 kind（Int、Struct、Slice...），用于分支逻辑
- `(v Value) IsValid() bool`：判断该 `Value` 是否有效（比如找不到字段时返回的零 `Value` 就是 invalid）
- `(v Value) IsZero() bool`：判断该值是否为零值（全部字段为零）
- `(v Value) Interface() any`：把 `Value` 还原成 `interface{}`；是退出反射的标准方式
- `(v Value) Equal(u Value) bool`：等价性比较
- `(v Value) Comparable() bool`：类型比较
- `(v Value) IsNil() bool`：Value 是否为 nil 空值
- `(v Value) CanAddr() bool`：可寻址判定 API
- `(v Value) CanSet() bool`：能否修改当前 Value 对应的底层值（需要：可寻址 + 导出等条件）
- `(v Value) CanInterface() bool`：有些受保护的值（如未导出字段）不允许 `Interface()`，这里可提前检测
- `(v Value) CanConvert(to Type) bool`：是否可以类型转换
- `(v Value) CanInt() / CanUint() / CanFloat() / CanComplex()`：是否可以使用对应 Getter 取值

### `Type()`

返回 v 的返回静态类型 `reflect.Type`（含包名、字段、方法），类型层级比 Kind 更细。

```
func (v Value) Type() Type
```

- 会区分指针，`*int` `int` 是不同 Type
- 可以配合 `New/Zero/MakeSlice` 等从 Type 创建实例
- 常用于动态结构体解析、序列化、ORM字段映射

### `Kind()`

返回底层 Kind 枚举类型，如 `Int/Struct/Map/Slice...`，不带复杂结构，见 [TypeKind.md](TypeKind.md)

```
func (v Value) Kind() Kind
```

- 如果 v 的值是零（`Value.IsValid` 返回 false），则 Kind 返回 Invalid
- 主要用于在反射时判断走哪种分支，在写反射框架时第一个判断句（先 Kind 分支，再 Type 深解析）

| Type             | Kind   |
|------------------|--------|
| `map[string]int` | Map    |
| `[]byte`         | Slice  |
| `*User`          | Ptr    |
| `User`           | Struct |

### `IsValid()`

判断 Value 是否有效，也就是 `v` 是否代表一个值，无效 Value 是反射中的 Null

```
func (v Value) IsValid() bool
```

- 找不到字段/方法、ValueOf(nil)、MapIndex Miss 这些情况会返回 False
- IsValid 是唯一安全检查入口，任何后续操作前必须先检查
- 需要先判断 IsValid 在进行操作 Kind/Type，invalid Value 取 Kind/Type 会 panic

| 来源               | 示例行为         |
|------------------|--------------|
| ValueOf(nil)     | 直接得到 invalid |
| FieldByName 不存在  | 返回 invalid   |
| MapIndex(key不存在) | 返回 invalid   |
| 接口为 nil 且不是指针    | 也可能 invalid  |

### `IsZero()`

判断值是否为类型零值

```
func (v Value) IsZero() bool
```

- struct 要求所有字段为零才 true（深度递归判断）
- map/slice/ptr 为 nil 返回 true
- 非 nil 但空 slice 返回 false
- JSON 自定义序列化会常用到

- IsNil：仅适用于判断引用类型（slice / map / chan / pointer / func / interface）是否为空

### `IsNil`

判断 v 的值是否为 nil，相当于检查底层是否为 nil

- `v.Kind()` 必须为 Chan/Func/Interface/Map/Ptr/Slice/UnsafePointer 等类型支持 nil的
- Int/Bool/Struct/Array/String 调用 IsNil 会导致 panic

```
func (v Value) IsNil() bool
```

### `CanInterface()`

是否允许 `Interface()` 导出为普通 Go 类型，涉及可见性（导出/未导出字段）

```
func (v Value) CanInterface() bool
```

- 未导出字段大多返回 false（reflect取结构体时最常遇到）
- 如果 false 且调用 `Interface()` 会导致 panic
- 不等于 CanSet

### `Interface()`

把 `reflect.Value` 还原成一个普通的 any

```
func (v Value) Interface() (i any)
```

以 `any` 的形式返回 v 的当前值，它等价于：`var i interface{} = (v 的底层值)`

- 若 v 是指针，返回 interface 也是指针
- 若 `v.IsValid` 为 true，调用 Interface 会导致 panic
- 若 `v.CanInterface` 为 false，调用页会 panic

> ValueOf -> Value -> Interface() -> interface{}恢复是反射最完整闭环

### `Equal()`

Go 1.20 引入的新方法，和 Go 的 `==` 比较一样，但反射保证：类型必须相同且值必须可比较 (`v.Comparable() == true`)

```
func (v Value) Equal(u Value) bool
```

- 若两者类型不同或 v 不可比较（`v.Comparable()==false`）会导致 panic
- Equal 与 `v.Interface()==u.Interface()` 并不完全相同，因为 `Interface()`可能 panic（比如未导出字段）

### `Comparable()`

判断该 Value 的类型是否可用于 Go 的 `==` 比较，常用于调用 Equal 前必须检查：`if v.Comparable() { ... }`或反射中判断某字段是否可以作为
map key

```
func (v Value) Comparable() bool
```

| Kind                       | 是否可比较          |
|----------------------------|----------------|
| Bool / Int / Uint / Float  | 可比较            |
| String                     | 可比较            |
| Array                      | 可比较（元素必须可比较）   |
| Struct                     | 可比较（所有字段必须可比较） |
| Ptr / Chan / UnsafePointer | 可比较（地址比较）      |
| Slice                      | 不可比较           |
| Map                        | 不可比较           |
| Func                       | 不可比较           |

### `CanAddr()`

判断 Value 是否有实际内存地址，是 CanSet 的前置条件，当然地址存在也不是说一定可以写入

```
(v Value) CanAddr() bool
```

- 反射能否修改 = 是否能找到底层实体 = 必须有地址
- 没地址 -> 没权限写 -> CanSet 直接 false
- 只要 Value 来自变量本身，而不是一个临时副本，它就可寻址

### `CanSet()`

判断 Value 是否可修改（是否允许 Set），CanSet == CanAddr && 可写

```
func (v Value) CanSet() bool
```

- 必须 可寻址 + 不是未导出字段
- SetXXX 可能会 panic，需要使用 CanSet 进行提前检查
- Go 反射永远不会去修改一个没有实体地址的值，即便它是变量类型也不行

`CanAddr()==true` 是 `CanSet()` 的必要条件，但不是充分条件，还必须满足：

- 是可导出的字段（大写）
- 类型匹配
- 非 map 索引
- 非接口内部不可写值

### `CanConvert()`

判断 v 是否能够安全 Convert 成另一个类型（不做转换，只判断），允许转换的情况：

- 数字类型之间的类型转换
- string -> `[]byte` 不可以（不是 convert，而是语义不同）
- 结构体之间必须 field 可 assign 或 tag 合法

```
func (v Value) CanConvert(t Type) bool
```

### `CanInt()` / `CanUint()` / `CanFloat()` / `CanComplex()`

用于检测能否用对应 Getter 取值

```
func (v Value) CanInt() bool
func (v Value) CanUint() bool
func (v Value) CanFloat() bool
func (v Value) CanComplex() bool
```

- `v.CanInt()` 为true -> 允许 `v.Int()`
- `v.CanFloat()` 为true -> 允许 `v.Float()`

## 基本类型读取 & 写入方法

1. 读取（只读，不改变原值）
    - `(v Value) Bool() bool`：仅 Kind 为 Bool
    - `(v Value) Int() int64`：Kind 为 Int 系列
    - `(v Value) Uint() uint64`：Kind 为 Uint 系列 + Uintptr
    - `(v Value) Float() float64`：Kind 为 Float32/Float64
    - `(v Value) Complex() complex128`：Kind 为 Complex64/Complex128
    - `(v Value) String() string`：Kind 为 String
    - `(v Value) Bytes() []byte`：Kind 为 Slice 且 elem 为 `uint8`（`[]byte`）
    - `(v Value) Pointer() uintptr`：Kind 为 Ptr/Map/Chan/Func/UnsafePointer 等指针型
    - `(v Value) UnsafePointer() unsafe.Pointer`：和 `Pointer` 类似但直接给 `unsafe.Pointer`
    - `(v Value) Addr() Value`：返回当前值的地址
    - `(v Value) Elem() Value`：取指向的值
    - `func (v Value) UnsafeAddr() uintptr`：返回底层不安全指针地址

2. 溢出检测：判断某个数值是否可以安全表示为 `v` 的类型（例如往 `int8` 里塞 128 会溢出）

    - `(v Value) OverflowInt(x int64) bool`
    - `(v Value) OverflowUint(x uint64) bool`
    - `(v Value) OverflowFloat(x float64) bool`
    - `(v Value) OverflowComplex(x complex128) bool`

3. 写入/类型转换（可变，受 CanSet / 可寻址 限制）
    - `(v Value) Set(x Value)`：通用赋值，要求类型可赋值，用于非基本类型也可以
    - `(v Value) SetBool(x bool)`
    - `(v Value) SetInt(x int64)`
    - `(v Value) SetUint(x uint64)`
    - `(v Value) SetFloat(x float64)`
    - `(v Value) SetComplex(x complex128)`
    - `(v Value) SetString(x string)`
    - `(v Value) SetBytes(x []byte)`
    - `(v Value) SetPointer(x unsafe.Pointer)`
    - `(v Value) SetZero()`
    - `(v Value) Convert(t Type) Value`

### `Bool()`

读取一个 bool 值。

```
(v Value) Bool() bool
```

- 只允许 `Kind()` 为 Bool 的值，如果 `v.Kind() != reflect.Bool` 会触发 panic
- 返回对应的 bool
- 不关心 `CanSet / CanInterface`，只要 Kind 对就能读
- 可用于结构体字段、slice 元素、map 值等，只要底层是 bool

### `Int()`

读取有符号整数。

```
(v Value) Int() int64
```

- 允许的 Kind：Int/Int8/Int16/Int32/Int64，如果 `v.Kind()` 不符合会触发 panic
- 返回值统一转为 int64
- 不会溢出，因为底层值已经是某个 intN，转成 int64 一定能表示
- 只是读取，不做任何转换检查，如果是要把一个新值塞进去才要考虑溢出需要使用 OverflowInt

### `Uint()`

读取无符号整数（包括 uintptr）

```
(v Value) Uint() uint64
```

- 允许的 Kind：Uint/Uint8/Uint16/Uint32/Uint64/Uintptr，如果 `v.Kind()` 不符合会触发 panic
- 返回值统一转为 uint64
- 同样不会溢出，纯读取，常用来处理各种 size 的无符号字段，比如结构体中的 `uint32 / uintptr`

### `Float()`

读取浮点数

```
(v Value) Float() float64
```

- 允许的 Kind：Float32/Float64，同样 `t.Kind()` 不符合会触发 panic
- 返回值统一转换为 float64
- 用 float64 是为了方便，之后如果要再写回小精度浮点，再配合 OverflowFloat 检查

### `Complex()`

读取复数

```
(v Value) Complex() complex128
```

- 允许的 Kind：Complex64/Complex128，同样 `t.Kind()` 不符合会触发 panic
- 返回值统一转换为 complex128
- 复杂度很低，Go 里复数不常用，知道存在就行

### `String()`

读取字符串

```
(v Value) String() string
```

- 只允许 `Kind()` 为 String
- 返回值为 string
- Kind 不是 String 会导致 panic `[]byte`、`[]rune` 都不行

> 对于 `[]byte`，要用的是 `Bytes()` 这一组，而不是 `String()`

### `Bytes()`

读取一个 `[]byte`

```
(v Value) Bytes() []byte
```

- `Kind()` 必须为 Slice，并且 elem 类型必须是 uint8（也就是 []byte），否则触发 panic
- 返回值是 `[]byte`，一般是与底层数据共享的切片（不是 copy）
- 常用于处理二进制字段、I/O Buffer 等
- 因为返回的是共享底层的 []byte，对这个 []byte 的修改会影响原值（这是设计上的性能考虑）

### `Pointer()`

以 uintptr 形式返回某些引用类型的指针值

```
(v Value) Pointer() uintptr
```

- `Kind()` 必须为： Chan/Func/Map/Ptr/Slice/UnsafePointer/或者对应类型值为 nil 时也可（返回 0）
- 这个值并不保证稳定且与 GC 交互非常敏感，不能把它当 C 风格裸指针长期保存
- 如果只是要正常读写值，不建议用 Pointer/unsafe，一般配合 `unsafe.Pointer`、NewAt 等高级用法才需要

### `UnsafePointer()`

直接暴露 `unsafe.Pointer`

```
(v Value) UnsafePointer() unsafe.Pointer
```

- `Kind()` 必须为 UnsafePointer
- 跟 `Pointer()` 不同，它是直接返回 Go 层的 `unsafe.Pointer`，属于完全裸奔模式
- 一般只有在已经在用 unsafe 包写底层库时才会接触，普通业务代码几乎不该用

### `Addr()`

返回当前值的地址（即 &v）

```
(v Value) Addr() Value
```

- 返回类型是 Value，其 Kind 一定是 Ptr 指针
- 调用 Addr 方法的前提条件是必须可寻址，也就是 `CanAddr()==true`（只能用于原始值不是拷贝的时候）
- 主要用于得到 `*T` 类型的 Value，用于传给函数或 SetPointer

### `Elem()`

取指向的值，如 `*T` -> T、`interface{}` -> 其内部值

```
(v Value) Elem() Value
```

- 反射中访问指针指向的值的关键方法，返回被只想的 Value（不是拷贝，而是实际的值）
- 如果 `v.Kind()` 不是 Ptr 或 Interface 会导致 panic
- 如果 `nil pointer` 或 `nil interface`，Elem 返回零 Value(invalid)，而不会 panic，但 `invalid Value` 再访问类型/字段会
  panic

> 反射中修改变量的标准方式是：`ValueOf(&x).Elem()`，因为 Elem 返回的才是实际变量，可以进行 Set 操作

### `UnsafeAddr()`

返回底层地址，但不是 Go 可用安全指针（uintptr，不是 `unsafe.Pointer`）

```
(v Value) UnsafeAddr() uintptr
```

- 前提条件是必须 `CanAddr()==true`
- 非常危险，基本真实环境不会使用
- 主要用于低层序列化/FFI/性能调优，配合 unsafe 包

### `OverFlowInt()`

判断一个 int64 的值 x 是否能被表示为 `v.Type()` 对应的有符号整数类型

```
(v Value) OverflowInt(x int64) bool
```

- `Kind()` 必须为：Int/Int8/Int16/Int32/Int64，否则会触发 panic
- 如果 x 在对应类型的取值区间内返回 false（不溢出），如果超出对应类型最大/最小值返回 true（会溢出）
- `Int()` 读取不会溢出，OverflowInt + SetInt 写入才需要考虑溢出

> v 是 int8 类型的 Value：
>
> x = 127 -> false
>
> x = 128 或 x = -129 -> true
> v 是 int 类型：区间取决于编译器 int 宽度（一般 32 或 64）

### `OverflowUint()`

判断 uint64 值 x 能否表示为 `v.Type()` 对应的无符号整数类型

```
(v Value) OverflowUint(x uint64) bool
```

- `Kind()` 必须为：Uint/Uint8/Uint16/Uint32/Uint64/Uintptr，否则会触发 panic
- x 在目标类型可表示范围内返回 false，超出范围返回 true

> 参数是 uint64，没有负数，不看 v 当前的值，只看 v 的类型 能否表达 x

### `OverflowFloat()`

判断一个 float64 值 x 能否被表示为 v 的浮点类型

```
(v Value) OverflowFloat(x float64) bool
```

- `Kind()` 必须为 Float32/Float64，否则触发 panic
- 当读到了一个 float64，想塞进 float32 字段之前，先检查会不会溢出，再 SetFloat

### `OverflowComplex()`

判断一个 complex128 值 x 能否被 v 的复数类型表示

```
(v Value) OverflowComplex(x complex128) bool
```

- `Kind()` 必须为 Complex64/Complex128，否则触发 panic
- v 是 complex64：检查实部/虚部分别能否用于 float32 表示
- v 是 complex128：complex128 -> complex128 不会溢出
- 有时只存 complex64，但计算时用 complex128 做中间结果，这时写回前用 OverflowComplex 检查

### `Set()`

最通用的 Setter，用于将另一个 Value 替换当前值。

```
(v Value) Set(x Value)
```

- 要求必须 `v.CanSet()==true` 并且 `x.Type()` 必须 可赋值给 `v.Type()`

### `SetBool()`

为 boolean 值赋值

```
(v Value) SetBool(x bool)
```

- 直接 bool 不需 Value 包装
- `Kind != Bool` 或 `CanSet==false` 会触发 panic

### `SetInt()`

为 int 类型赋值

```
(v Value) SetInt(x int64)
```

- `v.Kind()` 必须为 Int/Int8/Int16/Int32/Int64
- 需要手动使用 OverflowInt 检查溢出情况

### `SetUint()`

为 uint 类型赋值

```
(v Value) SetUint(x uint64)
```

- `v.Kind()` 必须为 Uint/Uint8/Uint16/Uint32/Uint64/Uintptr
- 需要手动使用 OverflowUint 检查溢出情况

### `SetFloat()`

为 float 类型赋值

```
(v Value) SetFloat(x float64)
```

- `v.Kind()` 必须为 Float32/Float64
- 写入 float32 容易损失精度，但不会报错，需要使用 OverflowFloat 提前告警

### `SetComplex()`

为 complex 类型赋值

```
(v Value) SetComplex(x complex128)
```

- `v.Kind()` 必须为 Complex64/Complex128
- 需要手动使用 OverflowComplex 检查溢出情况

### `SetString()`

为 string 类型赋值

```
(v Value) SetString(x string)
```

- `v.Kind()` 必须为 string
- 无类型转换，不会把 `[]byte` 自动转 string，更不会接收 `rune/[]rune`

### `SetBytes()`

复制 x 的内容到底层 Byte Slice

Setter 中最特殊的一个

```
(v Value) SetBytes(x []byte)
```

- `v.Kind()` 必须为 slice且 elem 为uint8，即 []byte only
- 不会 realloc，受 Len/Cap 影响，底层共享内存，修改影响原对象
- `Set()` 方法不适用这种操作，必须用 `SetBytes`

### `SetPointer()`

直接设定底层指针指向

```
(v Value) SetPointer(x unsafe.Pointer)
```

- `v.Kind()` 必须为 Ptr/UnsafePointer
- 越过 Go 类型安全，只适合 low-level，必须搭配 unsafe，否则没有意义
- `UnsafePointer` 写入意味着绕过 GC 与类型系统，非常高阶

### `SetZero()`

把 v 的值设为类型零值

```
(v Value) SetZero()
```

- 支持几乎全部类型，等价于 `Set(Zero(v.Type()))`
- 主要用于快速清空字段或 slice元素，相比手动调用每类 SetXxx，SetZero 是最省脑子的初始化方式

### `Convert()`

类型强转，返回将值 v 转换为类型 t 的结果，如果 `v.CanConvert()` 为 false 或者将 v 转换为类型 t 会导致
panic，则 Convert 会引发 panic。

```
func (v Value) Convert(t Type) Value
```

## 容器相关方法/函数

1. 构造函数

    - `MakeSlice(typ Type, len, cap int) Value`
    - `SliceAt(typ Type, p unsafe.Pointer, n int) Value`
    - `MakeMap(typ Type) Value`
    - `MakeMapWithSize(typ Type, n int) Value`
    - `MakeChan(typ Type, buffer int) Value`

2. 容量长度类

    - `(v Value) Len() int`
    - `(v Value) Cap() int`

3. 索引/键访问类
    - `(v Value) Index(i int) Value`
    - `(v Value) MapIndex(key Value) Value`
    - `(v Value) MapKeys() []Value`
    - `func (v Value) MapRange() *MapIter`

4. 容器修改操作相关（SetZero 等操作在基础数据操作已经学习，这里同样适用）
    - `(v Value) SetLen(n int)`
    - `(v Value) SetCap(n int)`
    - `(v Value) Grow(n int)`
    - `(v Value) SetMapIndex(key, elem Value)`
    - `(v Value) Close()`

5. 容器追加/流式操作系列
    - `Append(s Value, x ...Value) Value `
    - `AppendSlice(s, t Value) Value`
    - `(v Value) Slice(i, j int) Value`
    - `(v Value) Slice3(i, j, k int) Value`
    - `(v Value) Send(x Value)`
    - `(v Value) TrySend(x Value) bool`
    - `(v Value) Recv() (Value, bool)`
    - `(v Value) TryRecv() (Value, bool)`
    - `Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)`
    - `(v Value) SetIterKey(iter *MapIter)`
    - `(v Value) SetIterValue(iter *MapIter)`
    - `(v Value) Seq()`
    - `(v Value) Seq2()`

### `reflect.MakeSlice`

reflect 中最常用的 Slice 构造函数，用于动态创建一个 Slice，类似于 `make([]T, len, cap)`，可用于动态创建 T 类型切片，长度可变，可修改元素

```
func MakeSlice(typ Type, len, cap int) Value
```

### `reflect.SliceAt`

极度底层、高危、和 unsafe 深度绑定的构造方式

```
reflect.SliceAt(typ Type, p unsafe.Pointer, n int) Value
```

1. typ 同样必须是 slice 类型

    - 不是 slice 就 panic
    - 一般也是 `reflect.SliceOf(elemType)` 来的

2. 不会分配内存
    - 和 NewAt 类似：它只是用 p 这块内存伪装成一个 slice
    - p 必须指向一段已经存在的、连续的、适配的内存区域

3. 长度 / 容量问题
    - n 是 slice 的长度（len）
    - 容量（cap）通常也是 n（具体以文档为准，一般是 n）
    - 相当于说："在地址 p 开头，有 n 个元素，把这看成一个 slice 给我"

4. p 的要求非常严格：

    - p 这块内存要能容纳 n 个元素：`n * ElemType.Size()` 的字节数必须都合法可访问
    - 对齐要正确（元素类型的 alignment）
    - 这块内存在使用期间不能被释放/移动（比如来自 C malloc、全局内存、unsafe.SliceData 等）

5. 危险点：
    - 如果 p 指的是一块本不属于这个类型的内存会导致 UB（未定义行为）
    - 如果 n 写大了，读写越界 ⇒ 崩溃 / 数据损坏
    - 如果底层内存已经被 GC 回收或移动 ⇒ 悬空指针

### `reflect.MakeMap`

动态创建 map，相当于 `make(map[K]V)`，`typ.Kind()` 必须为 Map，返回的值是一个空 map

```
MakeMap(typ Type) Value
```

### `reflect.MakeMapWithSize()`

与 MakeMap 相同，但是可以提供初始容量，相当于 `make(map[K]V, size)`，适用于批量写入大 map，减少扩容次数

```
MakeMapWithSize(typ Type, n int) Value
```

- 同样要求 typ.Kind() == Map，否则 panic
- 返回一个空 map Value，带有预估大小 n，给 runtime 一个初始空间大小的 hint
- 跟 slice 不同，map 的 n 是预估 key 数量，不是硬性的容量限制

### `reflect.MakeChan()`

动态创建 chan，相当于 `make(chan T, buf)` `typ.Kind()` 必须为 Chan，否则会 panic

```
MakeChan(typ Type, buffer int) Value
```

- buffer == 0 -> 无缓冲 channel （同步）
- buffer > 0 -> 有缓冲 channel
- buffer < 0 -> panic

### `Len()`

获取容器长度，`v.Kind()` 必须为：Slice / Array / String / Map / Chan，否则会触发 panic

```
(v Value) Len() int
```

| Kind   | Len 返回值                  |
|--------|--------------------------| 
| Slice  | 当前长度，是可见长度，不一定等于 cap     | 
| Array  | 数组长度（固定）                 | 
| String | 字符数（按 byte 而不是 rune 数）   | 
| Map    | 键数量，会随着 SetMapIndex/删除变化 | 
| Chan   | 缓冲区当前已有的元素数量             | 

### `Cap()`

返回容器容量，`v.Kind()` 必须为：Slice / Array / Chan，否则会触发 panic

```
(v Value) Cap() int
```

- 对 Slice：cap 是底层数组的容量，可以用让 Grow / Append 判断需不需要拓容
- 对 Array：cap == len（编译时常量确定）
- 对 Chan：cap 是缓冲区大小（channel capacity）

### `Index()`

根据下标返回指定元素，`v.Kind()` 必须为：Slice / Array / String，否则会触发 panic

```
(v Value) Index(i int) Value
```

- 对 Slice：`v.Index(i)` 的可写性在于 slice 本体是否可寻址，如果 slice 是 `ValueOf(&s).Elem()` 是可写的，如果是
  `ValueOf(s)`（拷贝），则不可写
- 对 Array：若 array 是拷贝来的ValueOf(arr) -> Index 返回值 不可寻址不可写，如果 array 是 `ValueOf(&arr).Elem()` 可寻址可写
- 对 String：返回的是 uint8，而不是字符，因为字符串是不可变的，所以结果永远不可 Set，想处理 unicode 需要先转换为 `[]rune`

### `MapIndex()`

返回 map 映射 `v` 中与键 `key` 关联的值，只适用于 Map，否则直接 panic

```
(v Value) MapIndex(key Value) Value
```

- 如果 key 存在：返回 `map[key]` 对应的值的 Value（只读副本，不可寻址，不可 Set）
- 如果 key 不存在：返回 zero Value（invalid Value，注意不是零值，所以在使用前最好用 IsValid 进行检查）

### `MapKeys()`

返回一个包含 Map 中所有键的切片，顺序未指定，只适用于映射 Map，否则导致 panic

```
(v Value) MapKeys() []Value
```

- 返回 key 的 value 列表，顺序不保证稳定（和 Go 原生 map 一样无序）
- 返回值是 `[]Value`，slice 本身可以寻址，但是元素 Value 不可 Set，也就是对应 key 的可读副本（map key 自身没意义）
- 不是共享 map 内部结构，每个 key 都是新建的 `reflect.Value`

> 主要用于遍历 map，但是在大的 map 上使用可能有性能问题，因为每次查都会重新 hash key

### `MapRange()`

Go1.12 加入的新 API，用来安全且高效地遍历 map，同样只适用于 map

```
(v Value) MapRange() *MapIter
```

- 返回的 `*MapIter` 是个迭代器：
    - `.Next()` -> bool：是否有下一个元素
    - `.Key()` -> Value：当前 key
    - `.Value()` -> Value：当前 value

- 同样顺序无保证
- 返回的 key/value 都不可 Set，要修改 map 元素必须用 `SetMapIndex(k, newVal)`

> 官方推荐遍历 map 使用 MapRange

### `Slice()`

返回 `v[i:j]`，`v.Kind()` 必须为：Slice / String，否则会触发 panic，如果 i/j 越界或 i>j 也会导致 panic

```
(v Value) Slice(i, j int) Value
```

- 对 slice：共享底层数组的新 Slice 切片
- 对 string：返回新 string，但其内部也是共享原 memory，只是不可修改

### `Slice3()`

`v[i:j:k]` 三参数切片，只能用于 slice 切片操作

```
(v Value) Slice3(i, j, k int)
```

- 允许指定 cap 为 k - i
- 常用于控制不能访问原 slice 未被暴露的后半部分（安全 slice）
- `i <= j <= k <= cap(v)` 必须成立，同时 `v.Kind()` 必须为 slice

### `SetLen()`

将 v 的长度设置为 n。相当于 `s = s[:n]`

如果 `v.Kind()` 不是 Slice，或 n 为负数或大于切片的容量，或者 `Value.CanSet` 返回 false，则会引发 panic

```
func (v Value) SetLen(n int)
```

### `SetCap`

将 v 的容量设置为 n（十分危险的操作，因为 Go 本身没有语法允许修改 cap）

如果 `v.Kind()` 不是 Slice，或者 n 小于切片的长度或大于切片的容量，或者 `Value.CanSet` 返回 false，则会引发 panic

```
func (v Value) SetCap(n int)
```

### `Grow()`

扩容 slice 的容量，使它至少增加 n 个容量（推荐操作）

```
func (v Value) Grow(n int)
```

- `v.Kind()` 必须为 Slice
- 可能重新分配底层数组（类似 append 扩容）
- Grow 不改变 Len，只改变 Cap

> 用于构建高效 decoder，可以一次扩容多个空间，需要提前保证 slice 完成多次 append 仍不需重新分配

### `SetMapIndex()`

将 map 映射中与键 key 关联的元素赋值为 elem

```
func (v Value) SetMapIndex(key, elem Value)
```

- 如果 `v` 的 `Kind` 不是`Map` 类型，则该函数会引发 panic
- elem 非空，相当于 `m[key] = elem`
- elem 为 invalid Value：相当于删除 key（delete(m,key)）

### `Clear()`

容器通用方法，用于清空容器

```
func (v Value) Clear()
```

Clear 对不同容器有不同效果：

| Kind  | Clear 行为               |
|-------|------------------------|
| Slice | 将每个元素设为零值（不改变 len/cap） |
| Array | 同上，清零所有元素              |
| Map   | 清空所有键值对（等同 delete all） |
| 其他    | panic                  |

### `Close()`

用于关闭通道 v。如果 v 的 Kind 不是 Chan 类型，或者 v 是仅接收通道，则会触发 panic

```
func (v Value) Close()
```

### `reflect.Append`

`append` 方法将值 x 追加到切片 s 中，并返回一个新的 Slice Value 结果切片

```
func Append(s Value, x ...Value) Value
```

- `s.Kind()` 必须为 Slice 且不是不可用的 Value，否则直接 panic
- 完全等价于 Go 内置 append：`s2:=append(s, x...)`

### `reflect.AppendSlice`

用于拼接两个 slice，等价于 `append(s, t...)`

```
func AppendSlice(s, t Value) Value
```

> `s.Kind()` 和 `t.Kind()` 必须都为 Slice 且 `t.Type().Elem()` 必须与 `s.Type().Elem()` 匹配

### `Send()`

将一个值 x 发送到通道 v，等价于 `v <- x`，必须 `v.Kind() == Chan`，否则 panic

```
func (v Value) Send(x Value)
```

- 若 channel 无缓存则会阻塞直到有人接收，若 channel 已关闭会导致 panic
- `x.Type()` 要与 channel 的元素类型一致，reflect 不做自动 Convert

### `TrySend()`

`Send` 的非阻塞操作方法，

```
func TrySend(x Value) bool
```

- 返回 true 表示 send 成功，返回 false 表示失败，可能通道塞满或为准备好
- 和 Send 一样：类型不匹配、chan 不对、已关闭会导致 panic
- 是构建非阻塞并发逻辑的关键

### `Recv()`

用于从管道接收数据，等价于 `v, ok := <-ch`，`v.Kind()` 必须为 chan

```
func (v Value) Recv() (Value, bool)
```

- 返回实际接收到的值和是否接收到
- 阻塞式 receive，若 chan 已关闭且为空返回零值和 false

### `TryRecv()`

非阻塞 recv 接收方法

```
func (v Value) TryRecv() (Value, bool)
```

等价于：

```
select {
case v, ok := <-ch:
    return (v, ok)
default:
    return (zeroValue, false)
}
```

- 第二个值 ok 为 true 表示实际收到值或收到关闭状态，为 false 表示没值可读（未准备好）
- `TryRecv` 和 `TrySend` 一样用来构建 polling 或 select-less 非阻塞逻辑

### `reflect.Select`

reflect 提供的 反射版 select 操作，强大但使用频率不高

类似原生 `select { case ... }` 语句，用于 多个 channel 的选择，会阻塞，直到至少有一个 case 可以执行。
返回所选 case 的索引，如果该 case 是接收操作，则返回接收到的值以及一个布尔值，该布尔值指示该值是否对应于通道上的发送操作（而不是由于通道已关闭而接收到的零值）。
`SELECT` 语句最多支持 65536 个 case。

```
func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)
```

- 多路 wait / send / recv 随机选择就绪的 case
- 可设置 default 分支
- 可进行 recv / send / default

返回值：

- chosen int：被选中的 case 下标
- recv Value：当该 case 是 Recv 时返回的值
- recvOK bool：类似 <-ch 的 ok

> case.Dir 不合法 / 类型不匹配 / chan 不是 Channel 类型 / chan 已关闭而做 Send 会导致 panic

### `SetIterKey()`

把 v 设为 iter 当前的 Key，等价于 `v.Set(iter.Key())`，但避免了分配新的 `Value` 对象，v 通常是 mapKeyType 的 Value

如果 `Value.CanSet` 返回 false，则会引发 panic，`iter.Key().Type()` 必须与 `v.Type()` 相匹配

```
func (v Value) SetIterKey(iter *MapIter)
```

这是为了支持 `for key := range map` 这种单变量遍历加入的新能力。

### `SetIterValue()`

把 v 设为 iter 当前的 Value，等价于 `v.Set(iter.Value())`,但避免了分配新的 `Value` 对象

`Value.CanSet` 必须为 true 且 类型匹配 `map.value` 类型

```
func (v Value) SetIterValue(iter *MapIter)
```

> SetIterKey / SetIterValue 是支持反射 API 实现 Go 的for range map的赋值语义，真实使用基本用不到，主要了解即可

### `Seq()`

Go 1.23 新增方法，返回一个 `iter.Seq[Value]`单值序列，该单值序列会遍历 v 的元素，
等价于一个 `yield-like` 生成器，可以用 Go 的新 iterator 模式遍历 map。

```
func (v Value) Seq() iter.Seq[Value]
```

- `v.Kind()` 必须为 map，否则导致 panic
- 每次迭代返回一个 `reflect.Value` 表示 key，顺序不保证（和 map range 一致）

### `Seq2()`

Go 1.23 新增方法，返回一个 `iter.Seq2[Value, Value]`，该 `iter.Seq2[Value, Value]` 会遍历 v 的元素，
这是 reflect 对 双变量迭代（k,v） 的官方支持。

```
func (v Value) Seq2() iter.Seq2[Value, Value]
```

每次迭代:

- keyVal ：`reflect.Value`
- valVal ：`reflect.Value`
- 都不可 Set（map 元素是副本）

> `Seq` 和 `Seq2` 主要配合 iter 包进行使用，这是 Go 新语法提供的 iterator ecosystem

## Struct 结构体相关

`TypeKind/CanSet/...` 这些通用的这里不再讲解了

1. 字段相关
    - `(v Value) NumField() int`
    - `(v Value) Field(i int) Value`
    - `(v Value) FieldByName(name string) Value`
    - `(v Value) FieldByNameFunc(match func(string) bool) Value`
    - `(v Value) FieldByIndex(index []int) Value`
    - `(v Value) FieldByIndexErr(index []int) (Value, error)`

2. 方法相关

    - `(v Value) NumMethod() int`
    - `func (v Value) Method(i int) Value`
    - `func (v Value) MethodByName(name string) Value`

### `NumField()`

返回回这个 struct 类型有多少个直接字段，只适用于 `v.Kind()==Struct`，否则导致 panic，字段包括：

- 普通字段
- 匿名字段（embedded field）也算一个 field
- 导出+未导出一起计算

```
func (v Value) NumField() int
```

常用于做遍历所有字段时，先 `NumField()` 再 `for i := 0; i < NumField(); i++ { Field(i) }`

> 它只是告诉你这个 Type 的字段数，不管能不能访问/能不能 Interface

### `Field()`

按 index 取字段 Value，是最基础的字段访问

```
func (v Value) Field(i int) Value
```

- 只适用于 Struct；其他 Kind 会 panic
- index 必须是 `[0, NumField())` 范围内，否则 panic
- 引用的是结构体里的真实字段，不是值拷贝，所以如果 struct 本身是可寻址的（ValueOf(&s).Elem()），这个字段通常也可寻址
- `CanSet()` 决定能不能写字段：
    - `struct Value` 来自指针 `.Elem()` 且字段是导出（大写），通常 `CanSet()` 为 true
    - 如果字段未导出，即便可寻址，`CanSet()` 仍然为 false
- `CanInterface()` 决定能不能 `.Interface()` 出来：
    - 导出字段：true
    - 未导出字段：false，调用 `Interface()` 会 panic
- 只看字段名，不看 tag，tag 还是要通过 `reflect.Type` 去看

### `FieldByName()`

按字段名查找字段（只查当前这一层 struct 的字段，不会自动深度递归嵌套 struct）

```
func (v Value) FieldByName(name string) Value
```

- 找到字段就返回字段的 Value（规则和 `Field(i)` 一样）
- 没有找到字段就返回 `invalid Value`（`IsValid()==false`）

### `FieldByNameFunc()`

用一个匹配函数来自定义字段名查找规则

```
func (v Value) FieldByNameFunc(match func(string) bool) Value
```

- 会对这个 struct 的所有可见字段名调用 match(name)
- 第一个返回 true 的字段返回对应字段 Value
- 没有匹配到返回 invalid Value
- 若匹配多个字段会导致panic（因为不唯一）

**使用场景**

主要用于实现各种模糊匹配策略，比如：

- 忽略大小写：strings.EqualFold(name, "xxx")
- 支持多个候选名（如 "ID", "Id", "id"）
- 自定义规则（前缀 / 后缀）

### `FieldByIndex()`

返回与索引对应的嵌套字段，配合 `reflect.Type` 的 `StructField.Index` 专用的，

```
func (v Value) FieldByIndex(index []int) Value
```

在 `reflect.Type` 里，每个 [StructField](StructStructField.md) 都有一个 `Index []int` 字段，表示访问这个字段时的嵌套路径，FieldByIndex
就是专门给这个 Index 用的：

```
sf := someType.Field(i)     // 反射 Type 得到 StructField
idx := sf.Index             // []int
fv  := v.FieldByIndex(idx)  // 直接拿到对应 value
```

### `FieldByIndexErr()`

这是 FieldByIndex 的不 panic 版本

```
func (v Value) FieldByIndexErr(index []int) (Value, error)
```

行为同 FieldByIndex，但：

- 如果 index 无效、不匹配、访问出错，返回 error 而不是 panic
- 正常时 error 为 nil

### `NumMethod()`

返回可导出方法（exported methods）的数量

```
func (v Value) NumMethod() int
```

- 只统计导出方法（Name大写），结构体中有小写方法不会统计在内
- 遵循 method set 方法集规则：值方法（`T receiver`）属于 T 与 `*T` 都可见，指针方法（`*T receiver`）只有在 `*T` 上可见

### `Method()`

按索引返回绑定方法（bound method）

```
func (v Value) Method(i int) Value
```

- 返回一个 `reflect.Value`（Kind = Func），这个函数已经绑定了接收者（receiver），是可以 call 调用的
- `i < 0` 或 `i >= NumMethod()` 会导致 panic
- v 不是 `struct / ptr-to-struct` 时也可用，但如果没有方法会 panic ，例如 `v.Kind=Int;NumMethod=0;Method(0)` 会 panic
- reflect 保证方法按照字典序排序（方法名排序），这样一个类型在不同机器上也是顺序一致的

### `MethodByName()`

按照方法名返回对应的方法

```
func (v Value) MethodByName(name string) Value
```

- 若找到导出方法返回绑定方法（Value）
- 若找不到返回 `invalid Value`（`IsValid()==false`,`Kind()==Invalid`），调用 `.Call` 会导致 panic

## 方法/函数相关

- `(v Value) Call(in []Value) []Value`
- `func (v Value) CallSlice(in []Value) []Value`
- `MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value`

### `Call()`

`v.Call([]Value{arg0, arg1, ...})`等价于`v(arg0, arg1, ...)`这个 API 可以调用：

- function（`Value.Kind()==Func`）
- method（通过 `Method(i)` 或 `MethodByName` 返回）
- interface 内的函数值

```
func (v Value) Call(in []Value) []Value
```

- 参数数量必须和函数签名匹配（除非是 variadic 可变参数）
- 参数类型必须与函数签名的参数类型可赋值（AssignableTo）

**Call 的 variadic（可变参数）规则**

函数若定义为：`func F(a int, bs ...string)`则你可以传：

```
Call([]Value{1, "x", "y", "z"})
```

反射会把最后 n 个参数自动当“展开的 variadic 参数”，变成：

```
F(1, []string{"x","y","z"}...)
```

> `.Call` 不能传 slice 本身作为 variadic 参数，这正是 `.CallSlice` 存在的原因

### `CallSlice()`

把 slice 本身作为 variadic 参数整体传入，用于：`F(1, []string{"x","y"}...)`

```
func (v Value) CallSlice(in []Value) []Value
```

基本规则：

- 被调用的函数必须是 variadic（可变参数）函数，否则 panic。
- in 数组的长度必须和函数参数数量一致 （variadic 部分仍然算 1 个）
- `in[last]` 必须是一个 slice 的 Value
- 这个 slice 的元素类型必须和 variadic 元素类型一致


### 5. `reflect.MakeFunc`

比较特殊，可以动态生成一个函数值（闭包），并可以自定义函数体。会返回一个类型为 typ 的函数对象，fn 会作为函数逻辑替代执行

```
func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value
```

