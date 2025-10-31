## writer 辅助方法

### Buffered

```
func (b *Writer) Buffered() int
```

返回当前缓冲区中（已经写入但还没被 flush 到底层 writer）的字节数，常用于判断是否有数据未写入

> 如果缓冲区大小小于字节数则返回 Flush 后缓冲区大小

**用途**

| 用途   | 说明                          |
|------|-----------------------------|
| 性能监控 | 判断是否有未写出的数据                 |
| 调试   | 检查是否忘记调用 `Flush()`          |
| 写入优化 | 配合 `Available()` 判断何时写出大块数据 |

### Available

```
func (b *Writer) Available() int
```

返回当前缓冲区中剩余可写入的字节数（即还没被占用的空间）。也就是还可以再写多少字节才会填满缓冲区

### Size

返回 `bufio.Writer` 缓冲区的总大小（固定值）。默认是 4096 字节，也可以在 `NewWriterSize()` 时手动指定。

- Available() + Buffered() = Size()缓冲区总大小（即 Writer 的容量）
- Writer 默认容量是 4096 字节（4KB），除非用 NewWriterSize() 手动指定

### Reset

```
func (b *Writer) Reset(w io.Writer)
```

复用缓冲区，把当前 Writer 绑定到新的底层 `io.Writer`，但保留原有的缓冲区内存。
这可以避免重复分配内存，提高性能。

### Flush

Flush 将任何缓冲数据写入底层 `io.Writer`。
大部分 write 方法会在缓冲区满时自动调用，但是还需要最后手动调用一次，否则可能会有剩余内容未写入。

### ReadFrom

```
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
```

从一个 `io.Reader` 连续读取数据，并直接写入到当前 `bufio.Writer` 的缓冲区，
自动处理填充和 `Flush()`，直到读取完成（遇到 EOF 或错误）。
就是把一个 Reader 的数据写进 Writer 里。

> 较少手动调用 io.Copy() 内部就会优先调用 ReadFrom() 来加速


Go 的 `io.Copy(dst, src)` 会自动检测：

- 如果 dst 实现了 `io.ReaderFrom`（即拥有 `ReadFrom()` 方法），就调用它；

- 否则，执行普通的 `Read / Write` 循环。

因此：`io.Copy(bufio.NewWriter(f), src)`实际上就是在内部使用 `ReadFrom()`，自动进行高效传输。
