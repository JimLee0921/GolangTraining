# `bufio` 包

`bufio`（buffered I/O 缓冲输入输出）包实现了缓冲 I/O。它包装了一个 `io.Reader` 或 `io.Writer` 对象，并创建了另一个对象（Reader
或
Writer），该对象也实现了该接口，但提供了缓冲功能以及一些文本 I/O 的帮助。

它是 os 和 io 之间的中间层，让文件读写、标准输入等操作更高效、更方便。

`bufio` = `Buffered I/O`
全称就是 “带缓冲的输入输出”。

Go 中的文件和网络操作本质上是通过 os.File、net.Conn 等对象进行的，
但它们的读写是“裸 I/O”，一次系统调用可能性能较低。

`bufio` 在它们外面包一层 缓冲层（buffer）：

- 减少系统调用次数，提高性能
- 提供更方便的按行 / 按词 / 按分隔符读取
- 支持带缓冲的写入（只在 Flush 时真正写出）

## 核心类型（结构体）

输入（Reader 系列）、输出（Writer 系列）、输入+输出（ReadWriter / Scanner）

| 模块             | 类型                 | 作用                 | 典型用途        | 是否需要 Flush |
|----------------|--------------------|--------------------|-------------|------------|
| **Reader**     | `bufio.Reader`     | 带缓冲的读取器            | 文件读、网络流读取   | 否          |
| **Writer**     | `bufio.Writer`     | 带缓冲的写入器            | 文件写、日志、网络输出 | 是          |
| **ReadWriter** | `bufio.ReadWriter` | Reader + Writer 组合 | 网络通信、双向流    | 是          |
| **Scanner**    | `bufio.Scanner`    | 高层封装扫描器            | 逐行或逐词读取文本   | 否          |

```
Reader  -> 带缓冲的读取器
Writer  -> 带缓冲的写入器
ReadWriter -> Reader + Writer 组合
Scanner -> 高层封装的读取器（自动分词/分行）
```

### `bufio.Reader`

带缓冲的输入读取器，Reader 为 io.Reader 对象实现了缓冲。
通过调用NewReader或NewReaderSize可以创建一个新的 Reader ；或者，也可以在对
Reader 调用 `[Reset]` 后使用其零值。

**定义创建**

`bufio.NewReader(rd io.Reader) *Reader` 和  `NewReaderSize(rd io.Reader, size int) *Reader` 都可用于创建

| 创建方式   | 函数                                           | 缓冲区大小         | 说明        |
|--------|----------------------------------------------|---------------|-----------|
| 默认缓冲区  | `bufio.NewReader(r io.Reader)`               | 4096 字节（4 KB） | 直接使用默认大小  |
| 自定义缓冲区 | `bufio.NewReaderSize(r io.Reader, size int)` | 自定义           | 可根据场景优化性能 |

```
r := bufio.NewReader(io.Reader)
r2 := bufio.NewReaderSize(file, 16*1024)
```

返回一个缓冲区具有默认（4KB） 或指定缓冲区大小的新`Reader` ，它在普通 io.Reader 外包一层缓冲区，让读取更快、更灵活

- 读取速度快（减少系统调用）
- 支持按字节、按行、按分隔符读取
- 支持预读（Peek）、回退（Unread）
- NewReaderSize 返回一个新的Reader ，其缓冲区至少具有指定的大小。如果参数 io.Reader 已经是一个具有足够大大小的Reader，则返回底层Reader

### `bufio.Writer`

带缓冲的输出写入器，Writer 实现了io.Writer对象的缓冲功能。
如果写入Writer时发生错误，则不会再接受任何数据，并且所有后续写入操作以及Writer.Flush都会返回错误。
所有数据写入完成后，客户端应调用 Writer.Flush 方法，以确保所有数据都已转发至底层io.Writer。

**定义创建**

| 创建方式   | 函数                                           | 缓冲区大小        | 说明       |
|--------|----------------------------------------------|--------------|----------|
| 默认缓冲区  | `bufio.NewWriter(w io.Writer)`               | 4096 字节（4KB） | 默认缓冲     |
| 自定义缓冲区 | `bufio.NewWriterSize(w io.Writer, size int)` | 指定的大小        | 可优化性能或延迟 |

```
// 默认 4KB 缓冲
w1 := bufio.NewWriter(file)
// 自定义 16KB 缓冲
w2 := bufio.NewWriterSize(file, 16*1024)
```

- 数据先写入内存缓冲区
- 当缓冲区满时或调用 Flush() 时
- 才真正写到底层 io.Writer（比如文件或 socket）

| 优点           | 说明                   |
|--------------|----------------------|
| 性能提升         | 减少底层写入次数（例如每次写磁盘都很慢） |
| 小块数据聚合       | 多次 Write 变成一次大块写入    |
| 适合日志、文件、网络输出 |                      |

### `bufio.ReadWriter`

可以同时用于读写，常见于网络连接（如 socket 通信）。
**定义创建**

```
rw := bufio.NewReadWriter(reader, writer)

// 内部包含
type ReadWriter struct {
    *Reader
    *Writer
}
```

### `bufio.Scanner`

高层封装的输入扫描器

**定义创建**

```
scanner := bufio.NewScanner(io.Reader)
```

- 比 Reader 更轻量、易用
- 自动分割输入（默认按行）
- 每次调用 Scan() 读取一段内容
- 用 Text() 取出字符串

### 缓冲区大小设置

| 场景          | 推荐大小       | 说明      |
|-------------|------------|---------|
| 一般文件读取      | 4KB（默认）    | 已足够     |
| 读取超大文本文件    | 8KB ~ 64KB | 提高读取吞吐量 |
| 网络流、管道      | 1KB ~ 8KB  | 减少延迟    |
| 特殊场景（如大行读取） | ≥ 行长度      | 防止行被拆断  |

### Reader 和 Writer 对称关系

| 类型             | 默认函数          | 指定大小函数            | 默认缓冲大小 | 是否需要 Flush |
|----------------|---------------|-------------------|--------|------------|
| `bufio.Reader` | `NewReader()` | `NewReaderSize()` | 4KB    | 否          |
| `bufio.Writer` | `NewWriter()` | `NewWriterSize()` | 4KB    | 是 （必须手动调用） |
