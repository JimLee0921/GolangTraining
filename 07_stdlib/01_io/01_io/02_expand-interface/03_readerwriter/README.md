## `io.RedeWriter`

`io.ReadWriter` 是 `Go I/O` 体系中最常见的组合接口之一，表示 既能读（Reader）又能写（Writer）的对象。
没有自己的新方法，就是简单把 Reader 和 Writer 组合在一起。

```
type ReadWriter interface {
    Reader
    Writer
}
```

等价于

```
type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
}
```

### 使用场景

| 场景      | 为什么使用 ReadWriter  |
|---------|-------------------|
| 网络连接    | TCP 连接既能读又能写      |
| 内存缓冲区   | 可以读写同一个内存流        |
| 自定义流处理器 | 上层只关心“能读能写”，不关心实现 |

只要一个对象同时实现 Read 和 Write，它就自动满足 io.ReadWriter。

典型例子：

- `*os.File`
- `net.Conn`（网络连接 socket）
- `bytes.Buffer`
- `bufio.ReadWriter`