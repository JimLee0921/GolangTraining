## 基础序列化（Marshal）和反序列化（Unmarshal）

## `Marshal` 序列化

在 Go 中要把一个结构体转成 JSON 字符串（序列化），用 `json.Marshal` 或 `json.MarshalIndent`

### json.Marshal

```
func Marshal(v any) ([]byte, error)
```

- 只会序列化可导出的字段（首字母要大写）
- 可以输入任意可序列化的 Go 值（结构体、map、slice、基本类型等）
- 字段名默认使用结构体字段名（如果没有 json tag 标签）
- 输出是紧凑格式，没有换行和缩进
- 返回的是 []byte 类型的 JSON 数据，通常要 string() 转成可读文本
- 如果结构体里有不支持的类型，会返回 `error`

### json.MarshalIndent

```
func MarshalIndent(v any, prefix, indent string) ([]byte, error)
```

- prefix：每行行首前缀（一般设为 ""）
- indent：缩进字符（常用 " " 或 "\t"）

| 方法                   | 输出格式       | 适用场景        |
|----------------------|------------|-------------|
| `json.Marshal`       | 一行，无空格，无换行 | 网络传输、存储、最常用 |
| `json.MarshalIndent` | 多行、带缩进     | 人类查看、调试、日志  |

## `Unmarshal` 反序列化

JSON -> 结构体（反序列化），核心函数是：

```
func Unmarshal(data []byte, v any) error
```

- `v any`：必须是指针，因为需要修改传入的变量中的值。

### 基础反序列化

```
err := json.Unmarshal([]byte(jsonStr), &v)
```

Go 结构体字段同样必须首字母大写才能参与反序列化

> JSON 动态结构下章讲解

### JSON 数组反序列化

```
jsonStr := `[{"name": "A", "age": 10}, {"name": "B", "age": 20}]`

var users []User
json.Unmarshal([]byte(jsonStr), &users)

fmt.Println(users)
// [{A 10} {B 20}]
```

## Json Tags

Go 结构体字段必须大写开头才能被序列化，但 JSON 通常使用 小写 + 下划线 风格如：user_name, age, created_at。

为了控制 JSON 字段名，可以使用结构体 tag。

> 具体结构体 tags 见 struct 学习中的 tags 章节，这里主要列出 json 序列化反序列化中中常用的 tag

`encoding/json` 本身识别基类标签写法：

| 写法                      | 示例                                      | 作用                | 序列化            | 反序列化           |
|-------------------------|-----------------------------------------|-------------------|----------------|----------------|
| `json:"name"`           | `Name string "json:\"name\""`           | **字段名映射**         | 使用 `"name"` 作键 | 读取 `"name"` 字段 |
| `json:"name,omitempty"` | `Age int "json:\"age,omitempty\""`      | **零值不输出**         | 零值会被省略         | 正常解析           |
| `json:"-"`              | `Password string "json:\"-\""`          | **字段完全忽略**        | 不输出            | 不写入            |
| `json:"name,string"`    | `ID int64 "json:\"id,string\""`         | **数字/布尔以字符串形式传输** | 输出 `"123"`     | 从 `"123"` 解析   |
| `json:",omitempty"`     | `Nickname string "json:\",omitempty\""` | 字段名不变，但 **零值不输出** | 省略零值           | 正常解析           |

> 标签只能控制「可访问字段」的名称和行为，不能改变字段的可见性。如果是未导出字段，反射看不到，就无法读写

```
type User struct {
    name string `json:"name"` // 无效
    Age  int    `json:"age"`  // 有效
}
```

## 补充

### 必须首字母大写才能进行序列化和反序列化

Go 的 JSON 是基于 反射 (reflection) 实现的。只有 导出的字段（public，开头大写）才能在反射中访问。

```
type User struct {
    name string // 不会被序列化和反序列化
    Age  int    // 会被序列化和反序列化
}
```