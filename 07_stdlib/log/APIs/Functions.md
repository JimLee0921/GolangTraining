# log 模块顶层函数

这些顶层函数其实都是用一个隐藏的全局 Logger（叫 std）来输出日志

## 顶层函数底层实现

这些顶层函数全部都在用 log 模块中定义的一个隐藏的全局变量：

```
var std = New(os.Stderr, "", LstdFlags)
```

- 输出到 stderr
- 默认前缀 ""
- 默认 `flags = Ldate | Ltime`

### `log.Default`

log.Default() 这个函数的实现就是返回 log 全局模块定义的的 `std` 变量

```
func Default() *Logger
```

> `log.Default()` 与 `&std` 在概念上是一致的

### 普通输出

* `Print(v ...any)`
* `Println(v ...any)`
* `Printf(format string, v ...any)`

### Panic 系列（输出 + panic）

* `Panic(v ...any)`
* `Panicln(v ...any)`
* `Panicf(format string, v ...any)`

### Fatal 系列（输出 + os.Exit(1)）

* `Fatal(v ...any)`
* `Fatalln(v ...any)`
* `Fatalf(format string, v ...any)`

### 配置类函数

* `SetOutput(w io.Writer)`
* `SetFlags(flag int)`
* `SetPrefix(prefix string)`
* `Flags()`
* `Prefix()`

