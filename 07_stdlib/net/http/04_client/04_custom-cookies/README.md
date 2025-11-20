## 请求添加 Cookie

在 Go 里，每个请求 *http.Request 都有一个字段：

```
Cookies []*http.Cookie
```

但通常不直接操作这个切片，而是使用专门的方法：

```
req.AddCookie(cookie)
```

或用 CookieJar 自动管理

> 其实 Cookie 最终也是加到了 req.Header 中，所以也可以直接从 req.Header 中添加 Cookie

### AddCookie

`req.AddCookie(c *http.Cookie)` 会自动在 Header 里生成：

```
Cookie: session_id=xyz-123; user=JimLee
```

- 自动拼接多个 cookie
- 自动转义值（=、; 等特殊字符）
- 避免手动管理 Header 拼接错误

> 不要直接用 req.Header.Set("Cookie", "...") 除非非常确定格式

### 读取服务器返回 Cookie

```
resp, _ := client.Do(req)
defer resp.Body.Close()

for _, c := range resp.Cookies() {
	fmt.Printf("%s = %s\n", c.Name, c.Value)
}
```

这会读取所有 `Set-Cookie` 头

### CookieJar

更推荐使用 CookieJar 自动管理 Cookie，这样就像浏览器一样自动保存、自动发送 Cookie。
Go 会根据 Domain / Path / Secure 等属性自动管理。

1. 自动保存服务器返回的 `Set-Cookie`
2. 在下次请求时，自动带上该域名下有效的 Cookie
3. 不用再手动 `AddCookie()`