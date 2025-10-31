## `io.Closer`

| 接口                   | 能力        | 场景              |
|----------------------|-----------|-----------------|
| `io.Closer`          | 关闭资源      | 所有需要释放资源的对象     |
| `io.ReadCloser`      | 可读 + 可关闭  | 文件读取 / HTTP 响应体 |
| `io.WriteCloser`     | 可写 + 可关闭  | 压缩器 / 网络发送流     |
| `io.ReadWriteCloser` | 读 + 写 + 关 | TCP 连接 / 文件     |

### 定义

```
type Closer interface {
    Close() error
}
type ReadCloser interface {
    Reader
    Closer
}
type WriteCloser interface {
    Writer
    Closer
}
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

出现需要 Close 的对象，推荐用 `defer xxx.Close()`，不然可能会造成资源泄漏。

尤其是：

- HTTP Body 不关 -> 连接池耗尽
- DB 不关 -> 连接占满
- Writer 不关 -> 数据不落盘
- File 不关 -> FD 泄漏导致 too many open files