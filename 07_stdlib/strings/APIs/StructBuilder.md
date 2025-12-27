# `strings.Builder`

`strings.Builder` 是一个用于高效，连续构造字符串的专用工具类型

```
type Builder struct {
	// contains filtered or unexported fields
}
```

## 存在意义

在 Go 中字符串是不可变的：

```
s := "a"
s = s + "b"
s = s + "c"
```

每次 `+` 都会创建新的 string 并复制已有的内容，时间复杂度趋近于 `O(n²)`

多次拼接是一个非常常见但是很危险的场景，比如循环、条件分支中拼接字符串时并不知道最终字符串有多长，而使用 `strings.Builder`时
用一次可以增长的 buffer，承载多次写入，最终一次性生成 string，中间阶段是通过 `[]byte` ，而用户只关心最终的 string 结果

- 是为构造型字符串专门设置的根据
- 为替代 `s += xxx` 而存在
- 是一个写完及丢的一次性对象
- 最大限度减少内存复制

## 方法集

`strings.Builder` 方法可用分为四类：

1. 状态查询类（state）：`Len`、`Cap`
2. 容量管理类（capacity）：`Grow`
3. 写入类（write）：`Write`、`WriteByte`、`WriteRune`、`WriteString`
4. 生命周期类（lifecycle）：`String`、`Reset`

### 状态查询类方法

#### `Len`

返回累计字节数，也就是已经写入的字节数。`b.Len() == len(b.String())`

```
func (b *Builder) Len() int
```

#### `Cap`

返回构造器底层字节切片的容量，视为正在构造的字符串分配的总空间，包括所有已经写入的字节。

```
func (b *Builder) Cap() int
```

> `Cap >= Len`，因为扩容是指数式的

### 容量管理类方法

#### `Grow`

增加构造器底层字节切片的容量，确保在不发生重新分配的情况下，Builder 还能再写入至少 n 个字节

```
func (b *Builder) Grow(n int)
```

- 如果 `Cap - Len >= n` 则什么都不做
- 如果不足会扩容到 `Len + n`
- `n > 0` 必须成立，否则会触发 panic

> Grow 只是性能提示，不是必要调用，即使不调用，底层在遇到容量不足时也会自动扩容，只是为了减少扩容次数，避免多次内存分配+拷贝

### 写入类方法

把不同形态的数据，以最低成本的路径，追加到内部的 `byte buffer`。

#### WriteString

最重要，最常用的 Write，把一个 string 直接追加到 b 中

```
func (b *Builder) WriteString(s string) (int, error)
```

- 返回写入的字节数和一个 nil 错误（错误永远为 nil，因为实现了 `io.StringWriter` 接口，目的是生态兼容，而不是处理错误）
- 因为 `Write([]byte(s))` 每次都需要分配一个新的 `[]byte`，需要二次拷贝，使用 `WriteString` 运行 Builder 直接按 string
  的字节序列写入

#### WriteRune

面向 Unicode 语义的入口，用于把一个 `rune` 按照 `UTF-8` 的编码写入 Builder

```
func (b *Builder) WriteRune(r rune) (int, error)
```

- 不是写一个单位，而是写 `utf8.RuneLen(r)` 个字节
- 用于字符串遍历、Unicode 映射或手写 parser/tokenizer

#### WriteByte

极小粒度的性能优化，向 Builder 追加一个字节，微优化 API，直接写入一个字节，用于替代 `b.Write([]byte{c})`

```
func (b *Builder) WriteByte(c byte) error
```

主要用于写分隔符、手动格式化逻辑或高性能日志拼接等

#### `Write`

接口兼容型写入，把一个 byte slice 的内容拷贝进 Builder，这也是 Builder 实现 `io.Writer` 接口的原因

```
func (b *Builder) Write(p []byte) (int, error)
```

- 是为了能让 Builder 当作 `io.Writer` 进行传递，接入通用 IO 生态（fmt/log/encoder），而不是为了字符串拼接本身
- 注意只有拿到 `[]byte` 时才能使用 `Write`

### 生命周期类方法

#### `String`

Builder 的终点方法，返回当前 Builder 中内容对应的 string，最核心也是最危险的方法，一旦调用了 String 方法，就应当停止继续操作这个
Builder

```
func (b *Builder) String() string
```

Builder 是不能被复制的，里面有个关键字段 `addr *Builder` 就是为了防止在已经调用过 `String()` 的情况下，Builder 被非法复制或继续操作，
因为如果 `String()` 返回的 string 和 buffer 共享内存，那么再写 Builder 就会篡改已经返回给外部的 string

#### `Reset`

清空 Builder 的内容，使其回到未使用状态，重点是状态而不是内存

```
func (b *Builder) Reset()
```

调用后会把 `b.Len()` 设为 0，清空内部状态标记，并允许重新开始写入。