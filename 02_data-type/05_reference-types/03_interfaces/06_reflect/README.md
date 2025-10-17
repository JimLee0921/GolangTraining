## 反射（Reflection）

反射（reflection） 是 Go 提供的一种强大机制，允许程序在运行时动态地检查类型信息，以及操作变量的值。

换句话说：

> 反射可以在运行时知道一个变量的类型（type）和值（value），甚至修改它

反射的核心包是：`import "reflect"`

> 文档地址 [Go Reflect](https://pkg.go.dev/reflect)


---

## 反射的三大核心概念

Go 反射主要围绕以下三个类型展开：

| 概念   | 类型              | 作用        |
|------|-----------------|-----------|
| 类型信息 | `reflect.Type`  | 描述变量的类型   |
| 值信息  | `reflect.Value` | 描述变量的值    |
| 接口入口 | `interface{}`   | 所有类型的通用容器 |

---

## 反射操作

### 基础反射

反射最核心的两个概念就是 Type 与 Value ：

- reflect.Type：表示类型信息（如 int、float64、[]string 等）
- reflect.Value：表示值信息（即变量当前持有的内容）

所有反射操作的入口都是一个 `interface{}` 类型的值。

```
var x float64 = 3.4
t := reflect.TypeOf(x)
v := reflect.ValueOf(x)
```

这里 t 获取的是类型（float64），
而 v 则是一个可以操作的值对象，通过 v.Float() 能取出具体数值。

结论：

- TypeOf 查看类型
- ValueOf 查看值
- v.Kind() 能看到类型的种类（底层类别，如 int、slice、struct）。

### 从反射对象回到普通变量

* **从 `reflect.Value` 转回具体值**：

```
v := reflect.ValueOf(3.4)
f := v.Float() // 取出 float64
fmt.Println(f) // 3.4
```

* **取回 interface{} 再断言**：

```
x := v.Interface()
fmt.Println(x.(float64)) // 3.4
```

### 反射修改变量值

反射默认得到的 Value 是只读的，如果想通过反射修改变量，必须传入变量的指针。

1. `使用 reflect.ValueOf()` 时传入指针 &x
2. 使用 `v.Elem` 返回接口 v 包含的值或指针 v 指向的值，用于解引用指针，类似普通语法中的 *p
3. 只有当 v.CanSet() 为 true 时，才能修改值

```
x := 3.4
v := reflect.ValueOf(&x) // 传入指针
v = v.Elem()             // 取出指针指向的值

if v.CanSet() {
	v.SetFloat(777)
}
fmt.Println(x) // 7.1
```

- ValueOf(x) -> 不可修改
- ValueOf(&x).Elem() -> 可修改
- SetXXX() 方法必须与类型匹配（如 SetFloat、SetInt、SetString）

### Type 与 Kind 的区别

Go 的类型系统分两层：

* `Type`：完整类型，如 `[]int`, `map[string]int`, `MyStruct`
* `Kind`：底层类别，如 `slice`, `map`, `struct`

```
arr := []int{1, 2, 3}
t := reflect.TypeOf(arr)
fmt.Println(t)        // []int
fmt.Println(t.Kind()) // slice
```

> Type -> 表示是什么类型
> Kind -> 表示属于哪一类类型

### 结构体字段反射与Tag

反射可以获取结构体的字段、方法、tag 等信息，通过 reflect.TypeOf(struct) 可以获取结构体类型，
使用 NumField() 和 Field(i) 可以遍历结构体的所有字段。

```
type User struct {
Name string `json:"name"`
Age  int    `json:"age"`
}
```

循环中：

```
for i := 0; i < t.NumField(); i++ {
    f := t.Field(i)       // 字段定义信息
    v := value.Field(i)   // 字段对应的值
    f.Tag.Get("json")     // 取出标签
}
```

- Field(i) -> 字段定义信息
- Value.Field(i) -> 字段值
- Tag.Get("json") -> 读取结构体标签内容
- 框架（如 ORM、序列化）通常依赖这个特性

### 通过反射修改结构体字段值

修改结构体的字段值时，也必须先传入结构体指针。

```
u := User{"Tom", 18}
v := reflect.ValueOf(&u).Elem()

nameField := v.FieldByName("Name")
ageField := v.FieldByName("Age")

nameField.SetString("Jerry")
ageField.SetInt(30)
```

- 只能修改导出字段（即首字母大写的字段）
- FieldByName("Name") 按名称访问
- CanSet() 判断字段是否可写

### 动态调用函数

反射允许在运行时调用函数，即使不知道函数的具体类型，只要能获得它的引用。

```
func Add(a, b int) int { return a + b }

v := reflect.ValueOf(Add)
args := []reflect.Value{
    reflect.ValueOf(2),
    reflect.ValueOf(3),
}
result := v.Call(args)
fmt.Println(result[0].Int()) // 5
```

- reflect.ValueOf(func) -> 获取函数
- Call([]reflect.Value) -> 执行
- 返回值是 []reflect.Value 数组
- 框架可通过此方式调用任意方法

---

## 反射的注意事项

| 注意点     | 说明                 |
|---------|--------------------|
| 性能低     | 反射会带来运行时开销，不适合频繁调用 |
| 可读性差    | 代码难理解，不易维护         |
| 类型安全性丧失 | 反射绕过编译器的类型检查       |
| 用法建议    | 一般只在框架底层或通用库中使用反射  |

---

## 核心函数速查表

| 函数                    | 说明                             |
|-----------------------|--------------------------------|
| `reflect.TypeOf(i)`   | 获取类型                           |
| `reflect.ValueOf(i)`  | 获取值                            |
| `v.Kind()`            | 获取值的基础类别                       |
| `v.Interface()`       | 转回 interface{}                 |
| `v.CanSet()`          | 判断是否可修改                        |
| `v.SetXXX()`          | 修改值（XXX 可为 Int、String、Float 等） |
| `t.NumField()`        | 获取字段数                          |
| `t.Field(i)`          | 获取字段信息                         |
| `v.FieldByName(name)` | 按字段名取值                         |
| `v.Call([]Value)`     | 调用函数                           |

