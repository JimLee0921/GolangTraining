## ReadString 字符串读取

ReadBytes() 的字符串友好版。
功能完全一样，只是返回类型改成了 string，用起来更直观，非常适合处理文本内容（如日志、配置、输入行等）。

### 方法定义

```
func (b *Reader) ReadString(delim byte) (string, error)
```

参数：

- `delim byte`：分隔符（单个字节），通常是 '\n' 或 ','

返回值：

- string：包含分隔符的完整字符串
- error：错误信息

### 功能说明

从缓冲区中持续读取数据，直到遇到分隔符 delim 或到达 EOF。返回包含该分隔符的完整 字符串。

- `ReadBytes()` + `string()` 转换
- 自动扩容、自动拷贝
- 不会出现 ErrBufferFull
- 不会共享底层缓冲（返回的是新字符串）

