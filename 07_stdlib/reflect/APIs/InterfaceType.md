# `reflect.Type`

`reflect.Type` 是反射中最基础的概念，是 Go 编译器生成的类型描述信息（Metadata），反射通过它知道一个值的类型结构是什么样的。

在 Go 源码中，Type 是一个接口，遮住了内部复杂的实现（不同类型对应不同底层 `struct`，如 `rtype`、`arrayType`、`structType` 等）。

## Type 结构

`reflect.Type` 作为一个接口，里面定义的主要都是类型元信息访问的一些方法，在最后会对这些方法进行分类简单介绍。

```
type Type interface {
    ...
}
```

需要知道 `reflect.Type` 只是一个接口，真正持有信息的是底层的 `rtype` 结构，所以下面 `TypeOf` 等方法返回的基本都是以
`*rtype` 作为或它的派生作为实现。

## 主要入口函数

前两个是用的比较多的顶层函数用于获取一个值或type的类型，后面几个主要用于动态创建一个新的 `reflect.Type`

### 1. `reflect.TypeOf`

这是反射中获取类型最常用的入口函数，可以根据一个值返回它的类型，是反射最基础的入口。
反射类型体系所有 `Type / Value API` 基本全部都要靠 `reflect.TypeOf()` 做入口。

```
func TypeOf(i any) Type
```

- 需要传一个值 value 而不是类型，这个值可以是任何类型（`int/struct/slice/pointer/...`）
- 如果传入 nil interface 则返回 nil
- 返回这个值动态类型（dynamic type），类型是 `reflect.Type`（底层是实现了 Type 的 `*rtype`）
- 指针和值类型 Type 不一样，`reflect.TypeOf(User{}).Kind()=Struct`， 而 `reflect.TypeOf(&User{}).Kind()=Pointer`

**常用于**

- 获取类型信息（配合 Kind/Name/Field）
- 检查接口实现或赋值关系（与 Implements / AssignableTo）
- 容器拆解（配合 Elem/Key/Len）
- 反射方法来源（配合 NumMethod/MethodByName）

### 2. `reflect.TypeFor`

Go 1.22+ 新增的类型构造函数（现代泛型友好），根据 泛型 推导类型，而不是值。

```
func TypeFor[T any]() Type
```

- TypeFor 返回表示类型参数 T 的 类型
- 只需要一个类型不需要实例就看也获取 Type
- 常用于反射泛型，构造类型，一定要类型而非值

以往获得类型必须 `TypeOf(value)`：

```
reflect.TypeOf(123)
reflect.TypeOf(User{})
```

但如果想直接访问类型本身（非值），以前必须构造临时变量：

```
reflect.TypeOf((*User)(nil)).Elem()
reflect.TypeOf(User{})      // 但这是值，不是裸类型语义
```

Go1.22后可以直接使用 `reflect.TypeFor[User]()`，重点是不需要值也不需要实例，直接通过一个类型就看也获得 Type

```
reflect.TypeFor[User]()     // 直接得到 User 的 reflect.Type
reflect.TypeFor[*User]()    // 指针类型也OK
reflect.TypeFor[[]User]()   // 任意复杂类型可直接获取
```

TypeOf 需要一个值，所以在很多场景下非常不方便，例如：

- 只知道类型，不想创建实例
- 在写泛型库，而 TypeOf 做不到泛型的类型推导

这时就建议直接使用 TypeFor。

### 3. `reflect.ArrayOf`

动态创建一个固定长度数组类型，例如想创建 `[5]int` 这种数组类型，但数量是运行时才能确定：

```
func ArrayOf(length int, elem Type) Type
```

- length：数组长度（固定大小）
- elem：数组元素的类型（reflect.Type）
- 返回一个新创建的 数组类型，具体类型形状类似 `[length]elemType`
- 返回值类型依然是 reflect.Type

可以用它创建：`[3]string`，`[10]float64`，`[N]User` 等

### 4. `reflect.SliceOf`

动态创建一个 slice 类型

```
func SliceOf(elem Type) Type
```

- elem：切片中元素的类型（reflect.Type）
- 返回一个切片类型，类型形状类似 `[]elemType`，比如 elem 表示 int 类型，则 `SliceOf(elem)` 表示 `[]int`

可以创建：`[]User`，`[]map[string]int`，`[][]byte` 等

### 5. `reflect.PointerTo`

动态创建一个指针类型，如：`*int`, `*User`, `*[]string`，PointerTo 返回包含元素 t 的指针类型。例如，如果 base 表示类型 Foo，则
`PointerTo(base)` 表示 `*Foo`

```
func PointerTo(base Type) Type
```

- base: 基础类型（reflect.Type）
- 返回指向该类型的 指针类型，类型形状类似 `*baseType`

### 6. `reflect.MapOf`

动态创建一个 map 类型，如：`map[string]User`，`map[int][]byte`，`map[string]map[int]`。

string返回具有给定键和元素类型的映射类型。例如，如果 k 表示整数，e 表示字符串，则 `MapOf(k, e)` 表示 `map[int]string`。
如果键类型不是有效的映射键类型（即，如果它没有实现 Go 的 `==` 运算符），则 MapOf 会引发 panic。

```
func MapOf(key, elem Type) Type
```

- key：key 的类型（reflect.Type）
- elem：value 的类型（reflect.Type）
- 返回一个 map 类型，形状类似 `map[keyType]valueType`

### 7. `reflect.ChanOf`

动态创建一个 channel 类型，返回具有给定方向和元素类型的通道类型。

```
func ChanOf(dir ChanDir, elem Type) Type
```

- dir：channel 的方向（reflect.ChanDir）
    - SendDir：只能发送
    - RecvDir：只能接收
    - BothDir：双向（普通 chan）
- elem：通道中元素的类型（reflect.Type）
- 返回一个新的 channel 类型，类型形状取决于方向：`chan T`，`<-chan T`（只读），`chan<- T`（只写）

### 8. `reflect.FuncOf`

动态创建一个函数类型，这个非常强大，可以构造任意函数签名，包括：高阶函数，返回多个值的函数，带可变参数的函数

```
func FuncOf(in, out []Type, variadic bool) Type
```

- `in []Type`：输入参数类型列表，顺序严格代表函数参数列表
- `out []Type`：输出参数类型列表，对应函数返回值顺序
- `variadic bool`：是否为可变参数函数（最后一个 in 类型视作 `...T`）
- 返回一个具有指定参数签名的新函数类型

主要用于动态构建函数签名，用于 RPC、调用代理、反射代理函数等

### 9. `reflect.StructOf`

```
reflect.StructOf(fields []StructField) Type
```

- 一个 []reflect.StructField 切片
- 每个 StructField 包含：
    - Name：字段名
    - Type：字段类型
    - Tag：tag 字符串（如 json:"xxx"）
    - Anonymous：是否匿名字段
    - Offset 等底层字段（通常不需要手动设置）
- 返回一个在运行时构造出来的新结构体类型，结构体字段的顺序按照切片顺序

主要用于动态构造 struct 类型，动态 schema、ORM、序列化框架常用

# Type 元信息方法

Type Metadata Accessors，也就是 Type 接口中定义的那些类型元信息访问方法/类型反射接口方法，主要分为下面几个级别：

## 类型基础信息相关方法

这类方法主要用于判断类型身份、类别、大小、包名、打印表示等。

### 1. Kind

Kind 主要用于判断类型类别（非常重要，是判断类型行为的入口）

```
Kind() Kind
``` 

返回类型的底层分类，参考 [TypeKind.md](TypeKind.md)，开发中基本都需要根据 Kind 进行分类然后再决定下一步怎么解析。

### 2. Name

返回类型在当前包内的名字

```
Name() string
```

可以通过 ` t.Name()` 是否为 `""`来判断是否为自定义类型

- 对于已定义的类型，Name 返回该类型在其包内的名称
- 对于其他（未定义的）类型，它返回空字符串

| 类型                       | Name()         |
|--------------------------|----------------|
| `type User struct{}`     | `"User"`       |
| `type MyInt int`         | `"MyInt"`      |
| `*User`                  | `""`（指针不是定义类型） |
| `[]int` / `map` / `chan` | `""`（无定义名字类型）  |

### 3. PkgPath

PkgPath 返回自定义类型的包路径，即导入路径，也就是这个类型属于哪个包，如果不是自定义类型返回 `""`

```
PkgPath() string
```

主要用于：

- 自动加载模块
- 动态路由
- 反射生成文档
- 类型白名单规则

| 类型                             | PkgPath()         |
|--------------------------------|-------------------|
| struct User in `model/user.go` | `"project/model"` |
| fmt.Stringer                   | `"fmt"`           |
| int/string/内建类型                | `""`（系统类型）        |
| 指针类型                           | `""`（不是定义类型）      |

### 4. String

返回类型的字符串表示，字符串表示形式可能使用缩写的包名，比如 `main.User`

```
String() string
```

`String()` 不保证在所有类型中唯一，比较类型是否相同一定要用：`t1==t2`

### 5. Size

返回类型数据所占内存字节数，类似于 `unsafe.Sizeof`

```
Size() uintptr
```

底层性能优化、有序排列 struct 字段防 padding 时会用到

- int32 → 4
- int64 → 8
- string → 16 或更多   (变量结构大小)
- struct{A int;B int} → 16

### 6. Align / FieldAlign

Align 返回此类型值在内存中分配时的字节对齐方式。FieldAlign 返回此类型值在结构体中用作字段时的字节对齐方式。

```
Align() int
FieldAlign() int
```

`FieldAlign() >= Align()` ，字段如果不能被对齐，Go编译器会自动 padding 填充。

## Struct 相关方法

这些方法主要用于从 Type 获取 StructField 相关信息。

下面这些在 Value 中也有相关操作，但是完全不一样，比如 `Type.Field(i)` 与 `Value.Field(i)` 区别：

- `t.Field()` -> 字段信息 (StructField)
- `v.Field()`-> 字段真实值 (Value)

### 1. NumField

NumField 方法返回结构体类型的字段数量，如果 `t.Kind()` 不是 struct 结构体会引发 panic。

```
NumField() int
```

主要用于遍历字段，需要注意每个匿名字段（embedded）也算一个，而匿名字段内部字段不自动展开，需要 Index 或 VisibleFields。

### 2. Field

Filed 根据下标返回结构体类型的第 i 个字段。如果`t.Kind()` 不是 struct会引发 panic。
如果 i 不在 `[0, NumField())` 范围内，页会引发 panic。

```
Field(i int) StructField
```

### 3. FieldByName

按字段名进行查找对应的字段信息，大小写敏感、仅匹配导出字段

```
FieldByName(name string) (StructField, bool)
```

返回具有给定名称的结构体字段以及一个布尔值，指示是否找到该字段。
如果返回的字段是从嵌入式结构体提升而来，则返回的 StructField 中的 Offset 是嵌入式结构体中的偏移量。

### 4. FieldByNameFunc

可以自定义规则查字段名，而不是精确匹配。

```
FieldByNameFunc(match func(string) bool) (StructField, bool)
```

- 只取最浅层字段
- 同深度多个字段匹配会导致冲突取消，返回 false
- 完全遵守 Go 原生匿名字段查找规则

### 5. FieldByIndex

FieldByIndex 主要用于匿名字段和嵌套字段，返回与索引序列对应的嵌套字段。

```
FieldByIndex(index []int) StructField
```

- 等价于对每个索引 i 依次调用 Field 方法
- 如果 `t.Kind()` 不是 Struct，同样会引发 panic

## Method 相关方法

Type 的 Method 相关方法主要用于 动态 RPC / HTTP 路由 / IoC 注入 / 框架级调用。
需要注意 Method 读取信息的入口在 Type ，而Method 调用行为的入口在 Value 二者配合才是完整的反射调用方式。

Type 的 Method 相关方法可见性

| Category         | Method / MethodByName / NumMethod 可见性规则 |
|------------------|-----------------------------------------|
| *struct / struct | 只能访问 导出方法（大写开头）                         |
| interface        | 可访问 所有方法（导出 + 非导出）                      |

### 1. NumMethod

返回可通过 Method 访问的方法数量。对于非接口类型，它返回导出的方法数量。对于接口类型，它返回导出和未导出方法的数量。

```
NumMethod() int
```

- 结构体的私有方法不可反射访问
- interface 方法本来就是签名描述，不涉及值可见性

### 2. Method

通过传入下标返回类型方法集中的第 i 个方法，返回类型为 `reflect.Method`，参考 [StructMethod.md](StructMethod.md)

```
Method(int) Method
```

注意这里也只是拿到方法信息，包括参数返回值等，但是并不是可执行对象，调用还是要走 `Value.Method(i)`

### 3. MethodByName

按名字查方法获取方法信息，如果没有找到会返回 false。

```
MethodByName(string) (Method, bool)
```

### 容器类型相关方法

主要用于 JSON / ORM / 反序列化 / map->struct绑定。

### 1. Elem

获取容器内元素类型（最为重要），如果 `t.Kind()` 不是 `array / channel / map / pointer / slice` 会触发 panic。

```
Elem() Type
```

适用类型：

| Kind    | Elem含义  |
|---------|---------|
| Pointer | 指向元素类型  |
| Slice   | 元素类型    |
| Array   | 元素类型    |
| Map     | value类型 |
| Chan    | 传输数据类型  |

> `ORM/JSON` 解码就是靠 Elem 一层层剥开类型的

### 2. Len

返回 array 数组类型的长度。如果 `t.Kind()` 不是 Array，则会引发 panic。

```
Len() int
```

> Slice 没有 Len()（只有运行时长度可通过 `Value.Len` 获取）

### 3. Key

返回映射类型的键类型。如果 `t.Kind()` 不是 Map，则会引发 panic。

```
Key() Type
```

### 4. ChanDir

返回通道类型的方向，返回值类型 ChanDir 参考 [TypeChanDir.md](TypeChanDir.md)。如果 `t.Kind()` 不是 Chan，则会引发 panic。

```
ChanDir() ChanDir
```

## 类型关系相关方法

写 DI 容器 / 插件系统 / 通用工具库 / 反射工厂 时非常关键的一组 API。

### 1. Implements

用于判断类型是否实现了接口类型 u，也就是 t 是否实现接口类型 u。

```
Implements(u Type) bool
```

- u 必须是 接口类型，否则行为没意义
- t 可以是任意具体类型（struct、ptr、别名等）
- 本质就是编译器那套方法集合规则在运行时的版本

### 2. AssignableTo

用于判断一个值是否可以赋给类型 u。如果有一个变量 x 类型是 t，能不能写 `var y u; y=x`。

```
AssignableTo(u Type) bool
```

某些情况比如自定义 MyInt 类型和 int 构造出的值虽然底层都是 int kind，但是类型不同，不能直接赋值（包含接口情况 +
同类型/别名兼容情况等）。

### 3. ConvertibleTo

是否允许显式类型转换，能否写 `var x t; y := u(x)` 这种转换（不一定安全，会可能 panic）

```
func (t Type) ConvertibleTo(u Type) bool
```

> 就算 ConvertibleTo 返回 true，转换时还是可能 panic（主要是一些特殊 case，如切片长度不够转换为数组指针等，注释里其实有写）

### 4. Comparable

用于判断这个类型的值能不能合法地用 `==` 比较

```
func (t Type) Comparable() bool
```

- `slice / map / func` -> 不可比较
- `array / struct` -> 看内部所有字段/元素是否可比较
- interface -> 值可比较，但实际比较时可能因为内部动态类型不可比较而 panic

**使用场景**

- 深度比较前先判断能否比较
- 做 key 时限制类型
- 实现自己的 `DeepEqual / set` 类型时过滤类型

## Function 相关方法

这几个方法只对 `t.Kind() == Func` 的 Type 有意义，用来看清楚一个函数长什么样。

### 1. NumIn / In

NumIn 返回函数类型的输入参数数量，如果 `t.Kind()` 不是 Func 会引发 panic。
In 返回函数类型的第 i 个输入参数的类型。如果 `t.Kind()` 不是 Func 或 i 不在 `[0, NumIn())` 范围内会引发 panic。

```
In(i int) Type
NumIn() int
```

> 可变参数(...) 也算 1 个 input

### 2. IsVariadic

用于判断是否是为可变参数函数。

```
IsVariadic() bool
```

- 可变参数函数的最后一个参数类型，在 Type 里表现为 `[]T`
- 想判断是不是真的是可变参数，要看 `IsVariadic()`

### 3. NumOut / Out

NumOut 返回函数类型的输出参数数量，如果 `t.Kind()` 不是 Func 会引发 panic。
Out 返回函数类型的第 i 个输出参数的类型，如果 `t.Kind()` 不是 Func 或 i 不在 `[0, NumOut())` 范围内会引发 panic。

```
Out(i int) Type
NumOut() int
```

## Overflow 类型溢出检测相关方法

| 方法                                   | 适用类型                   | 入参         | 作用         |
|--------------------------------------|------------------------|------------|------------|
| `OverflowInt(x int64) bool`          | int / intX             | int64      | x 是否能装进该类型 |
| `OverflowUint(x uint64) bool`        | uint / uintX           | uint64     | 同理但无符号     |
| `OverflowFloat(x float64) bool`      | float32 / float64      | float64    | 能否表示该浮点数   |
| `OverflowComplex(x complex128) bool` | complex64 / complex128 | complex128 | 同理复数       |

这些都要求 `t.Kind` 必须与对应类型相匹配，否则直接 panic。

主要用于：

1. 用于反射赋值前验证是否会溢出
2. 用于实现安全的 `Convert()`
3. 用于自制序列化/反序列化框架

## CanSeq / CanSeq2

这两个是 Go1.22+ 新增 API，用于配合 `Value.Seq / Seq2`，让非数组非切片类型也能像 iterator 一样迭代。

```
CanSeq() bool
```

判断是否可以使用 `Value.Seq` 进行遍历。

| 类型         | Seq 能否迭代        |
|------------|-----------------|
| 数组 `Array` | 可以              |
| 切片 `Slice` | 可以              |
| string     | 可以（按 rune）      |
| map        | 可以（单 key/value） |
| chan       | 可以（阻塞消费）        |

```
CanSeq2() bool
```

是否能用 `Value.Seq2` 双值迭代，和上面类似，但支持返回两个值，比如：

| 类型     | Seq 可迭代 | Seq2 可迭代 | 返回值含义       |
|--------|---------|----------|-------------|
| slice  | 支持      | 支持       | index,value |
| map    | 支持      | 支持       | key,value   |
| string | 支持      | 不支持      | 不可双值        |
| chan   | 支持      | 不支持      | 不可双值        |
