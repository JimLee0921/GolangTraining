# HTTP 顶层函数

在真正深入 `http.Server` / `Handler` / `Router` 之前，先把 `net/http` 包提供的顶层函数搞清楚会对整体运行流程更清晰。

---

## 服务端启动类（Server Start Functions）

| 函数                                                         | 用途                      | 备注                                     |
|------------------------------------------------------------|-------------------------|----------------------------------------|
| `http.ListenAndServe(addr, handler)`                       | 启动 HTTP 服务器             | 最常用一行式启动方式                             |
| `http.ListenAndServeTLS(addr, certFile, keyFile, handler)` | 启动 HTTPS 服务器            | PEM 格式证书+私钥文件                          |
| `http.Serve(l net.Listener, handler)`                      | 使用已有 Listener 启动 Server | 更高级控制方式（常用于自定义端口、代理、graceful shutdown） |

## 路由注册类（Handler & 路由注册）

| 函数                                      | 用途                  | 备注                                                              |
|-----------------------------------------|---------------------|-----------------------------------------------------------------|
| `http.Handle(pattern, handler)`         | 注册 `handler` 到默认路由器 | Handler 实现必须实现 `ServeHTTP`                                      |
| `http.HandleFunc(pattern, handlerFunc)` | 注册一个函数为 handler     | 最常用，handlerFunc 的签名是 `(w http.ResponseWriter, r *http.Request)` |

## 发起 HTTP 请求（Client Side Functions）

| 函数                                  | 用途         | 本质                                                |
|-------------------------------------|------------|---------------------------------------------------|
| `http.Get(url)`                     | 简单发 GET 请求 | 底层调用 `http.DefaultClient.Get`                     |
| `http.Post(url, contentType, body)` | 发 POST 请求  | 适合简单用法                                            |
| `http.PostForm(url, data)`          | 表单 POST    | `Content-Type: application/x-www-form-urlencoded` |
| `http.Head(url)`                    | 发 HEAD 请求  | 只获取响应头，不要 body                                    |

## 构建 Response / Request 的辅助函数

| 函数                                                | 用途                   |
|---------------------------------------------------|----------------------|
| `http.Error(w, msg, code)`                        | 向客户端返回错误响应           |
| `http.NotFound(w, r)`                             | 返回 404               |
| `http.Redirect(w, r, url, code)`                  | 发重定向                 |
| `http.ServeFile(w, r, filename)`                  | 直接返回静态文件             |
| `http.ServeContent(w, r, name, modtime, content)` | 比 ServeFile 更灵活，可处理流 |

## 总结思维导图

```
net/http 顶层函数
├── 1. 启动服务器
│   ├── ListenAndServe
│   ├── ListenAndServeTLS
│   └── Serve
│
├── 2. 注册路由
│   ├── Handle
│   └── HandleFunc
│
├── 3. 发起客户端请求
│   ├── Get / Post / PostForm / Head
│   └── 推荐使用自定义 http.Client
│
└── 4. 构建响应 / 工具函数
    ├── Error / NotFound / Redirect
    ├── ServeFile / ServeContent
    └── 这些函数内部都是围绕 ResponseWriter 写数据
