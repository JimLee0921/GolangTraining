## 结构体+方法创建 Handler

只要类型实现了 ServeHTTP 方法，它就是一个 Handler。

```
type HelloHandler struct{} // 一个空结构体，表示 Handler

// 实现 ServeHttp 方法
func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from struct handler!")
}

// 启动服务时传入一个结构体的值
http.ListenAndServe(":8080", HelloHandler{})
```

### 与函数 Handler 对比

| 方式                       |   是否可带状态   | 常见场景              |
|--------------------------|:----------:|-------------------|
| `http.HandlerFunc(func)` | 需要闭包才能持状态  | 写简单路由处理逻辑         |
| 结构体 + ServeHTTP          | 天然支持业务依赖注入 | 真实业务服务、MVC 模块、控制器 |

结构体对比简单函数可以 带状态（保存变量、依赖、配置、数据库连接等）。

```
type UserHandler struct {
	DBName string
}

func (h UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Using database: %s", h.DBName)
}

func main() {
	handler := UserHandler{DBName: "users.db"}
	http.ListenAndServe(":8080", handler)
}
```

访问 / 会输出：`Using database: users.db`

> 这意味着 Handler 可以和业务上下文绑定，这就是可维护的工程代码风格

### 推荐用指针接收者

推荐创建 ServeHTTP 方法使用 指针接收者：

- 避免复制结构体（尤其当结构体很大）
- 允许内部修改状态（计数器、缓存）

```
type CounterHandler struct {
	Count int
}

func (h *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Count++
	fmt.Fprintf(w, "Visited %d times\n", h.Count)
}

func main() {
	http.ListenAndServe(":8080", &CounterHandler{})
}
```

> 访问页面会看到计数 + 1，因为是 同一个 Handler 实例
>
> 用结构体 + 指针接收者 可以做有状态服务。

### `http.Handle` 函数

http.Handle 是包级函数，向默认的 DefaultServeMux 注册路由

```
func Handle(pattern string, handler Handler) {
    DefaultServeMux.Handle(pattern, handler)
}
```

> 所以默认情况下是在向 DefaultServeMux 注册路由