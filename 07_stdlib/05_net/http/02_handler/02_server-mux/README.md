# http.ServeMux

ServeMux 是标准库自带的路由器（HTTP 请求复用器），实现了 `http.Handler`，所以本质上也是一个 Handler。
它会将每个传入请求的 URL 与已注册的模式列表进行匹配，并调用与 URL 最匹配的模式对应的处理程序。

```
type ServeMux struct {}
```

ServeMux = 路由分发器

- 根据 URL Path，为请求选择对应的 Handler
- Server 启动后，所有请求都会交给 `mux.ServeHTTP`，再由 mux 决定转给哪个 Handler

> ServeMux 就像地图：URL -> Handler 的映射表

### 创建 Mux

最常见的就是使用 `http.NewServeMux()` 进行创建：

```
mux := http.NewServeMux()
```

还有几种等价/相关方式：

1. 默认路由器：`http.DefaultServeMux`，不手动创建，直接用包级函数注册到默认 mux
    ```
    http.HandleFunc("/hello", hello)            // 注册到 DefaultServeMux
    http.ListenAndServe(":8080", nil)           // Handler=nil → 用 DefaultServeMux
    ```

2. 零值可用：&http.ServeMux{} 或 变量零值。ServeMux 的零值就是空路由器，可以直接使用
    ```
    mux := &http.ServeMux{}                     // 等价于 NewServeMux()
    // 或 var mux http.ServeMux
    mux.HandleFunc("/hello", hello)
    http.ListenAndServe(":8080", mux)
    ```

## 注册路由

可以使用 `Handle` 或者 `HandleFunc` 进行路由注册

```
func (mux *ServeMux) Handle(pattern string, handler Handler)
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

- Handle：接收实现了 ServeHTTP 的对象
- HandleFunc：接收一个普通函数，内部会自动包装成 HandlerFunc

## 路由匹配规则

`http.ServeMux` 的匹配规则如下

### 最长前缀匹配（Longest Match Wins）

* 先找更长、更具体的 pattern，找不到再退而求其次
* 顺序与注册先后无关

```
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "root") })
mux.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "user group") })
mux.HandleFunc("/user/profile", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "profile") })
// /user/profile → 命中 "/user/profile"
// /user/42      → 命中 "/user/"
// /anything     → 命中 "/"
```

### 末尾斜杠的语义：目录 vs 具体路径

* 以 `/` 结尾的 pattern**（如 `/user/`）表示这棵子树下的所有路径
* 不以 `/` 结尾的 pattern（如 `/user`）只匹配完全相等的路径

```
mux.HandleFunc("/user/", h) // 匹配 /user/、/user/42、/user/a/b...
mux.HandleFunc("/user",  h) // 只匹配 /user，不匹配 /user/ 或 /user/xxx
```

### 自动重定向（PathSlash 规范化）

* 注册了 `/user/`，当访问 `/user`（少了斜杠）-> ServeMux 自动 301 重定向到 `/user/`
* 反之不成立：注册 `/user` 时访问 `/user/` 不会跳回 `/user`，只会继续按最长匹配寻找（通常落到 `/`）

> 这就是为什么静态资源常用 `/static/`：
> 访问 `/static` 会被自动 301 到 `/static/`，再向下匹配子文件

### 路径清洗（CleanPath）与 301

ServeMux 会把非规范路径变成规范路径（必要时 301）：

* 多重斜杠：`//a///b` → `/a/b`
* `.` / `..`：`/a/./b/../c` → `/a/c`
* 确保以 `/` 开头；空路径等价 `/`

这能提升健壮性，但也会意外看到 301——是帮助规范 URL

### Host + Path 双维度匹配（少用但强大）

pattern 允许带主机名：`"api.example.com/v1/"`

* 当请求的 `r.Host` 命中该主机名时，才会匹配到这条规则
* Host 不区分大小写；Path 区分大小写

```
mux.HandleFunc("api.example.com/v1/", v1)
mux.HandleFunc("/v1/", fallback) // 其他主机走这里
```

### 参与匹配的只有 Host & Path（不含 Method / Query）

* HTTP 方法（GET/POST/PUT等）不参与匹配；要在 Handler 里自己分支
* 查询串（`?a=1`）不参与匹配；ServeMux 只看 `r.URL.Path`
* Path 是大小写敏感的：`/User` 不等于 `/user`

```
mux.HandleFunc("/items", func (w http.ResponseWriter, r *http.Request) {
switch r.Method {
case http.MethodGet: // ... 
case http.MethodPost: // ...
default:
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
})
```

### 注册规则与错误

* 同一个 pattern 只能注册一次（重复会 `panic`）
* 生产里建议集中构建路由树（启动期即暴露冲突）

## 组织路由的推荐方式

用路径前缀做分组：

```
func registerUser(mux *http.ServeMux) {
	mux.HandleFunc("/user/login", login)
	mux.HandleFunc("/user/logout", logout)
	mux.HandleFunc("/user/profile", profile)
}

func registerOrder(mux *http.ServeMux) {
	mux.HandleFunc("/order/create", create)
	mux.HandleFunc("/order/detail", detail)
}

func main() {
	mux := http.NewServeMux()
	registerUser(mux)
	registerOrder(mux)
	http.ListenAndServe(":8080", mux)
}
```

若要把一整棵子路由交给另一个处理器，用 `http.StripPrefix`：

```
api := http.NewServeMux()
api.HandleFunc("/v1/ping", ping)
api.HandleFunc("/v1/time", now)

root := http.NewServeMux()
root.Handle("/api/", http.StripPrefix("/api", api)) // /api/v1/* → api 的 /v1/*

http.ListenAndServe(":8080", root)
```

服务器收到请求 `/api/v1/ping` 时不想让 root 处理它
想把它转交给 api 路由器去处理，但是 api 路由器定义的路径是 `/v1/ping`，不带 /api ，
所以需要把路径前缀 `/api` 去掉之后再交给 api。

| 好处      | 为什么                                  |
|---------|--------------------------------------|
| 子模块路由独立 | `/v1/*` 不受 `/api` 前缀影响，模块更清晰         |
| 可组合     | 可以把 `/api` 路由树独立放到另一个包或 microservice |
| 支持版本路由  | `/api/v1`、`/api/v2` 可以清晰分离           |
