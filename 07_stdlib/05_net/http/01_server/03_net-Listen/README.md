## 先监听再交给 Server.Serve

可以先使用 `net.Listen` 进行端口监听再交给 Server 进行接管。

* 手动管理监听端口
* 自定义监听方式（重用端口、SO_REUSEPORT、优雅重启）
* 限制最大连接数
* 接入自定义 `net.Listener`（比如 TLS、Unix Socket）

```
ln, _ := net.Listen("tcp", ":8080")
srv := &http.Server{}
srv.Serve(ln)
```

### 启动服务

在创建时不指定 Addr 端口，在启动时使用 `Serve()` 而不是 `ListenAndServe()`

因为 `Serve()` 接收的是：

```
func (srv *Server) Serve(l net.Listener) error
```