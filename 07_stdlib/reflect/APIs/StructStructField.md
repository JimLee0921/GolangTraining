# reflect.StructField

反射读取结构体字段信息的核心数据结构，用得非常频繁，`ORM/JSON/DI`都围绕它展开，歪瑞重要！

## 结构

想了解 struct 的字段，StructField 就是信息源。

```
type StructField struct {
    Name      string      // 字段名
    PkgPath   string      // 包路径，非导出字段才有
    Type      Type        // 字段类型（reflect.Type）
    Tag       StructTag   // 字段tag，`xxx:".."` 就是这里
    Offset    uintptr     // field 在内存中的偏移量
    Index     []int       // 匿名字段层级索引路径
    Anonymous bool        // 是否匿名字段（embedding）
}
```

| 字段          | 意义                 | 重点用途                 |
|-------------|--------------------|----------------------|
| `Name`      | 字段名（必须导出才能访问）      | 用于字段定位、序列化、ORM映射     |
| `PkgPath`   | 非导出字段才有包路径         | 判断字段是否可导出访问          |
| `Type`      | 字段类型（reflect.Type） | 用于判断类型、生成实例、解码       |
| `Tag`       | 字段标签 `tag`         | JSON/GORM/YAML注解解析核心 |
| `Offset`    | 字段相对结构体内存偏移        | 做unsafe优化或写ORM驱动时用   |
| `Index`     | 字段层级索引（支持匿名字段）     | 查找嵌套字段必需             |
| `Anonymous` | 是否为匿名字段（嵌入）        | Go继承/embedding支持的关键  |

## 反射体系的定位

在反射体系两大核心概念 Type 和 Value 中

- Type = 知道这个 struct 有哪些字段，而 StructFiled 就是每个字段的描述（`Name/Type/Tag/Index/Offset`），主要是查看
- Value = 字段的实际值（`Value.Filed(i)`），可以进行读/改值

## 获取 StructFiled

StructField 永远来自 `reflect.Type`

```
t := reflect.TypeOf(User{})
t.NumField()  // 获取到 Field 数量，可以用于遍历
```

1. 按索引获取：使用 `t.Field(i)` 按索引取字段，最常用，0 到 `NumField()-1`
2. 按名称取字段：使用 `t.FieldByName(name)` 按字段名取字段，如果找不到返回 false
3. 按嵌套路径取字段（高级用法）：使用 `t.FieldByIndex([]int{...})`，支持匿名字段层级索引

## 获取实际运行值

使用 `reflect.Value` 获取字段实际上运行的值并进行类型转换，也可以选择使用指针 + `Elem()` 进行修改运行值。

1. `Field(i)`：按索引获取字段 Value，最基本、遍历使用
2. `FieldByName(name)`：按字段名取值，字段名已知但顺序未知时使用
3. `FieldByIndex([]int)`：按层级路径取嵌套字段，匿名字段 / 多层嵌套 struct 访问
4. `FieldByNameFunc(fn)`：依据自定义规则搜索字段，模糊匹配 / 自定义过滤逻辑

> 在 Value 章节还会学习

### 类型转换

1. 使用对应类型的 Getter (推荐)进行修改：
    - `v.Filed(i).Int()`
    - `v.Field(i).String()`
    - `v.Field(i).Bool()`
    - `v.Field(i).Float()`
    - ...

2. 转成 interface 再断言：`val := v.Field(i).Interface().(int)`
    - 适合不确定类型或泛型场景

### 修改字段运行值

想修改必须用 指针 + Elem()

```
u := User{"JimLee", 22}

v := reflect.ValueOf(&u).Elem()  // 必须传指针

v.FieldByName("Age").SetInt(99)  // 修改运行值
fmt.Println(u.Age) // 99 已经改掉
```

- Value = 操作运行数据
- Pointer = 授权你修改它

> ValueOf(u) 拿到的是副本，而 `ValueOf(&u).Elem()` 才是真对象可写视图

## 相关函数和方法

### `reflect.VisibleFields`

`VisibleFields` 顶级函数接受一个 `reflect.Type`（且 Kind 必须是 struct），返回一个 `[]StructField` 列表。

返回的是所有可见 (visible) 字段 (fields)，不只是当前 struct 自己声明的字段，还包括匿名嵌入 (embedded / anonymous) struct
中提升 (promoted)出来的字段 (promoted fields) ，
并且正确处理字段名冲突 (覆盖 /隐藏) 的情况。

```
func VisibleFields(t Type) []StructField
```

### IsExported 方法

用来判断该字段是否为导出 (exported) 字段。也就是 Go 语言中大写开头、外包可访问的字段。

Go1.17 之前，判断是否导出通常是看 `StructField.PkgPath` 是否为空 (`PkgPath == ""` 表示导出) 这也一直是 reflect 的惯例。

`IsExported()`使得代码更加语义化，如果用到了私有字段/导出字段过滤，就写 `if f.IsExported() {...}`，比 `f.PkgPath == ""`。

```
func (f StructField) IsExported() bool
```