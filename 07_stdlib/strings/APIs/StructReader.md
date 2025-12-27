# `strings.Reader`

`strings.Reader` 是把一个内存中的 string 适配为可读的字节/字符流，从而接入 Go 的 `io.Reader` 生态。
如果现在有一个字符串需要交给一些只接受 `io.Reader`/ `io.Seeker` 等接口的库/函数，就需要用到 `strings.Reader`

```
type Reader struct {
	// contains filtered or unexported fields
}
```

本质上是读指针+源字符串，不修改字符串内容，只移动内部位置。
Reader 对象通过读取字符串来实现 `io.Reader`、`io.ReaderAt`、`io.ByteReader`、`io.ByteScanner`、`io.RuneReader`、
`io.RuneScanner`、`io.Seeker` 和 `io.WriterTo` 接口。Reader 的值为零时，其行为类似于读取空字符串的 Reader 对象。

## 构造函数

Reader 的起点，用一个 string 创建一个 Reader 对象，读指针从 0 开始

```
func NewReader(s string) *Reader
```

- 不复制 string 内容，时间复杂度为 O(1)
- Reader 只是持有 string 引用+位置（本质上返回的 Reader 中有一个 `s string` 数据源和 `i int64`
  当前读取位置还有其它一些辅助字段）
- 是轻量对象创建，可以频繁创建，不存在 Builder 那种声明周期风险

## 方法集

Reader 方法可以分为四种

1. 状态查询类：`Len`、`Size`
2. 顺序读取类：`Read`、`ReadByte`、`ReadRune`
3. 访问与定位类：`ReadAt`、`Seek`
4. 回退与输出类：`UnreadByte`、`UnreadRune`、`WriteTo`、`Reset`

### 状态读取类

还有多少可读 / 总大小是多少

#### Len

返回剩余未读的字节数，不是原始字符串长度，也不是 rune 数量，而是 `len(r.s) - len(r.i)`

```
func (r *Reader) Len() int
```

永远 >= 0 且会随着 Read/Seek 方法改变

#### Size

Reader 的全局视角，返回底层字符串的总字节数（固定不变的）为

```
func (r *Reader) Size() int64
```

### 顺序读取类

解决像读流一样往前读

#### `Read`

从当前读指针开始，把数据读入 b 并推进读指针，`strings.Reader` 最核心的方法，用于实现 `io.Read` 接口

```
func (r *Reader) Read(b []byte) (n int, err error)
```

- 从 `r.i` 开始读取，最多读 `len(b)` 个字节，实际读取量为 `min(len(b), r.Len())`
- 如果成功读到数据： ` err==nil `
- 如果读到结尾：`n == 0` 并且 `err == io.EOF`

> Read 永远按照字节工作，不关心 UTF-8/rune

#### `ReadByte`

字节级别顺序读取，从当前位置读取一个字节并把读指针前移1，对 `io.ByteReader` 接口的实现

```
func (r *Reader) ReadByte() (byte, error)
```

- 如果还有数据返回 `byte, nil` 并且 `r.i+1`
- 如果已经读到末尾返回 0 和 `io.EOF`

> 主要为了避免 `Read([]byte{1})` 这种低效写法，可以用于判断分隔符等

#### `ReadRune`

Unicode 语义读取，从当前位置按 UTF-8 解码读取一个 rune

```
func (r *Reader) ReadRune() (ch rune, size int, err error)
```

- 不是读一个字节，而是读 1-4 个字节解码成一个 rune
- `ch rune` 为读取到的 Unicode 字符
- `size int` 为该 rune 占用的字节数
- 正常 err 返回 nil，如果到了末尾返回 `io.EOF`，如果是非法 UTF-8 返回 `utf-8.RuneError`并且 size 为 1
- 同样会推进读指针：`r.i+=size`

### 访问与定位类

主要用于从指定位置读或跳到某个位置继续读

#### `ReadAt`

从字符串的 off 位置开始读数据到 b，但不改变当前读指针，实现 `io.ReadAt` 接口

```
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
```

- `off < 0` 会触发 panic，`off >= len(s)` 会返回 0 和 `io.EOF`
- 因为只读，所以是线程安全的
- 和 `Read` 相比不会推进读指针，也就是 `r.i` 不会改变
- 让 `Reader` 具备了随机访问文件的能力

#### `Seek`

移动当前读指针到指定位置并返回新的相对位置，实现了 `io.Seeker` 接口，会改变 `r.i` 也就是读指针位置

```
func (r *Reader) Seek(offset int64, whence int) (int64, error)
```

- 按照 whence 模式移动 offset 个字节
- whence 三种模式：
    - `io.SeekStart`：从字符串开头
    - `io.SeekCurrent`：从对当前位置
    - `io.SeekEnd`：从字符串末尾
- `len(s)<目标位置<0` 会触发 panic
- 成功更新 `r.i`，并返回新的绝对位置

### 回退与输出类

把刚读的东西退回去或高效把剩余内容写到 writer

#### `UnreadByte`

字节级别回退一步，撤销最近一次成功的 `ReadByte`，把读指针回退一个字节，实现的 `io.ByteScanner` 接口

```
func (r *Reader) UnreadByte() error
```

`UnreadByte` 是单步回退机制，不是通用回滚，不是随便能调用的，必须满足：

- 上一次操作是成功的 `ReadByte`
- 中间不能有任何其他 Read / Seek / ReadRune
- 否则会返回 error

#### `UnreadRune`

字符级别回退，撤销最近一次成功的 `ReadRune`，回退该 rune 占用的字节数，实现的 `io.RuneScanner` 接口。

```
func (r *Reader) UnreadRune() error
```

比 `UnreadByte` 更重要一些

- 因为在 文本解析中判断通常都是 rune 级别（字母/数字/空白/符号），所以回退也是 rune 几倍
- 如果用 `ReadRune` 读取了一个字符，不能用 `UnreadByte` 进行回退，必须使用 `UnreadRune`
- 通用只能撤销最近一次成功的 `ReadRune` 并且期间不能有任何其它读取或 Seek 操作

#### `WriteTo`

Reader 的收尾方法，把 Reader 中剩余未读的内容直接写入 w，实现了 `io.WriteTo` 接口

```
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
```

> 如果写 `io.Copy(w, r)`，`io.Copy` 会先检查 r 是否实现了 `WriteTo` 接口

#### `Reset`

Reader 数据源切换器，用于使用一个新的 string 替换 Reader 的数据源并会把读指针重置到起点

```
func (r *Reader) Reset(s string)
```

等价于：

```
r.s = s // 换数据源
r.i = 0 // 重置读指针
r.prevRune = -1 // 清空可回退 rune 的状态
```

> 不复制字符串，不分配内存，和 `NewReader` 功能上是等价的，Reset 主要是为了对象复用，避免重新分配 Reader


