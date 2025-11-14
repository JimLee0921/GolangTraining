## `type URL`

URL 是一个结构体，用来 结构化地表示一个 URL。这些方法的存在目的主要是：

- 在结构体内部读取信息（而不是手动做字符串解析）
- 对 URL 做变更后能重新正确序列化
- 安全处理可能涉及敏感信息或编码问题

```
type URL struct {
	Scheme      string    // 协议，如 "http" / "https"
	Opaque      string    
	User        *Userinfo 
	Host        string    // 域名 + 端口，如 "example.com:8080"
	Path        string    // 路径，如 "/v1/user/info"
	RawPath     string    
	OmitHost    bool      
	ForceQuery  bool      
	RawQuery    string    // 原始查询字符串，如 "id=10&debug=true"
	Fragment    string    // #后的锚点，如 "section1"
	RawFragment string   
}
```

## 主要方法

### String()

将结构化 URL 重新组装成有效的 URL 字符串。

```
func (u *URL) String() string
```

当修改了 Scheme / Host / Path / RawQuery 等字段之后，
需要将其重新转换为字符串来发请求或打印。

打印格式有两种：

```
scheme:opaque?query#fragment
scheme://userinfo@host/path?query#fragment
```

第二种形式适用于以下规则：

- 如果 u.Scheme 是空的，那么 scheme: 这一部分会被省略
- 如果 u.User 是 nil，那么 userinfo@ 这一部分会被省略
- 如果 u.Host 是空的，那么 host/ 这一部分会被省略
- 如果 u.Scheme 和 u.Host 都是空的，并且 u.User 也是 nil，那么整个 scheme://userinfo@host/ 这一段都会被省略。
- 如果 u.Host 非空，并且 u.Path 以 / 开头，那么输出时不会额外再加一个 /（即不会变成 //host//path）
- 如果 u.RawQuery 为空，那么 ?query 这部分会被省略
- 如果 u.Fragment 为空，那么 #fragment 这部分会被省略

> URL.String() 会根据字段是否为空，自动决定是否输出对应部分。也就是说，URL 是按需组装的，不会多，也不会少

### `Query()`

用于获取查询参数集合。

```
func (u *URL) Query() Values
```

* 查询参数不是放在 `URL` 字段里，而是放在 `RawQuery` 字符串中
* `Query()` 的作用是把 `RawQuery` 解析成 `url.Values`（一个映射结构）
* 返回值是可读可写的 map，但修改后不会自动写回 URL

> 因此，常见流程是：获取->修改->重新编码->写回

### `Hostname()` / `Port()`

用于从 `Host` 中拆分出域名和端口。

```
func (u *URL) Hostname() string
func (u *URL) Port() string
```

* `Host` 字段可能同时包含域名 + 端口（如 `example.com:8080`）
* 这两个方法帮助无需手动分割字符串
* 更安全、更清晰地处理目标服务器信息

### `EscapedPath()`

用于获取编码后的路径表示。

```
func (u *URL) EscapedPath() string
```

* `Path` 有两种表示：原始（解码） 和 转义（编码）
* `Path`：人类可读，例如包含中文或特殊字符时是未转义的
* `EscapedPath()`：适合直接放在 HTTP 请求行里的版本，是正确的 URL 编码形式
* 区分这两种形式的理由：显示给人看 vs 发送给服务器的实际合法字符序列

### `Redacted()`

安全输出 URL（隐藏凭据）。

```
func (u *URL) Redacted() string
```

* 如果 URL 中包含 `Userinfo`（用户名/密码），普通 `String()` 会保留明文
* `Redacted()` 会自动把密码替换为 `xxxxx`
* 作用：用于日志和调试时避免泄露敏感信息

---

### `ResolveReference()`

合并相对 URL与基准 URL。

```
func (u *URL) ResolveReference(ref *URL) *URL
```

* 在浏览器、HTTP 重定向、爬虫中，经常会遇到相对路径URL
* `ResolveReference()` 的行为类似 HTML `<base>` 解析相对链接
* 它会根据 RFC 标准来正确拼接，而不是简单字符串合并
* 目的是：生成符合 URL 规范的绝对 URL

### `IsAbs()`

```
func (u *URL) IsAbs() bool
```

- 用于判断 URL 是否为绝对路径
- 绝对路径意味着 URL 具有非空的 scheme

### `JoinPath()`

```
func (u *URL) JoinPath(elem ...string) *URL
```

JoinPath 返回一个新的URL，其中提供的路径元素与任何现有路径连接，并且生成的路径会清除所有 `./` 或 `../` 元素。
任何多个 `/` 字符的序列都将被简化为单个 `/`。

> `url.JoinPath()` 函数底层调用的就是这个

### `EscapedFragment()`

获取片段（Fragment）的编码形式。

```
func (u *URL) EscapedFragment() string
```

通常，任何 Fragment 都有多种可能的转义形式。当 u.RawFragment 是 u.Fragment 的有效转义形式时，EscapedFragment 返回
u.RawFragment。否则，EscapedFragment 会忽略 u.RawFragment 并自行计算转义形式。URL.String方法使用 EscapedFragment
来构造其结果。通常，代码应该调用 EscapedFragment 而不是直接读取 u.RawFragment。

### `Parse()`

在 当前 URL u 的上下文基础上 解析另一个 URL（可能是相对路径）

```
func (u *URL) Parse(ref string) (*URL, error)
```

`URL.Parse(ref)` 本质上等价于 `u.ResolveReference(parsedRef)`，但是它把 ref 的解析和相对合并 合在一起了。

它将参数 ref 作为相对 URL 或绝对 URL解释：

- 如果 ref 是绝对 URL（包含 scheme），那么它直接返回它自己，不参考 u
- 如果 ref 是相对 URL，则会以调用者 u 为基准，进行路径/查询/片段级别的合并
- 可以使用路径符号 `./`、`../` 或 `/` 进行路径层级操作

**对比包级函数 `url.Parse()`**

| 对比点           | `url.Parse()`（包级函数） | `u.Parse(ref)`（URL 方法）   |
|---------------|---------------------|--------------------------|
| 是否依赖原 URL 上下文 | 不依赖                 | 依赖 `u` 作为基准              |
| 解析相对 URL      | 允许但不会自动合并           | 自动合并上级路径、host、scheme     |
| 常用场景          | 初次解析 URL            | 处理相对链接 / 路径跳转 / redirect |
