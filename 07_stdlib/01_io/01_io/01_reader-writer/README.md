# IO.Reader & IO.Writer

`io.Reader` 和 `io.Writer` 是 `Go I/O` 的最核心接口，分别是包装 Read 和 Write 方法的接口

## 接口定义

```
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

```

### Read

从数据源读入字节到缓冲区

**参数**

- `p []byte`：缓存区，从数据源中读取字节到 p

**返回值**

- n: 实际读到的字节数（可能 < len(p)）
- err: 错误或结束标志。如果返回 io.EOF 表示读取完毕（End Of File）

**官方文档解释**

```text
Read 最多读取 len(p) 个字节到 p 中，返回读取的字节数 (0 <= n <= len(p)) 以及遇到的任何错误。
即使 Read 返回 n < len(p)，它也可能在调用期间将整个 p 用作临时空间。
如果存在可用数据但未达到 len(p) 个字节，Read 通常会返回可用数据，而不是等待更多数据。

当 Read 在成功读取 n > 0 个字节后遇到错误或文件结束条件时，它会返回读取的字节数。它可能在同一个调用中返回（非 nil）错误，也可能在后续调用中返回错误（且 n == 0）。
这种一般情况的一个例子是，Reader 在输入流末尾返回非零字节数，可能返回 err == EOF 或 err == nil。下一个 Read 应该返回 0，即 EOF。

调用者应始终先处理返回的 n > 0 个字节，然后再考虑错误 err。这样做可以正确处理读取一些字节后发生的 I/O 错误以及两种允许的 EOF 行为。
如果 len(p) == 0，Read 应该始终返回 n == 0。如果已知某些错误条件（例如 EOF），它可能会返回非零错误。

不鼓励 Read 的实现返回零字节数和 nil 错误，除非 len(p) == 0。调用者应该将返回 0 和 nil 视为没有发生任何事情；特别是它不表示 EOF。
```

### Writer

Write 向目标写入 p 中的数据，并返回实际写入的字节数 n 与错误 err

**参数**

- `p []byte`：要写入的数据（通常是切片、字符串转换的字节数组）

**返回值**

- `n int`：实际写入的字节数（可能 < len(p)）
- `err error`：如果出错，返回非 nil，否则为 nil

**官方文档解释**

```text
Writer 是包装基本 Write 方法的接口。

Write 将 p 中的 len(p) 个字节写入底层数据流。
它返回从 p 写入的字节数（0 <= n <= len(p)）以及导致写入提前停止的任何错误。
如果返回 n < len(p)，则 Write 必须返回非零错误。Write 不得修改切片数据，即使是临时修改。
```

## 工作原理

`Reader + Writer` 流式传输：在 Go 中，I/O 是流式的（streaming），意味着数据是分批读出、分批写入的，而不是一次性全部处理

```
buf := make([]byte, 4)
for {
    // 从 Reader 读取数据
    n, err := r.Read(buf)
    if n > 0 {
        // 写入 Writer（注意只写有效部分）
        if _, wErr := w.Write(buf[:n]); wErr != nil {
            fmt.Println("写入错误:", wErr)
            break
        }
        fmt.Printf("写入 %d 字节: %s\n", n, buf[:n])
    }

    if err == io.EOF {
        // 已经读完所有数据
        break
    }
    if err != nil {
        fmt.Println("读取错误:", err)
        break
    }
}
```

**执行流程**

1. 创建一个缓冲区 buf（例如 4 字节）
2. 不断从 `r.Read(buf)` 中读取下一段数据
3. 每次读到的数据长度 n 可能不同
4. 将 `buf[:n]` 写入 `w.Write()`
5. 当 `err == io.EOF` 时表示数据读完（不是错误）
6. 其它错误（网络断开、磁盘错误）才是异常终止。

**特点**

| 特性                | 说明                   |
|-------------------|----------------------|
| **流式读写**          | 数据分段读写，不会一次性加载到内存    |
| **部分读/写**         | 每次读到的字节数 `n` 可能小于缓冲区 |
| **`io.EOF` 不是错误** | 表示数据结束而非失败           |
| **双向处理**          | 每次读取后立即写入，适合管道、网络转发  |
| **健壮性高**          | 即使发生部分读写，也能安全继续循环    |

## 常见实现类型

| 分类        | 包                                | Reader 实现                                        | Writer 实现                         | 典型用途               |
|-----------|----------------------------------|--------------------------------------------------|-----------------------------------|--------------------|
| **内存**    | `strings`, `bytes`               | `strings.Reader`, `bytes.Reader`, `bytes.Buffer` | `bytes.Buffer`, `strings.Builder` | 读/写内存数据            |
| **文件**    | `os`                             | `os.File`                                        | `os.File`                         | 文件读写               |
| **缓冲**    | `bufio`                          | `bufio.Reader`                                   | `bufio.Writer`                    | 提高读写效率             |
| **网络**    | `net`                            | `net.Conn`                                       | `net.Conn`                        | 网络通信（TCP/UDP）      |
| **压缩**    | `compress/gzip`, `zlib`, `flate` | `gzip.Reader`                                    | `gzip.Writer`                     | 压缩/解压流             |
| **加密编码**  | `encoding/base64`, `hex`         | `base64.NewDecoder`                              | `base64.NewEncoder`               | 数据编码/解码            |
| **归档**    | `archive/zip`, `tar`             | `zip.File`, `tar.Reader`                         | `zip.Writer`, `tar.Writer`        | 打包/解包文件            |
| **对象序列化** | `encoding/json`                  | `json.Decoder`                                   | `json.Encoder`                    | 结构体 ↔ 文本/二进制       |
| **管道**    | `io`                             | `io.PipeReader`                                  | `io.PipeWriter`                   | 协程间数据传输            |
| **多路复用**  | `io`                             | `io.MultiReader`                                 | `io.MultiWriter`                  | 拼接多个 Reader/Writer |
