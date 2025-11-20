# `type Header`

> `type Header` Go 的 net/http 中非常基础，但后面写 Request / Response / Handler 都会一直用到。

在 HTTP 协议中，每一次请求和响应，除了主体内容（Body）以外，还会携带一些 键值对格式的元信息，比如：

```
Content-Type: application/json
User-Agent: curl/7.88
Accept-Encoding: gzip
```

这些就是 HTTP 头部（Header）。
在 Go 里，它的类型定义是：

```
type Header map[string][]string
```

- Header 本质上是一个 map
- key 是 string（例如 `Content-Type`, `User-Agent`）
- value 是 string 列表 []string（因为 HTTP 运行同一个 header 字段可以出现多次）
- HTTP 头字段对大小写不敏感，所以 `content-type`、 `content-type`、 `CONTENT-TYPE` 是一样的，并且 Go
  在内部自动规范化为首字母大写形式，所以不需要自己处理大小写

## Go 中使用场景

`type Header` 主要用于存储 HTTP 请求和响应的元信息字典，本质是 `map[string][]string`，用于读取客户端传来的头部和设置服务器返回的头部信息

| 角色                | 存放方式         | 示例                                                   |
|-------------------|--------------|------------------------------------------------------|
| 客户端发送请求（Request）  | `req.Header` | `req.Header.Get("User-Agent")`                       |
| 服务端接收请求（Handler里） | `r.Header`   | `r.Header.Get("Content-Type")`                       |
| 服务端返回响应（Response） | `w.Header()` | `w.Header().Set("Content-Type", "application/json")` |

### 请求头读取

```
func handler(w http.ResponseWriter, r *http.Request) {
    ua := r.Header.Get("User-Agent")
    fmt.Println("Client:", ua)
}
```

### 响应头设置

```
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("Hello"))
}
```

## 主要方法

`Header` 是 `map[string][]string`，但标准库为它补充了一些常用方法，用来更方便地获取、设置、添加和删除 HTTP 头部字段。

### `Get()`

```
func (h Header) Get(key string) string
```

* 用来读取头部字段的值
* 如果该字段有多个值，只会返回第一个值
* 如果该头字段不存在，则返回空字符串
* **注意**：返回值是逻辑上的字符串视图，不会改变内部存储结构

### `Set()`

```
func (h Header) Set(key, value string)
```

* 用来设置/覆盖一个头部字段
* 如果该字段之前存在多个值，调用后会只保留一个值
* 适用于这个头部只需设置一个值的场景，例如 `Content-Type`

### `Add()`

```
func (h Header) Add(key, value string)
```

* 用来追加一个值，而不是覆盖
* 如果该字段本身允许或需要多个值（如 `Set-Cookie`），使用此方法
* 调用后该 key 会对应一个字符串切片，保存多个值

### `Del()`

```
func (h Header) Del(key string)
```

* 删除该 key 对应的所有值
* 删除后，`Header` 中不会再出现该字段

### `Values()`

```
func (h Header) Values(key string) []string
```

* 获取该 key 下的所有值，以切片形式返回
* 可用于处理多值头部字段

### `Write()`

```
func (h Header) Write(w io.Writer) error
```

* 将整个 header 格式化为标准 HTTP 头部格式并写入底层流
* 每一对值会写成：`Key: value1, value2, ...`
* 一般由 HTTP 内部使用，不需要手动调用

### `WriteSubset()`

```
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
```

* 与 `Write` 类似，但可选择排除指定字段
* HTTP 内部用于控制一些关键头不被重复写出

### `Clone()`

```
func (h Header) Clone() Header
```

* Clone 用来 深拷贝 一个 Header，返回一个新的 Header 对象（如果 Header 为 nil 则返回 nil）
* 返回的新 Header 和原来的 不是同一个 map，key -> value 的每个 []string 也会被复制一份
* 修改新副本 不会影响原 Header，修改原 Header 也不会影响副本
* 主要用于安全传递 header、避免共享数据被篡改

## 补充

### 大小写规则说明

虽然 HTTP 头字段大小写不敏感，但 `Header` 会将 key 规范为首字母大写+中划线分隔的形式存储，例如：

```
content-type -> Content-Type
user-agent -> User-Agent
```

不需要手动处理大小写转换，`Header` 内部自动保证一致性。

### 设置时机
