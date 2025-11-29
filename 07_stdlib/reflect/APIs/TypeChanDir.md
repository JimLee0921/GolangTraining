# `reflect.ChanDir`

在 reflect 包里，ChanDir 定义为

```
type ChanDir int。
```

表示一个 channel 类型的方向（direction），用来描述某个 channel 是可接收 (`<-chan`)、可发送 (`chan<-`)、还是双向 (`chan`)

对应的常量包括：

```
const (
    RecvDir ChanDir = 1 << iota  // <-chan
    SendDir                      // chan<-
    BothDir = RecvDir | SendDir  // chan (双向)
)
```

- `RecvDir` 表示 “只接收 channel”（`<-chan T`）
- `SendDir` 表示 “只发送 channel”（`chan<- T`）
- `BothDir` 表示 “普通双向 channel”（`chan T`）

## 方法

ChanDir 类型有一个方法 String

```
func (d ChanDir) String() string
```

可以把方向打印成字符串 (`<-chan / chan<- / chan`)

## 存在意义

这种方向区分不仅在编译期影响类型，也在反射 / 类型检查 / 动态生成类型时有意义。

在使用 `reflect.ChanOf()` 动态创建一个 channel 类型需要传入对应方向常量。

当通过反射拿到一个 `reflect.Type`，它可能对应普通 channel，也可能对应只读 channel / 只写 channel。
对 reflect 来说，需要一个清晰、类型安全的方法用来知道这是哪种 channel。

某些泛型、动态构造、类型比较、序列化/反序列化库（或者自己写的框架）如果要支持 channel 类型，就必须知道 channel 的方向，
因为在 Go 中，`chan T`、`<-chan T`、`chan<- T` 是三种不同的类型。

因此，在 reflect 里引入 ChanDir，就让人可以通过反射安全地获取到 channel 的方向。