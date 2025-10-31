# `bufio.Reader`

`bufio.Reader`带缓冲的输入读取器，Reader 为 io.Reader 对象实现了缓冲。
通过调用 `NewReader` 或 `NewReaderSize` 可以创建一个新的 `Reader`，或者，也可以在对
Reader 调用 `[Reset]` 后使用其零值。

`bufio.Reader` 不会每次都直接读取底层（例如`os.File`），而是：

1. 一次性从底层读一大块数据进内存缓冲区
2. 之后用户的多次 `Read()`、`ReadByte()`、`ReadString()` 等操作，都从这块内存中取
3. 缓冲区读完后再重新填充

## 定义创建

`bufio.NewReader(rd io.Reader) *Reader` 和  `NewReaderSize(rd io.Reader, size int) *Reader` 都可用于创建 `Reader` 对象

| 创建方式   | 函数                                           | 缓冲区大小             | 说明                             |
|--------|----------------------------------------------|-------------------|--------------------------------|
| 默认缓冲区  | `bufio.NewReader(r io.Reader)`               | 默认s 4096 字节（4 KB） | NewReaderSize的包装函数直接使用默认大小4096 |
| 自定义缓冲区 | `bufio.NewReaderSize(r io.Reader, size int)` | 自定义               | 可根据场景优化性能                      |

```
r := bufio.NewReader(io.Reader)
r2 := bufio.NewReaderSize(file, 16*1024)
```

返回一个缓冲区具有默认（4KB） 或指定缓冲区大小的新`Reader` ，它在普通 io.Reader 外包一层缓冲区，让读取更快、更灵活

- 读取速度快（减少系统调用）
- 支持按字节、按行、按分隔符读取
- 支持预读（Peek）、回退（Unread）
- 如果参数 io.Reader 已经是一个具有足够大大小的Reader，则返回底层Reader

## 补充

1. 字符串读取时可以使用 `strings.NewReader` 方法将 string 包装成一个实现了 `io.Reader` 接口的对象然后传给
   `bufio.NewReader` 进行使用

2. 在 Go 的 `bufio.Reader.ReadString(delim)`（或 `ReadBytes`、`ReadSlice`）里，delim 参数必须是一个单个字节（byte 类型），
   而不是一个字符或者任意符号。英文的逗号`,`可以，但是中文的逗号`，`不行
