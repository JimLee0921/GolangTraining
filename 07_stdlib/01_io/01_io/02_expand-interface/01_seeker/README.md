## `id.Seeker`

`io.Seeker` 是 Go 中的一个接口，用于控制读写位置（文件指针）在可读写数据对象（比如文件）中的移动。
它让你能跳到文件中的任意位置读/写，而不仅仅是从头到尾顺序处理。

### 接口定义

```
type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}
```

- offset：偏移量
- whence：参考点
- 返回：移动后的位置

**`wherce` 三种取值**

| 常量                   | 含义     | 参考点    | 示例含义                                |
|----------------------|--------|--------|-------------------------------------|
| `io.SeekStart` (0)   | 从开头算   | 文件开头   | `Seek(10, SeekStart)` -> 跳到第 10 字节处 |
| `io.SeekCurrent` (1) | 从当前位置算 | 当前读写位置 | `Seek(-3, SeekCurrent)` -? 向前退 3 字节 |
| `io.SeekEnd` (2)     | 从文件末尾算 | 文件末尾   | `Seek(0, SeekEnd)` -> 跳到文件最后        |