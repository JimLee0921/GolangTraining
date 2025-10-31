# `bufio.ReadWriter`

可以同时用于读写（Reader + Writer 混合结构），用于同时处理输入与输出流的缓冲。常见于网络连接（如 socket 通信）

## 类型定义

```
type ReadWriter struct {
    *Reader
    *Writer
}
```

就是把一个 `*bufio.Reader` 和一个 `*bufio.Writer` 封装在一起

## 创建方式

```
func NewReadWriter(r *Reader, w *Writer) *ReadWriter
```

示例：

```
r := bufio.NewReader(os.Stdin)
w := bufio.NewWriter(os.Stdout)
rw := bufio.NewReadWriter(r, w)
```

- rw.Reader 提供所有读取方法（Read, ReadString, Peek, ...）
- rw.Writer 提供所有写入方法（Write, WriteString, Flush, ...）

> ReadWriter 可用让一个对象上同时拥有所有 Reader 和 Writer 的方法