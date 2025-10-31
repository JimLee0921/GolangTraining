## 常用 io 包拓展接口

## 重点接口（Reader/Writer 的扩展）

| 接口                                 | 功能               | 出现的场景               |
|------------------------------------|------------------|---------------------|
| `io.Closer`                        | 关闭资源             | 文件、网络连接、HTTP Body   |
| `io.Seeker`                        | 定位 read/write 位置 | 文件、内存流、实现快进/回退      |
| `io.ReaderAt` / `io.WriterAt`      | **随机读写，不依赖当前位置** | 数据库引擎、存储引擎、文件 mmap  |
| `io.ReadWriter`                    | 同时读又写            | TCP 连接、bytes.Buffer |
| `io.ReadCloser` / `io.WriteCloser` | “读/写 + 关闭”组合接口   | HTTP、文件             |

**os.File** 同时实现了这些接口：

```
f *os.File:
io.Reader
io.Writer
io.Closer
io.Seeker
io.ReaderAt
io.WriterAt
```

所以，可以把文件当作：

* 一个流 (`Reader`)
* 一个输出目标 (`Writer`)
* 一个可关闭资源 (`Closer`)
* 一个可移动光标的介质 (`Seeker`)
* 一个支持随机访问的存储 (`ReaderAt`/`WriterAt`)