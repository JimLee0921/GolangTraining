## WriteRune 写入字符

WriteByte() 的「Unicode 加强版」，可以写入任何字符（包括中文、emoji、特殊符号等）。

### 方法签名

```
func (b *Writer) WriteRune(r rune) (size int, err error)
```

写入单个 Unicode 代码点，返回写入的字节数和任何错误。

### 功能说明

WriteRune() 用于写入一个 Unicode 字符 到缓冲区。
它会自动将 rune（即 int32 类型）编码为 UTF-8 字节序列，
然后写入 `bufio.Writer` 的内部缓冲区。


- `WriteByte()`：写入单个字节（仅支持 ASCII）
- `WriteRune()`：写入单个字符（支持所有 Unicode）