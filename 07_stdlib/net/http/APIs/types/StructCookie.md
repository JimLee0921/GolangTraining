# http.Cookie

`type Cookie` 是 net/http 里的结构体，用来表示 单个 HTTP Cookie。
既用于服务端设置响应 Cookie，也可用于客户端请求里携带 Cookie。

## 核心字段

```
type Cookie struct {
    Name     string
    Value    string
    Path     string
    Domain   string
    Expires  time.Time
    RawExpires string
    MaxAge   int
    Secure   bool
    HttpOnly bool
    SameSite SameSite // SameSiteDefaultMode/Lax/Strict/None
    Raw      string
    Unparsed []string
}
```

### `Name` / `Value`

* 必须设置 `Name`；`Value` 是字符串值
* 值里不要放逗号、分号、控制字符；需要复杂内容请自行编码（如 `url.QueryEscape`/base64）
* Go 会在序列化时尽量生成合法 `Set-Cookie` 头，但不会做业务编码

### `MaxAge` 与 `Expires`（寿命/过期）

* `MaxAge`（单位秒）优先级更高：

    * `> 0`：从现在起存活 N 秒（持久化）
    * `= 0`：不指定（会变成会话 Cookie，浏览器退出可能清除）
    * `< 0`：让浏览器立刻删除这个 Cookie（常用删除法）
* `Expires`（绝对时间）是历史兼容字段；现代实现以 `MaxAge` 为准。两者都设时，以 `MaxAge` 为准
* 删除 Cookie 的两种常见做法：

    * `MaxAge = -1`（推荐）
    * 或 `Expires` 设成过去的时间

### `Domain`（作用域：主域/子域）

* 不设置：成为`host-only` cookie，只发给当前完整主机名（不含子域）
* 设置（如 `example.com`）：会发给该域及其子域（`api.example.com` 等）
* 现代规范不要求在 `Domain` 前加点（`.example.com`），加不加效果等价
* 不要跨越公共后缀（如直接设 `com`），浏览器会拒绝

### `Path`（路径作用域）

* 限定 URL 路径前缀（默认一般是当前请求路径的目录部分）
* 例如 `Path=/app`，则只有访问 `/app/...` 时才会携带

### `Secure`

* 仅在 HTTPS 请求中发送。对含敏感信息的 cookie 必开

### `HttpOnly`

* 阻止前端 JavaScript 通过 `document.cookie` 读取，降低 XSS 窃取风险
* 强烈建*对会话/登录态开启

### `SameSite`

* 浏览器的 CSRF 对策。可选值
* 一般建议：会话态 `Lax`，如需跨站 iframe/POST 则用 `None+Secure`

### `Raw` / `Unparsed` / `RawExpires`

* 协议级保留/兼容字段：

    * `Raw`：原始的 `Name=Value` 片段
    * `Unparsed`：未能识别的属性列表
    * `RawExpires`：未解析的 Expires 文本
* 业务中很少直接用，更多是库内部或代理/中间件场景

## 相关包级函数

### `http.ParseCookie()`

用于解析请求头中的 Cookie: 行（不是 Set-Cookie）

```
func ParseCookie(line string) ([]*Cookie, error)
```

请求里的 Cookie 格式是：`Cookie: name1=value1; name2=value2; name3=value3`

- 同一行可能有多个 cookie
- Cookie 没有属性（Path/Domain/SameSite 等不会出现在 Cookie 头中）
- 所以 ParseCookie() 返回：[]*Cookie（多个）

**使用场景**

- 手写 HTTP 代理 / 抓包工具 / 日志分析工具
- 想直接解析原始请求头而不是通过 `r.Cookies()`

> 请求里的 Cookie 是压缩在一行中的，需要按 ; 分割，这就是 ParseCookie 做的事

### `http.ParseSetCookie()`

解析 响应头中的 Set-Cookie 行（服务端下发 Cookie 时用）

```
func ParseSetCookie(line string) (*Cookie, error)
```

响应格式是：`Set-Cookie: name=value; Path=/; Domain=example.com; HttpOnly; Secure; SameSite=Lax`

与请求侧不同：

- 一条 Set-Cookie 只设置一个 cookie
- 所以返回值是 *Cookie（单个）

**使用场景**

- 写反向代理、API 网关 时，需要解析并重写 `Set-Cookie` 属性
- 调试/检查 cookie 配置

## 主要方法

### String()

将 Cookie 序列化为 Set-Cookie 头格式的字符串。

```
func (c * Cookie) String() string
```

也就是把结构体变为：`name=value; Path=/; HttpOnly; Secure; SameSite=Lax; Max-Age=3600`

**使用场景**

- 服务端手写中间件，自己写 response header
- 做 cookie proxy / mirror / 复制 cookie
- 调试查看 cookie 内容

> String() 的输出用于 响应头，不是请求头。请求头 Cookie 的格式由 `req.AddCookie()` 负责，不用手写。

### Valid()

检查 Cookie 是否 符合规范，返回 nil 表示合法 cookie，error 表示不合法，不应该发送

```
func (c *Cookie) Valid() error
```

| 项                | 检查内容      |
|------------------|-----------|
| Name / Value     | 是否包含非法字符  |
| Path / Domain    | 是否是合法文本   |
| Expires / MaxAge | 是否符合规范    |
| SameSite         | 是否是有效选项   |
| Raw / Unparsed   | 若有冲突则校验失败 |

**使用场景**

- 生成 Cookie 时的健壮性检查
- 代理 / 中间件 在转发 Cookie 前验证安全性
- 防止写出非法 header 导致连接断开

## 补充

### 服务端设置 Cookie（响应）

* 使用 `http.SetCookie(w, &http.Cookie{...})`
* 这会在响应中写出一行：`Set-Cookie: ...`
* 多个 Cookie -> 多行 `Set-Cookie`

常见模式：

* 登录成功：设置会话 Cookie（`HttpOnly`, `Secure`, `SameSite=Lax` 或 `None+Secure`）
* 退出登录：同名 Cookie，`Path/Domain` 与原来一致，`MaxAge=-1`（或 `Expires` 过去）

### 客户端携带 Cookie（请求）

* 可以手工：`req.AddCookie(&http.Cookie{ Name, Value, ...})`
* 更推荐用 `Client.Jar`（实现 `CookieJar`），自动根据 `URL` 的 `Domain/Path/Secure` 规则存取与发送 Cookie（更像浏览器行为）

### Cookie 传输与匹配（你需要知道的规则）

* 请求头里是 `Cookie: name1=v1; name2=v2`（同名可能多值，行为由客户端决定；Go 在请求侧合并于一行）
* 响应头里是多行 `Set-Cookie`，每行只设置一个 cookie
* 浏览器/客户端发送时，会根据 `Domain/Path/Secure/SameSite` 选择是否携带
* 服务器匹配时，通常同名覆盖但要注意 Path 更具体者优先（浏览器行为层面）

### 安全与最佳实践

1. 会话 Cookie：`HttpOnly=true`、`Secure=true`、`SameSite=Lax`（默认安全）

    * 若必须跨站（如第三方嵌入）：`SameSite=None` 并确保 `Secure=true`
2. 删除 Cookie：同名同域同路径，`MaxAge=-1`（或 `Expires` 过去）
3. 敏感值：避免明文；至少做编码/签名，或只存短 token，服务端查真正数据
4. 作用域最小化：Domain/Path 设得尽量窄，减少泄露面
5. 大小限制：浏览器普遍限制单 Cookie ~4KB、每域总量有限；避免塞大数据

### 与 `CookieJar` 的关系

* `Cookie` 是单个 cookie 的结构描述
* `CookieJar` 是自动管理 cookie 的容器接口（记住/回放/持久化）
* 配合 `http.Client.Jar` 使用：`Do(req)` 时自动根据 URL 附带合规 cookie；`resp` 返回后自动保存 `Set-Cookie`

### type SameSite

提供了几个常量用于 Cookie 的 SameSite 字段使用。

```
type SameSite int

const (
 	SameSiteDefaultMode SameSite = iota + 1
 	SameSiteLaxMode
 	SameSiteStrictMode
 	SameSiteNoneMode
 )
```

* `SameSiteDefaultMode`：跟随浏览器默认（新浏览器多等价 Lax）
* `SameSiteLaxMode`：跨站 GET 导航可带，其他跨站请求不带（常用默认）
* `SameSiteStrictMode`：任何跨站都不带（最严格）
* `SameSiteNoneMode`：允许跨站发送，但必须配合 `Secure=true`（现代浏览器强制）

SameSite 允许服务器定义一个 cookie 属性，使浏览器无法将此 cookie 与跨站请求一起发送。
其主要目的是降低跨域信息泄露的风险，并提供一定的保护，防止跨站请求伪造攻击。

## http.CookieJar

主要用于在客户端侧自动管理 Cookie，配合 http.Client 一起使用。

CookieJar 是一个接口，作用是自动保存响应中的 Set-Cookie，并在下次请求时根据规则自动携带
Cookie，就像浏览器一样，可以维持登录态、会话、跨页面访问，而不用手动 AddCookie。

> CookieJar 是存储 + 匹配策略，不负责网络，不负责加密

### 主要字段

```
type CookieJar interface {
    SetCookies(u *url.URL, cookies []*Cookie)
    Cookies(u *url.URL) []*Cookie
}
```

只有两个方法：

| 方法                         | 作用                                 |
|----------------------------|------------------------------------|
| `SetCookies(url, cookies)` | 浏览器存 cookie：把 `Set-Cookie` 记下来     |
| `Cookies(url)`             | 浏览器带 cookie：根据 URL 取出应该附带的 cookies |

### 标准使用

```
import "net/http/cookiejar"

jar, _ := cookiejar.New(nil)
client := &http.Client{
    Jar: jar,
}
```

- 客户端收到响应 Set-Cookie -> 自动调用 jar.SetCookies 保存
- 下次请求同域名且路径匹配 -> 自动调用 jar.Cookies 附带 Cookie

> 可以自动维持登录状态，像浏览器一样

### 行为规则

CookieJar 在请求时会根据：

| 字段                 | 用于判断                                |
|--------------------|-------------------------------------|
| `Domain`           | 与 URL host 是否匹配（含不含子域）              |
| `Path`             | URL 的路径前缀是否匹配                       |
| `Secure`           | 如果 cookie 要求 HTTPS，而当前是 HTTP -> 不发送 |
| `SameSite`         | 控制跨站时是否发送（取决于是否是跨域请求）               |
| `MaxAge / Expires` | 若过期 -> 自动删除                         |

> CookieJar = 浏览器 Cookie 行为的核心实现

如果不用 CookieJar，必须

```
req.AddCookie(&http.Cookie{Name: "session", Value: "..."} )
```

每次手动取、手动带、手动更新、手动过期判断，非常麻烦。

CookieJar 自动做这些：

- 登录一次 -> 后续请求自动带 session cookie
- Cookie 过期 -> 自动删除，不会带旧 cookie
- 子域/路径 匹配正确
- HTTPS / SameSite 规则遵守

实际项目开发中，Client + CookieJar 是请求端的登录态保持方案

### 持久化存储

默认 `cookiejar.New(nil)` -> Cookie 只存内存，也就是程序退出就没了。

如果想跨进程持久化 Cookie（像浏览器一样），可以：

- 自己实现 CookieJar 接口 -> 存在 redis/sqlite/文件 
- 或用外部库（如 github.com/juju/persistent-cookiejar）