# Logger

Logger 是 log 包中的一个结构体，用来执行真正的日志输出动作。

默认情况下使用的 `log.Println("hi")` 其实底层调用的就是一个默认的 Logger 实例： std。
但真正可控，可拓展的日志是通过自己创建的 Logger 实例完成的。

## 创建 Logger

使用 `log.New()` 函数进行自定义 Logger 实例

```
func New(out io.Writer, prefix string, flag int) *Logger
```

**参数**

- out：日志往哪儿写（实现 `io.Writer` 接口），测试建议使用 `os.Stdout` 输出到控制台，后续换成文件或网络等
- prefix：每条日志前面的前缀字符串，比如 `[INFO]/[ERROR]`，主要用于区分不同类型日志
- flag：控制日志格式（要不要日期、时间、文件名等）

**常见 flag**

```
const (
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒：01:23:23.123123（依赖 Ltime）
    Llongfile                     // 完整路径：/a/b/c/d.go:23
    Lshortfile                    // 文件名：d.go:23（覆盖 Llongfile）
    LUTC                          // 使用 UTC 时间而不是本地时间
    Lmsgprefix                    // 前缀放在消息前而不是日志头前
    LstdFlags     = Ldate | Ltime // 默认 flag：日期 + 时间
)
```

> flag 可以通过按位或 `|` 进行组合：`logger := log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)`

## 常用方法

Logger 的方法主要分为四大类：普通输出、Fatal、Panic、配置方法。

### 普通输出方法

- `Print`：等价于 `fmt.Print`，输出不自动换行
- `Printf`：格式化输出
- `Println``：自动换行（最常用）

```
func (l *Logger) Print(v ...any)
func (l *Logger) Printf(format string, v ...any)
func (l *Logger) Println(v ...any)
```

### Fatal 终止程序

Fatal 系列方法输出日志之后，会直接退出程序，也就是执行 `os.Exit(1)`，
这三个 Fatal 方法都也是输出后立即退出，不执行 defer。

```
func (l *Logger) Fatal(v ...any)
func (l *Logger) Fatalln(v ...any)
func (l *Logger) Fatalf(format string, v ...any)
```

> Fatal 不会执行 defer，因此关闭数据库/文件的代码不会跑

### panic 方法

这些是带 panic 的方法（会抛出 panic）

```
func (l *Logger) Panic(v ...any)
func (l *Logger) Panicln(v ...any)
func (l *Logger) Panicf(format string, v ...any)
```

**行为**

1. 输出日志
2. 调用 `panic()`
3. 会执行 defer（和 Fatal 不同）
4. 如果没 recover，就导致程序崩溃

### 配置相关

Logger 的配置类方法（开发中非常常用）

```
func (l *Logger) SetFlags(flag int)
func (l *Logger) SetOutput(w io.Writer)
func (l *Logger) SetPrefix(prefix string)
```

- `SetPrefix`：用于修改前缀
- `SetFlags`：用于动态改变输出格式
- `SetOutPut`：动态改变输出目标，不如某些级别的日志不需要可以使用 `logger.SeOutPut(io.Discard)` 进行舍弃

### 查询方法 Getter

```
func (l *Logger) Flags() int
func (l *Logger) Prefix() string
func (l *Logger) Writer() io.Writer // Go1.17+
```

- `Flags`：返回当前 Logger 实例的 flag
- `Prefix`：返回当前 Logger 实例的前缀 prefix
- `Writer`：查看日志输出指向哪里，返回 `io.Writer` 的实现

### 底层方法

一般不直接使用，了解即可。

```
func (l *Logger) Output(calldepth int, s string) error
```

- `calldepth`：调用栈深度（为了定位正确的文件与行号）
- s：最终要写入的字符串（已经格式化好的）

`Output `是 logger 所有 Print/Fatal/Panic 的最底层方法，比如 Println 最终会调用：`l.Output(2, "xxx")`