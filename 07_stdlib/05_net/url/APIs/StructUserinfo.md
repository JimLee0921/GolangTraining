## `type Userinfo`

Userinfo 表示 URL 中的用户凭据信息：

```
scheme://user:password@host/path
         ↑   ↑
         |   └── password（可选）
         └────── username
```

在 Go 中定义为：

```
type Userinfo struct {
    username string
    password *string
}
```

只存两个东西： username 和 password（可选）

### 创建

永远不会直接构造它（因为它是内部结构），而是通过两个工厂函数：

| 函数                                     | 用途      | 场景                  |
|----------------------------------------|---------|---------------------|
| `url.User(username)`                   | 只带用户名   | 例如：`user@host`      |
| `url.UserPassword(username, password)` | 带用户名和密码 | 例如：`user:pass@host` |

示例：

| 调用                                | 构成的 URL 片段  |
|-----------------------------------|-------------|
| `url.User("root")`                | `root@`     |
| `url.UserPassword("root", "123")` | `root:123@` |

### 主要方法

| 方法                                             | 返回内容                       | 说明                  |
|------------------------------------------------|----------------------------|---------------------|
| `func (u *Userinfo) Username() string`         | 返回用户名（string）              | 无秘密                 |
| `func (u *Userinfo) Password() (string, bool)` | `password string, ok bool` | `ok = false` 表示没有密码 |
| `func (u *Userinfo) String() string`           | `user:pass`（不打码）           | 不要直接用于日志            |

通常使用 `type URL` 的 `Redacted()` 进行安全输出，可以隐藏密码：`user:xxxxx@host`

> username 和 password 分开存储是因为 password 是可选的。
> 很多 URL 只需要用户名（例如 SSH、ftp、git、匿名登录等）