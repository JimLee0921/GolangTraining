## struct tag

Go 结构体字段必须大写开头才能被序列化或参与反序列化。

但 JSON 等通常使用其它命名风格，比如小写+下划线风格：user_name, age, created_at，所以为了控制 JSON 字段名序列化和反序列化，可以使用结构体
tag

在 Go 中，可以给结构体字段添加一段特殊的字符串：

```
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

这段反引号 `` `...` `` 中的部分，就是 struct tag（结构体标签）。
它本质上是字段的额外注解（metadata），不会影响代码逻辑，但可以被反射读取。

> 用途：struct tag 让结构体在不同库或框架之间传递字段配置信息。
> 常用于：
>
> * JSON / XML 序列化
> * 数据库 ORM 映射（如 `gorm:"column:user_name"`）
> * 表单解析（如 `form:"username"`）
> * 验证框架（如 `validate:"required,email"`）

---

## 语法形式

结构体标签写在字段定义的最后，用反引号包裹：

```
FieldName Type `key1:"value1" key2:"value2"`
```

示例：

```
type Product struct {
    ID   int    `json:"id" db:"product_id"`
    Name string `json:"name" validate:"required"`
}
```

这里的标签字符串是：

```
json:"id" db:"product_id" validate:"required"
```

它由多个键值对组成，中间用空格分隔。

---

## 读取 struct tag（反射）

标签信息可以通过 `reflect` 包在运行时读取：

```
import "reflect"

type User struct {
    Name string `json:"name" db:"username"`
    Age  int    `json:"age"`
}

func main() {
    t := reflect.TypeOf(User{})
    f, _ := t.FieldByName("Name")
    fmt.Println(f.Tag.Get("json")) // 输出: name
    fmt.Println(f.Tag.Get("db"))   // 输出: username
}
```

* `f.Tag` 是 `reflect.StructTag` 类型，本质是字符串。
* `Tag.Get("key")` 可以安全地提取键对应的值。
* 如果 key 不存在，则返回空字符串。

---

## 标准库的 struct tag 示例

### 1. JSON 序列化

```
type User struct {
    Name string `json:"name,omitempty"`
    Age  int    `json:"age"`
}
```

* `json:"name"` -> 序列化时字段名改为 `"name"`；
* `omitempty` -> 当字段为空（零值）时跳过；
* `-` -> 忽略该字段。

### 2. 数据库 ORM（例如 GORM）

```
type User struct {
    ID   int    `gorm:"column:id;primaryKey"`
    Name string `gorm:"column:user_name"`
}
```

框架通过解析 tag 来生成 SQL 语句。

### 3. 表单解析 / 验证

```
type LoginForm struct {
    Email    string `form:"email" validate:"required,email"`
    Password string `form:"password" validate:"required,min=8"`
}
```

`form` 指定请求参数名，`validate` 告诉验证库如何校验。

---

## 结构体标签的底层原理

在编译期，Go 编译器会把 tag 信息**写入类型元数据中**，
运行时通过 `reflect` 就能取出这些字符串。

每个字段在反射信息里是一个 `StructField`，
定义如下（简化）：

```
type StructField struct {
    Name string
    Type reflect.Type
    Tag  StructTag
}
```

`StructTag` 本质上是一个字符串，但 Go 提供了便利函数：

```
type StructTag string

func (tag StructTag) Get(key string) string
func (tag StructTag) Lookup(key string) (string, bool)
```

---

## 注意事项

| 限制 / 注意点        | 说明                        |
|-----------------|---------------------------|
| 反引号 `` `...` `` | 只能用反引号包裹，不能用双引号           |
| key/value 必须紧凑  | 必须是 `key:"value"` 格式      |
| tag 不参与逻辑       | 编译器忽略 tag 内容，只是附加信息       |
| 字段需导出           | 只有导出字段（首字母大写）才能被反射读取到 tag |
| 相同 key 多次定义     | 只会读取第一个匹配的 key            |

---

## tag 的多键与组合

标签中可以包含多个键：

```
type User struct {
    Name string `json:"name" xml:"username" db:"user_name"`
}
```

读取时：

```
f, _ := reflect.TypeOf(User{}).FieldByName("Name")
fmt.Println(f.Tag.Get("xml")) // username
```

---

## 典型示例小结

| 框架/库           | Tag 示例                               | 含义               |
|----------------|--------------------------------------|------------------|
| encoding/json  | `json:"name,omitempty"`              | 控制 JSON 字段名与省略行为 |
| gorm           | `gorm:"column:user_name;primaryKey"` | 指定数据库列名与主键       |
| validator.v10  | `validate:"required,email"`          | 校验规则             |
| form / binding | `form:"username"`                    | 表单字段名映射          |
| bson (MongoDB) | `bson:"_id,omitempty"`               | 控制 BSON 序列化      |
| yaml           | `yaml:"config_path"`                 | YAML 映射字段名       |

