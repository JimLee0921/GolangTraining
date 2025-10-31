## `io.Copy` / `io.CopyN`

实现数据流转最核心的工具，用来在不同的 Reader / Writer 之间传输数据。

### io.Copy

```
func Copy(dst Writer, src Reader) (written int64, err error)
```

- 从 src 连续读取数据
- 写入到 dst
- 直到 src 读完（EOF） 或发生错误

> 本质就是：文件 / 网络 / 内存之间搬运数据

### io.CopyN

```
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
```

从 src 复制 固定 n 字节 到 dst，不会读到 EOF，而是严格复制指定长度

| 函数                      | 何时使用      | 说明        |
|-------------------------|-----------|-----------|
| `io.Copy(dst, src)`     | 复制整个流     | 直到 EOF 停止 |
| `io.CopyN(dst, src, n)` | 只复制前 n 字节 | 常用于文件头、分块 |
