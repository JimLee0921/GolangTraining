## `net/url` 常用顶层函数

### 解析 URL / 查询串

#### `url.Parse`

```
func Parse(rawURL string) (*URL, error)
```

用于解析 URL / 字符串，把完整 URL 字符串解析成 `*url.URL` 结构体（最常用）。
通用解析（含 scheme/host/path/query/fragment）。

> 结果的 `u.Query()` 返回 `url.Values`，方便读写查询参数

#### `url.ParseRequestURI`

```
func ParseRequestURI(rawURL string) (*URL, error)
```

按 HTTP 请求更严格的规则解析 URI。通常用于校验/解析来自 HTTP 请求行或头里的 URL（更严格，能过滤一些非法形式）。
不允许 `//example.com` 这种不带 scheme 的 URL，HTTP 请求行场景更安全。

除此之外，与 Parse 对比时发现 ParseRequestURI 解析的 URL 的 RawQuery 中 # 号后面的字符串并未被截断，而 # 后的位置的标识符是不需要的。

> 比 Parse 更挑剔，一些宽松写法会报错

#### `url.ParseQuery`

```
func ParseQuery(query string) (Values, error)
```

用于解析查询参数，把 `a=1&b=2&b=3` 解析成 `url.Values`（`map[string][]string`）。
只拿到查询，始终返回一个非空映射表，其中包含所有找到的有效查询参数，
如果存在解码错误，则返回 err，表示遇到的第一个错误。

> 输入应是 不含问号 的原始查询串；需要反向生成使用 `Values.Encode()`

### 转义 / 反转义

Path 和 Query 的转义规则不同，不要混用

#### `url.PathEscape` / `url.PathUnescape`

```
func PathEscape(s string) string
func PathUnescape(s string) (string, error)
```

针对路径段的转义/反转义（例如把中文、空格、安全处理为路径安全表示）。
常用于将变量拼入 Path（如 `/files/<filename>`。

> 不要用 QueryEscape 来转义路径，会产生不期望的加号等

#### url.QueryEscape / url.QueryUnescape

```
func QueryEscape(s string) string
func QueryUnescape(s string) (string, error)
```

针对 查询参数值 的转义/反转义。
常用于手工构造 RawQuery 或单独处理某个参数值。

> 在查询参数中，空格会被编码为 +；所以路径场景不要使用

### 路径拼接 `url.JoinPath`

Go 1.19+ 引入

```
func JoinPath(base string, elem ...string) (result string, err error)
```

安全地把多个路径段拼成一个 URL 字符串（自动处理斜杠、`.`、`..` 等）。
常用于在已知 base（可能带路径），还要追加 `/v1/users/123` 这类段落。

> 它返回 字符串 而不是 `*url.URL`；需要进一步改查询参数时，通常先 Parse 再改 若手里是 `*url.URL`，也可以直接改 `u.Path` 或用
> `path.Join(u.Path, more...)`，但转义要自己处理；JoinPath 会更省心。

### 构造 Userinfo（用户名/密码）

使用 url.User / url.UserPassword 进行构造用户名和密码

```
func User(username string) *url.Userinfo
func UserPassword(username, password string) *url.Userinfo
```

构造 `*url.Userinfo` 给 `u.User` 字段用（如 `https://user:pass@host`）。
需要在 URL 中放凭据（现代 HTTP 基本不推荐在 URL 里放敏感信息）时使用。

> 包含敏感信息，`u.Redacted()` 可打印已打码的 URL
