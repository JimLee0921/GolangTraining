## `Encoder`

`json.Encoder` 用来把 Go 值以 JSON 格式写入一个 `io.Writer`，并且可以连续写、不需要构造完整字符串。

常用在：

| 场景                 | 示例               |
|--------------------|------------------|
| HTTP 接口返回 JSON     | `ResponseWriter` |
| 向文件写入大 JSON 数据     | `os.File`        |
| 实现日志 / 数据流（NDJSON） | 一条条写，自动换行        |

### 创建 Encoder

```
enc := json.NewEncoder(writer)
```

`writer` 可以是：

| 类型                    | 示例            |
|-----------------------|---------------|
| `http.ResponseWriter` | Web 服务返回 JSON |
| `os.File`             | 写入本地文件        |
| `net.Conn`            | 网络通信 JSON 流   |
| `bytes.Buffer`        | 缓存到内存         |

### 核心方法

```
err := enc.Encode(v)
```

* 会将对象 `v` 转为 JSON 并写入 `writer`
* 自动在末尾加一个 `\n` 换行（对流式 / 日志特别有用）

## 示例

```go
package main

import (
	"encoding/json"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{"Alice", 20}
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(u) // 写到控制台
}
```

输出：

```
{"name":"Alice","age":20}
```

---

### 连续写多个 JSON

```
enc := json.NewEncoder(os.Stdout)

enc.Encode(User{"A", 10})
enc.Encode(User{"B", 20})
enc.Encode(User{"C", 30})
```

输出：

```
{"name":"A","age":10}
{"name":"B","age":20}
{"name":"C","age":30}
```

这等价于 NDJSON（Newline Delimited JSON），
常用于日志系统、数据流、WebSocket、SSE 等。

### 美化输出

可以使用 `SetIndent()` 方法进行美化输出

```
enc := json.NewEncoder(os.Stdout)
enc.SetIndent("", "  ")
enc.Encode(User{"Alice", 20})
```

输出：

```json
{
  "name": "Alice",
  "age": 20
}
```

> 注意：如果是生产环境（大数据/高频日志），不建议使用 SetIndent，因为会增加输出大小

---

### 关闭 HTML 转义

默认情况下：

```
"<script>" 会被转为 "\u003cscript\u003e"
```

如果不想转义（前端渲染更自然）可以使用 `SetEscapeHTML()` 方法关闭转义

```
enc := json.NewEncoder(w)
enc.SetEscapeHTML(false)
```

| 原字符 | 被转义为     |
|-----|----------|
| `<` | `\u003c` |
| `>` | `\u003e` |
| `&` | `\u0026` |

### Encoder vs Marshal 对比

| 功能                | `json.Marshal` | `json.Encoder.Encode` |
|-------------------|----------------|-----------------------|
| 输出位置              | 内存 `[]byte`    | 写入 `io.Writer`        |
| 是否流式              | 否              | 是                     |
| 是否自动换行            | 否              | 是                     |
| 大数据场景             | 不适合            | 最佳                    |
| NDJSON/log stream | 不方便            | 非常方便                  |

```
Marshal -> 得到字符串（内存）
Encoder -> 写出去（流式）
```