## ReadBytes 字节读取

### 方法定义

```
func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
```

参数：

- `delim byte`：分隔符（单个字节），读取会一直持续到遇到它为止

返回值

- line：字节切片，包含分隔符 的完整切片（已安全拷贝）
- err：错误信息，可能是 nil 或 io.EOF（读到结尾）

### 功能说明

继承了 ReadSlice() 的分隔符读取逻辑，但解决了后者易踩坑的问题（缓冲区太小、数据被覆盖等）。
从缓冲区和底层 reader 读取，直到遇到指定分隔符 delim 为止，返回包含该分隔符的完整字节切片。

**特点**

- 自动扩展内存缓冲区
- 不会出现 ErrBufferFull
- 返回的新切片不会被覆盖
- 性能略低于 `ReadSlice()`，但更安全

> 与 ReadSlice() 不同，ReadBytes() 不会抛出 ErrBufferFull，
> 而是自动把多次结果自动拼接成一个完整 []byte 返回。

