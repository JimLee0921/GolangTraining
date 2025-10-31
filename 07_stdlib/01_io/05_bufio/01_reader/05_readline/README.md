## ReadLine 按行读取

可以看作是 ReadSlice('\n') 的增强版，在保持高性能的同时，解决了 ReadSlice() 缓冲区太小时不会直接报错，而是自动拼接多段返回完整一行的问题

### 方法定义

```
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
```

返回值：

- line：字节切片，当前读取到的行数据（不包含换行符）
- isPrefix：是否为行的前半部分（表示行太长还没读完）
- err：错误对象

### 核心功能

逐行读取输入（以 `\n` 为行结束），自动拼接过长行，
`ReadLine()` 能保证最终能完整拿到一行数据，只是可能需要多次拼接。

区别于 ReadSlice('\n')：

- ReadSlice()：缓冲区不够时直接报错 ErrBufferFull
- ReadLine()：会返回 `isPrefix = true`，然后等待继续调用读取剩余部分

### 返回值逻辑

| 状态       | `line`       | `isPrefix` | `err`    |
|----------|--------------|------------|----------|
| 读到一行完整数据 | 当前行（不含 `\n`） | `false`    | `nil`    |
| 行太长，分段返回 | 当前部分         | `true`     | `nil`    |
| 到达文件末尾   | 剩余数据         | `false`    | `io.EOF` |
| 其他错误     | -            | -          | 非空       |
