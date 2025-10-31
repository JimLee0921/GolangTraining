## reader 其它辅助方法

| 方法                   | 作用             | 是否移动读指针 | 是否影响缓冲区内容 |
|----------------------|----------------|---------|-----------|
| `Buffered()`         | 查看缓冲区中还有多少未读字节 | 否       | 否         |
| `Discard(n)`         | 跳过（丢弃）n 个字节    | 是       | 是         |
| `Peek(n)`            | 偷看接下来 n 个字节    | 否       | 否         |
| `Reset(r io.Reader)` | 更换底层数据源        | 是（重置指针） | 是（清空缓存）   |
| `Size()`             | 返回缓冲区总容量（固定值）  | 否       | 否         |

### Buffered

```
func (b *Reader) Buffered() int
```

返回当前缓冲区中还未被读取的字节数。也就是还能在不触发底层 `I/O` 的情况下读取多少字节。

#### 示例

```
r := bufio.NewReaderSize(strings.NewReader("Hello, World!"), 8)
r.ReadByte() // 消费1个字节 'H'

fmt.Println("Buffered:", r.Buffered())
```

输出：

```
Buffered: 7
```

> 因为缓冲区初始容量 8，在读取 'H' 之后，还剩 7 个字节未读

#### 常见用途

- 判断当前缓冲区中是否还有数据
- 配合 `Peek()` 使用
- 在实现流解析器时查看剩余可读数据而不触发底层读取

### Discard

```
func (b *Reader) Discard(n int) (discarded int, err error)
```

跳过接下来 n 个字节（丢弃，不返回），返回丢弃的字节数。
如果 Discard 跳过的字节数少于 n，也会返回错误。
如果 `0 <= n <= b.Buffered()`，则 Discard 保证成功，无需从底层 `io.Reader` 读取。

#### 示例

```
r := bufio.NewReader(strings.NewReader("ABCDEFG"))
r.Discard(3)

b, _ := r.ReadByte()
fmt.Printf("跳过3个后读到: %c\n", b)
```

输出：

```
跳过3个后读到: D
```

> 跳过了 `A、B、C`下一次读取从 `D` 开始


> 注意：如果要丢弃的数据多于缓冲区，就会触发底层读；若中途遇到 `io.EOF`，会返回 `err=io.EOF` 和实际跳过的字节数

#### 常见用途

- 跳过固定头部：例如丢弃前 4 字节的 magic number
- 跳过无关字段： 如跳过分隔符或无意义填充
- 快速定位后续部分： 比手动 `Read()` 一次次更高效

### Peek

```
func (b *Reader) Peek(n int) ([]byte, error)
```

Peek 返回接下来的 n 个字节，但不推进读取器（指针）。这些字节在下一次读取调用时失效。
如有必要，Peek 会将更多字节读入缓冲区，以确保 n 个字节可用。
如果 Peek 返回的字节少于 n 个，它还会返回一个错误，解释读取时间短的原因。
如果 n 大于 b 的缓冲区大小，则 错误为 ErrBufferFull 。

调用 Peek 会阻止 `Reader.UnreadByte` 或 `Reader.UnreadRune` 调用成功，直到下一次读取操作为止。

#### 示例

```
r := bufio.NewReader(strings.NewReader("GoLang"))
b, _ := r.Peek(2)
fmt.Printf("Peek(2): %s\n", b)

ch, _ := r.ReadByte()
fmt.Printf("ReadByte(): %c\n", ch)
```

输出：

```
Peek(2): Go
ReadByte(): G
```

> Peek 不影响读取位置，所以之后的 `ReadByte()` 仍然从 'G' 开始

#### 注意事项

- Peek 返回的切片指向内部缓冲区
- 若之后调用了任何读操作，该切片可能被修改
- 若 Peek 读取的字节超过缓冲区容量，会返回 `bufio.ErrBufferFull`

### Reset

```
func (b *Reader) Reset(r io.Reader)
```

将现有的 bufio.Reader 重置为读取另一个数据源。
相当于：清空缓冲区 + 替换底层 Reader + 复用同一对象

#### 示例

```
r1 := bufio.NewReader(strings.NewReader("hello"))
buf := make([]byte, 5)
r1.Read(buf)
fmt.Println("第一次:", string(buf))

// 重置为另一个输入流
r1.Reset(strings.NewReader("world"))
r1.Read(buf)
fmt.Println("第二次:", string(buf))
```

输出：

```
第一次: hello
第二次: world
```

#### 使用场景

| 场景               | 说明                  |
|------------------|---------------------|
| 循环复用 `Reader` 对象 | 避免频繁分配新对象（节省内存与 GC） |
| 在同一函数内读取多个输入源    | 比如解析多个独立消息体         |
| 性能优化             | 对于高频解析器、网络框架特别重要    |

### Size

func (b *Reader) Size() int

返回缓冲区的总容量（即创建时指定的大小），对于 NewReader 就是默认的 4096

| 方法           | 含义               |
|--------------|------------------|
| `Size()`     | 缓冲区总容量（固定）       |
| `Buffered()` | 当前可用的未读字节数（动态变化） |
