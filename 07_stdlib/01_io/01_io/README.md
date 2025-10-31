# IO 包

> 官方文档：https://pkg.go.dev/io


io 是 Go 语言 I/O 模型的最底层抽象层。核心思想是：

- 只要能实现 Read 或 Write 方法，就可以参与所有输入输出操作
- 读文件、读网络连接、读内存、读压缩包、读 HTTP 请求体等全部都是 `io.Reader`
- 写文件、写网络、写日志、写缓冲、写加密流等全部都是 `io.Writer`

io 包提供了 I/O 原语的基本接口。主要作用是将现有的 I/O 原语实现（例如 os 包中的实现）包装成抽象功能的共享公共接口，并添加一些其他相关的原语。
这些接口和原语用各种实现包装了较低级别的操作，大部分不推荐直接使用。

### 核心理念

Go 的 `I/O` 是通过接口抽象的

```
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

所有实现了这些接口的类型（文件、缓冲区、socket 等）都可以互换使用

- `I/O` 是流式的
- `Read / Write` 不一定一次性完成，需要循环调用
- 错误值 io.EOF 表示流结束

### 四大基础接口

| 接口名         | 功能     | 常见实现                                                    |
|-------------|--------|---------------------------------------------------------|
| `io.Reader` | 读数据    | `os.File`, `bytes.Buffer`, `strings.Reader`, `net.Conn` |
| `io.Writer` | 写数据    | `os.File`, `bytes.Buffer`, `net.Conn`, `bufio.Writer`   |
| `io.Closer` | 关闭资源   | `os.File`, `net.Conn`                                   |
| `io.Seeker` | 定位读写位置 | 文件、内存缓冲区                                                |

> 注意：io.Copy 能工作，就是因为两个对象分别实现了 Reader 和 Writer

### I/O 包主要组成结构

| 分类             | 内容                                                               | 作用               |
|----------------|------------------------------------------------------------------|------------------|
| **接口定义**       | `Reader`, `Writer`, `Closer`, `Seeker`, `ReaderFrom`, `WriterTo` | 统一 I/O 行为标准      |
| **常用函数**       | `Copy`, `CopyN`, `ReadAll`, `ReadFull`, `WriteString`            | 快捷工具函数           |
| **Reader 组合器** | `LimitReader`, `MultiReader`, `TeeReader`                        | 构建复合 Reader      |
| **Writer 组合器** | `MultiWriter`, `Pipe`                                            | 构建复合 Writer 或流管道 |
| **常量与错误**      | `io.EOF`, `io.ErrShortWrite`, `io.ErrUnexpectedEOF`              | 标准错误值            |
