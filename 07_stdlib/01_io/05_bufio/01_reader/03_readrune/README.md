## ReadRune 按字符读取

- ReadByte() 是逐字节读
- ReadRune() 是逐字符读（自动处理 UTF-8 多字节）

### 方法定义

```
func (b *Reader) ReadRune() (r rune, size int, err error)
```

返回值

- r：rune类型，读到的字符（Unicode 码点）
- size：该字符所占的字节数
- err：错误信息（如 `io.EOF` 用于结束信号）

### 工作原理

1. 优先从缓冲区取数据
2. 根据 UTF-8 规则自动解码：

    - 若第一个字节以 `0xxxxxxx` 开头 -> 单字节字符
    - 若以 `110xxxxx`、`1110xxxx` 等开头 -> 组合多字节

3. 自动返回完整字符（rune）
4. 更新缓冲区指针
5. 如果遇到 EOF -> 返回 `io.EOF` 作为结束读取信号

## UnreadRune 退回字符

### 方法定义

```
func (b *Reader) UnreadRune() error
```

用于撤销上一次 `ReadRune()` 操作，将刚读出的字符回退到缓冲区中

### 使用限制

| 限制                           | 说明                                   |
|------------------------------|--------------------------------------|
| **只能回退一次**                   | 连续调用两次会出错：`bufio: can't unread rune` |
| **只能回退最近一次成功的 `ReadRune()`** | 不能在 `ReadByte()` 或 `Read()` 后调用      |
| **不能跨越缓冲区边界回退**              | 一般不影响普通使用，但要知道回退只针对最近一次字符            |
