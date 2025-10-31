## json.Decoder

`Decoder` 是专门用来 从 `io.Reader` 中流式解析 JSON 的解析器。

不像 `json.Unmarshal` 需要完整的 `[]byte`字节切片 在内存里，Decoder 可以边读边解析。

### 使用场景

| 场景        | 示例                           |
|-----------|------------------------------|
| 大文件       | 几百 MB 的日志 JSON               |
| 网络流       | HTTP 长连接持续传 JSON             |
| 连续多个 JSON | `{...}{...}{...}` 或 `NDJSON` |

### 创建

```
dec := json.NewDecoder(reader)
```

reader可以是

| 类型               | 示例    |
|------------------|-------|
| `os.File`        | 文件输入  |
| `net.Conn`       | 网络连接  |
| `bytes.Buffer`   | 内存缓冲区 |
| `strings.Reader` | 字符串流  |

### 基础用法

```
r := strings.NewReader(`{"name":"Alice","age":20}`)
dec := json.NewDecoder(r)

var u User
if err := dec.Decode(&u); err != nil {
    panic(err)
}

fmt.Println(u) // {Alice 20}
```

### 循环解析多个 JSON

如果输入流是多个 Json

```
{"name":"A","age":1}{"name":"B","age":2}{"name":"C","age":3}
```

可以这样：

```
dec := json.NewDecoder(r)
for {
    var u User
    if err := dec.Decode(&u); err != nil {
        if err == io.EOF { break }
        panic(err)
    }
    fmt.Println(u)
}
```

### 一行一个 JSON（NDJSON）

NDJSON 是`Newline-Delimited JSON`的缩写，意为换行符分隔的 JSON

比如有以下内容的文件 `users.ndjson`

```text
{"name":"A","age":1}
{"name":"B","age":2}
{"name":"C","age":3}
```

```
f, _ := os.Open("users.ndjson")
dec := json.NewDecoder(f)

for {
    var u User
    if err := dec.Decode(&u); err != nil {
        if err == io.EOF { break }
        panic(err)
    }
    fmt.Println(u)
}
```

> NDJSON 与日志流中极常见

### 数字转换问题

可以使用 `UseNumber()` 避免数字变成 float64

默认情况下：

```
map[string]any{"age": 18} // age 是 float64
```

`dec.UseNumber()` 告诉解析器：不要把数字直接变成 float64，而是先保留成 `json.Number`（其实是字符串形式）。
然后可以自己选择转成：

| 转换方式            | 说明          |
|-----------------|-------------|
| `age.Int64()`   | 转换成 int64   |
| `age.Float64()` | 转换成 float64 |
| `age.String()`  | 原样字符串，不损失精度 |

```
dec := json.NewDecoder(r)
dec.UseNumber()

var m map[string]interface{}
dec.Decode(&m)

age := m["age"].(json.Number)
i, _ := age.Int64()
fmt.Println(i) // 正确整数
```

### 严格模式

`DisallowUnknownFields()` 用来 禁止 JSON 中出现结构体里没有定义的字段。
如果 JSON 里多出任何结构体没有的字段，默认情况下 Go 会 忽略它们；而开启 `DisallowUnknownFields()` 开启严格模式后，解码会直接报错。

```
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func Demo() {
    data := []byte(`{"name":"Alice","age":20,"gender":"female"}`)

    var u User
    json.Unmarshal(data, &u)
    fmt.Printf("%+v\n", u)  // {Name:Alice Age:20}，gender 被悄悄忽略了。
}

func DemoStrict() {
    data := []byte(`{"name":"Alice","age":20,"gender":"female"}`)
    dec := json.NewDecoder(bytes.NewReader(data))
    dec.DisallowUnknownFields()

    var u User
    if err := dec.Decode(&u); err != nil {
        fmt.Println("error:", err)
        return
    }
    fmt.Printf("%+v\n", u)  // 开启了严格模式直接报错：error: json: unknown field "gender"
}


```

- 如果 JSON 中有结构体中不存在的 key 会直接报错，主要用于 API 输入校验非常实用
- 可以防止前端参数拼错字段名或有人恶意传递不期望处理的数据，提升数据安全性和可控性

### Token 模式（了解）

Token() 一次读取一个 JSON 符号（`{`, `}`, `[`, `]`, key, value...）

用于 只想读部分字段 / 快速跳过大块 JSON：

```
t, _ := dec.Token()
fmt.Printf("type=%T value=%v\n", t, t)
```