# reflect.Method

Method 是结构体、接口、任意类型中 导出方法的抽象类型描述体，反射中非常核心的对象，主要用于：

- 动态获取类型上有哪些方法
- 查看每个方法的参数 / 返回值签名
- 高级应用搭配 `Value.Method(i).Call()` 动态调用方法

## type Method 结构

```
type Method struct {
    Name    string      // 方法名（必须是导出 = 大写）
    PkgPath string      // 方法所在包，如果为空表示导出
    Type    Type        // 方法完整类型 = func(Recv, Args...) Returns...
    Func    Value       // 可调用的方法 Value（静态函数入口）
    Index   int         // 方法索引，Method(i) 时使用
}
```

| 字段        | 含义                        | 关键理解                                    |
|-----------|---------------------------|-----------------------------------------|
| `Name`    | 方法名                       | 只包含 导出方法（大写），小写方法反射不可见                  |
| `PkgPath` | 包路径                       | `""`=导出，非空=包私有方法                        |
| `Type`    | 方法完整类型签名                  | 包含 receiver 作为第一个入参                     |
| `Func`    | 方法类型，可调用的 `reflect.Value` | 可以 `.Call()` 调用                         |
| `Index`   | 方法编号                      | 用于 `Type.Method(i)` / `Value.Method(i)` |

> Method 只有一个方法：`func (m Method) IsExported() bool`，用于获取该方法是否已导出

## Method 获取与使用

使用 `reflect.TypeOf(x)` 获取方法信息，使用 `reflect.ValueOf(x)` 进行方法的调用

| 访问来源          | 获取对象                       |
|---------------|----------------------------|
| `t.Method(i)` | 方法结构描述（Method type 信息）     |
| `v.Method(i)` | 方法可调用 Value（没有 Type 信息！！！） |

```
t := reflect.TypeOf(v)
v := reflect.ValueOf(v)
```

### 获取方法信息

使用 `reflect.TypeOf(x)`可以获取方法信息（也就是 Method 结构体）

```
t.NumMethod()         // 方法数量
t.Method(i)           // 根据索引取方法结构体，可以访问上面的 Name / PkgPath 等属性
t.MethodByName("XXX") // 按名称查找
```

### 动态调用方法

通过 `reflect.ValueOf(x)` 可以调用方法

```
m := v.MethodByName("Hello") // 方法 Value

result := m.Call([]reflect.Value{
    reflect.ValueOf("JimLee")
})

fmt.Println(result[0].String()) // Hello, JimLee
```

> 在 Value 章节会再学习

| 行为                           | 说明         |
|------------------------------|------------|
| Call 参数必须是 `[]reflect.Value` | 不支持直接传原生类型 |
| 返回值也是 `[]reflect.Value`      | 要自己取出并转型   |
| 不支持调用未导出方法                   | 反射权限受限     |
