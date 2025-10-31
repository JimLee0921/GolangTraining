## Read 基础读取

最底层的读法，直接往 p 填数据，返回读入 p 的字节数

### 定义

```
func (b *Reader) Read(p []byte) (n int, err error)
```

参数：

- `p []byte`：提供的字节切片，用来存放读取到的数据

返回值：

- n int：实际读入的字节数
- err error：错误对象。常见的有：
    - nil：读取成功
    - io.EOF：到达输入流的末尾（End Of File）

### 行为说明

1. 优先从缓冲区读取
    - 如果内部缓冲区（b.buf）里已经有数据，就直接从那里复制到 p
    - 当缓冲区数据不足以填满 p 时，才会继续从底层 io.Reader 读取更多数据

2. 自动补充缓冲区
    - 当缓冲区用尽时，`bufio.Reader` 会调用底层的 Read 方法（比如 os.File.Read、strings.Reader.Read 等）填充缓冲区

3. 返回值规则
    - 如果读取成功（即使读到的字节数小于 len(p)），返回 (n, nil)
    - 如果读到部分数据后遇到 EOF，则返回 (n, io.EOF)，此时应当先处理这部分数据，再判断是否结束

### 注意事项

1. 不要忽略返回的 n
    - 即使 `err != nil`，n 也可能仍然 > 0（读到部分数据）
    - 只要还读到数据，这次就返回 `err = nil`，把 EOF 留到下一次
2. 使用 `io.EOF` 作为正常结束信号 `if err==io.EOF {break}`
3. 不要混用 `reader` 和 `bufio`
    - 一旦用 `bufio.Reader` 包装了底层 reader，不要再直接操作底层 reader，否则会打乱缓冲区状态
