# strconv 顶层方法

strconv 作为工具包，基本都是工具函数

## 字符可打印性/图形性判断

服务于 Quote/转义策略，这组通常不会单独用，而是在自定义转义、过滤输出、日志可读性时用到

- `IsPrint(r rune) bool`：是否为可打印字符
- `IsGraphic(r rune) bool`：是否为图形字符（通常不包括空格等）

### `strconv.IsPrint`

判断一个 rune 是否属于可打印字符（printable）

可打印字符指的是在文本中可以直接显示，不需要转义且不会破坏阅读结构的字符，该函数主要服务于：

- 字符串转义策略
- `Quote/AppendQuote` 的内部实现
- 日志打印调试输出的可读性判断

```
func IsPrint(r rune) bool
```

**参数**

- `r rune`：一个 Unicode code point

**返回值**

- `bool`：
    - `true`：该字符被认为是可以打印的
    - `false`：不可打印，通常需要转义

**注意事项**

1. 判定依据是 Unicode 定义的可打印字符集合
2. 包括：字母/数字/标点符号/空格
3. 不包括：控制字符（如 `\n`、`\t`、`\xOO`）
4. 并不等价于 ASCII 可打印字符，很多非 ASCII 字符（如中文）返回 true

### `strconv.IsGraphic`

判断一个 rune 是否属于图形字符，图形字符比可打印字符更严格，必须要是可见的，有形状的字符

```
func IsGraphic(r rune) bool
```

**参数**

- `r rune`：一个 Unicode code point

**返回值**

- `bool`：
    - `true`：该字符被认为是图形字符
    - `false`：该字符被认为是非图形字符

**注意事项**

1. 不包含空格，空格是可打印的，但不是图形字符
2. 包括：字母/数字/标点/表意文字（如中文）
3. 主要用于：`QuoteToGraphic`、`AppendQuoteToGraphic`
4. 目标是最大化可读性，最小化转义

## 引号/转义相关

生成或解析带引号字面量，可以把普通字符串/rune转成带引号并必要转义的形式，或者反过来解析

- 生成带引号字符串（返回 `string`）

    - `Quote(s string) string`
    - `QuoteToASCII(s string) string`
    - `QuoteToGraphic(s string) string`

- 生成带引号 rune（返回 `string`）

    - `QuoteRune(r rune) string`
    - `QuoteRuneToASCII(r rune) string`
    - `QuoteRuneToGraphic(r rune) string`

- 解析带引号字面量

    - `Unquote(s string) (string, error)`：把 `"..."`、`'...'`、`` `...` `` 这种形式还原成裸字符串
    - `UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)`：解析--一个--转义/字符并返回剩余
      tail（做词法扫描器时很常用）
    - `QuotedPrefix(s string) (string, error)`：从开头截出“一个完整的 quoted literal 前缀”（做流式解析/扫描时用）

- 反引号能力判断

    - `CanBackquote(s string) bool`：判断能否用反引号 `` `...` `` 表示（即无需解释转义、且内容满足反引号字面量约束）

### `strconv.Quote`

将一个普通字符串转换为带双引号、并进行必要转义的字符串字面量表示，生成的结果符合 Go 源码中的字符串字面量语法

```
func Quote(s string) string
```

**参数**

- `s string`：原始字符串内容（不带引号）

**返回值**

- `string`：带双引号的字符串字面量，例如 `"hello\nworld"`

**注意事项**

1. 返回值一定包含双引号 `"`
2. 控制字符，不可打印字符会被转义（`\n`、`\t`、`\xNN`等）
3. 使用的是 Go 字符串转义规则，不是 JSON 规则
4. 常用于调试输出，错误信息和生成 Go 源码片段

### `strconv.QuoteToASCII`

与 Quote 类似，但强制将所有非 ASCII 字符转义，确保返回字符串只包含 ASCII 字符

```
func QuoteToASCII(s string) string
```

**参数**

- `s string`：原始字符串内容

**返回值**

- `string`：只包含 ASCII 的带引号字符串，例如 `"\u4e2d\u6587"`

**注意事项**

1. 非 ASCII 字符会被编码为 `\uXXXX`
2. 可读性降低，但可移植性和终端兼容型更好
3. 适合纯 ASCII 日志系统和跨平台文本输出

### `strconv.QuoteToGraphic`

把字符串转换为带双引号形式，仅对非图形字符进行转义，尽可能保留可见字符的原样输出

```
func QuoteToGraphic(s string) string
```

**参数**

- `s string`：原始字符串内容

**返回值**

- `string`：带引号、偏向可读性的字符串表示

**注意事项**

1. 图形字符”的判断标准来自 `strconv.IsGraphic`
2. 相对于 Quote，转义更少，可读性更高
3. 适合人类可读日志和调试输出

### `strconv.QuoteRune`

将一个 rune 转换为带单引号的字符字面量，结果符合 Go 的 rune 字面量语法

```
func QuoteRune(r rune) string
```

**参数**

- `r rune`：一个 Unicode code point

**返回值**

- `string`：例如 `'a'`、`'\n'`、`'中'`

**注意事项**

1. 返回值一定使用单引号 `'`
2. 控制字符会被转义
3. 常用于字符级调试，错误提示和语法分析工具

### `strconv.QuoteRuneToASCII`

与 QuoteRune 相同，但强制使用 ASCII 表示

```
func QuoteRuneToASCII(r rune) string
```

**参数**

- `r rune`：一个 Unicode code point

**返回值**

- `string`：例如 `'\u4e2d'`

**注意事项**

1. 非 ASCII rune 会被转成 `\uXXXX`
2. 常用于 ASCII-only 输出环境或跨终端兼容日志

### `strconv.QuoteRuneToGraphic`

将 rune 转换为带单引号形式，仅对非图形字符进行转义，尽量保持可见性

```
func QuoteRuneToGraphic(r rune) string
```

**参数**

- `r rune`：一个 Unicode code point

**返回值**

- `string`：偏向可读的 rune 字面量表示

**注意事项**

1. 基于 `strconv.IsGraphic` 进行判断
2. 输出更贴近人类阅读

### `strconv.CanBackquote`

判断一个字符串是否可以原封不动的放入反引号（`` `...` `` 即 Raw String Literal）中进行表示，也就是能否作为 Go 原始字符串字面量

```
func CanBackquote(s string) bool
```

**参数**

- `s string`：原始字符串内容（不带引号）

**返回值**

- `bool`
    - `true`：可以使用反引号表示
    - `false`：必须使用双引号并转义

**注意事项**

1. 只要包含反引号字符 `` ` `` ，就返回 false
2. `\t` 是可以出现的， `\n`、`\r` 等控制符不行
3. 主要用于生成 Go 源码和服务于 `Quote` / `AppendQuote`

### `strconv.Unquote`

将一个完整的、带引号的字符串或 rune 字面量还原为裸字符串内容

```
func Unquote(s string) (string, error)
```

**参数**

- `s string`：带引号的字面量，例如：`"hello\n"`、`'a'`、`` `raw\ntext` ``

**返回值**

- `string`：还原后的字符串内容
- `error`：字面量非法或转移错误时返回

**注意事项**

1. 支持双引号字符串/单引号 rune(如果是单引号里面必须只有一个rune，结果仍是 string)/反引号字符串
2. 会解析并处理转义序列（`\n`、`\uXXXX` 等）
3. 输入必须是完整字面量

### `strconv.UnquoteChar`

从字符串开头解析一个字符（rune）或一个转义序列，并返回解析结果及剩余未处理部分，这是一个底层扫描函数，为了服务其它API，基本用不到

```
func UnquoteChar(
    s string,
    quote byte,
) (value rune, multibyte bool, tail string, err error)
```

### `strconv.QuotedPrefix`

从字符串开头解析并返回一个完整的、合法的带引号字面量前缀，如果开头不是合法的 quoted literal，则返回错误，
就是把一个字符串开头被引号包裹的部份完整抠出来。

```
func QuotedPrefix(s string) (string, error)
```

**参数**

- `s string`：以带引号字面量开头的字符串，双引号，单引号字符或反引号都可以

**返回值**

- `string`：截取出的完整的 quoted literal
- `error`：语法不合法时返回的错误

**注意事项**

1. 支持的引号格式参考 Unquoted
2. 只解析被引号包裹的前缀，不关心后面还有什么

## 面向 `[]byte` 追加输出

把结果追加到现有缓冲区，常用于日志、编码器、网络协议拼包，避免 `string` 分配与拼接成本。

- 基本类型追加
    - `AppendBool(dst []byte, b bool) []byte`
    - `AppendInt(dst []byte, i int64, base int) []byte`
    - `AppendUint(dst []byte, i uint64, base int) []byte`
    - `AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte`

- 把一个普通字符串/字符转换为带引号并进行必要转义的形式，生成结果符合 Go 字符串字面量语法（用于生成 Go/调试可读文本等）

    - `AppendQuote(dst []byte, s string) []byte`
    - `AppendQuoteToASCII(dst []byte, s string) []byte`
    - `AppendQuoteToGraphic(dst []byte, s string) []byte`
    - `AppendQuoteRune(dst []byte, r rune) []byte`
    - `AppendQuoteRuneToASCII(dst []byte, r rune) []byte`
    - `AppendQuoteRuneToGraphic(dst []byte, r rune) []byte`

### `strconv.AppendBool`

将一个 bool 值的文本表示（`"true"` 或 `"false"`）追加到已有的 `[]byte` 缓冲区中，用于构建输出内容而避免额外分配

```
func AppendBool(dst []byte, b bool) []byte
```

**参数**

- `dst []byte`：目标缓冲区，函数会在其末尾追加内容
- `b bool`：要被追加的布尔值

**返回值**

- `[]byte`：追加完成后的缓冲区（必须接收返回值）

**注意事项**

- 输出永远是小写 `"true"` / `"false"`
- 不返回错误（bool 没有非法值）
- 如果 dst 容量不足会自动扩容
- 常用于替代 `buf = append(buf, strconv.FromatBool(b)...)`

### `strconv.AppendInt`

将一个有符号整数按指定进制转换为字符串形式，并追加到已有 `[]byte` 中。

```
func AppendInt(dst []byte, i int64, base int) []byte
```

**参数**

- `dst []byte`：目标缓冲区，函数会在其末尾追加内容
- `i int64`：要转换的整数值（支持负数）
- `base int`：进制，范围 `2 <= base <= 36`

**返回值**

- `[]byte`：追加完成后的缓冲区（必须接收返回值）

**注意事项**

- 负号只由 AppendInt 处理
- base 超过 10 后使用 `a-z` 进行表示，如果 base 不合法时会触发 panic

### `strconv.AppendUint`

将一个 无符号整数 按指定进制转换为字符串形式，并追加到已有 `[]byte` 中

```
func AppendUint(dst []byte, i uint64, base int) []byte
```

**参数**

- `dst []byte`：目标缓冲区，函数会在其末尾追加内容
- `i uint64`：要转换的整数值（保证非负）
- `base int`：进制，范围 `2 <= base <= 36`

**返回值**

- `[]byte`：追加完成后的缓冲区（必须接收返回值）

**注意事项**

- 不会输出负号
- base 进制规则与 AppendInt 一致

### `strconv.AppendFloat`

将一个浮点数按指定格式、精度和位宽转换为字符串形式，并追加到已有 `[]byte` 中，strconv 中最复杂、最灵活的 Append 函数。

```
func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
```

**参数**

- `dst []byte`：目标缓冲区，函数会在其末尾追加内容
- `f float64`：要格式化的浮点值
- `fmt byte`：输出格式标识符：
    - `'f'` 定点
    - `'e' / 'E'` 科学计数法
    - `'g' / 'G'` 自动选择
- `prec int`：精度（含义取决于 fmt）
- `bitSize int`：输入值的位宽：32 或 64

**返回值**

- `[]byte`：追加完成后的缓冲区（必须接收返回值）

**注意事项**

- prec 含义随 fmt 改变：
    - fmt 为 `f`，prec 表示小数点后位数
    - fmt 为 `e/E`，prec 表示小数点后位数
    - fmt 为 `g/G`，prec 表示有效数字位数
    - fmt 为 `-1`，prec 表示最短且可逆表示
- bitSize 必须与数据来源一致
    - 32 -> float32
    - 64 -> float64
- 自动处理 `NaN`、`+Inf`、`-Inf`等特殊值

### `strconv.AppendQuote`

将一个普通字符串转换为 带双引号、并进行必要转义 的形式，然后追加到 dst，生成结果符合 Go 字符串字面量语法

```
func AppendQuote(dst []byte, s string) []byte
```

**参数**

- `dst []byte`：目标缓冲区，函数会在其末尾追加内容
- `s string`：原始字符串内容（不带引号）

**返回值**

- `[]byte`：追加完成后的缓冲区（必须接收返回值）

**注意事项**

- 输出一定 包含双引号
- 不可打印字符会被转义（比如 `\n`、`\t`、`\xNN` 等）
- 不是 JSON 转义规则，而是 Go 字符串规则
- 常用于日志调试/错误信息/生成 Go 源码片段

### `strconv.AppendQuoteToASCII`

与 AppendQuote 类似，但强制所有非 ASCII 字符转义，保证输出结果 只包含 ASCII 字符

```
func AppendQuoteToASCII(dst []byte, s string) []byte
```

参数和返回值与 AppendQuote 完全一致。

**注意事项**

1. 非 ASCII 字符会被转成 `\uXXXX`
2. 非常适合纯 ASCII 日志系统，终端兼容型要求高的输出
3. 可读性下降但是可移植性增强

### `strconv.AppendQuoteToGraphic`

将字符串转成带引号形式，仅对非图形字符进行转义，尽可能保留可见字符的原样输出

```
func AppendQuoteToGraphic(dst []byte, s string) []byte
```

参数和返回值同 AppendQuote

**注意事项**

1. 图形字符由 `strconv.Graphic` 判定
2. 通常比 AppendQuote 更偏向可读性
3. 常用于人可读日志/调试输出/REPL 显示

### `strconv.AppendQuoteRune`

将一个 rune 字符转换为带单引号的字符字面量并追加到 dst

```
func AppendQuoteRune(dst []byte, r rune) []byte
```

**参数**

- `dst []byte`：目标缓冲区
- `r rune`：Unicode code point

**返回值**

- `[]byte`：追加完成后的缓冲区

**注意事项**

1. 输出类型类似于：`a`、`\n`、`中`
2. 转义规则遵循 Go rune 字面量
3. 常用于字符级调试/语法分析器/错误提示

### `strconv.AppendQuoteRuneToASCII`

与 AppendQuoteRune 相同，但强制使用 ASCII 表示

```
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
```

参数和返回值同 AppendQuoteRune

**注意事项**

1. 非 ASCII rune 会被转为 `\uXXXX`
2. 适合： 纯 ASCII 输出环境/跨终端兼容日志

### `strconv.AppendQuoteRuneToGraphic`

将 rune 转为带引号形式，只转义非图形字符，尽量保持可见性

```
func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte
```

参数与返回值同 AppendQuoteRune

**注意事项**

1. 基于 `strconv.IsGraphic` 判断
2. 通常比 ASCII 版本更可读
3. 适合面向人的调试输出

## 基本类型转为 `string`

把值转成 `string`，相比 `fmt.Sprintf` 更轻、更可控，相比 Append 系列会产生 `string` 分配

- 基本类型格式化

    - `FormatBool(b bool) string`
    - `FormatInt(i int64, base int) string`
    - `FormatUint(i uint64, base int) string`
    - `FormatFloat(f float64, fmt byte, prec, bitSize int) string`
    - `FormatComplex(c complex128, fmt byte, prec, bitSize int) string`

- 便捷入口

    - `Itoa(i int) string`（等价于 `FormatInt(int64(i), 10)` 的便捷版）

### `strconv.FormatBool`

把 bool 值转为起字符串表示，`true` -> `"true"`，`false` -> `"false"`

```
func FormatBool(b bool) string
```

**参数**

- `b bool`：要格式化的布尔值

**返回值**

- string：`"true"` 或 `"false"`

**注意事项**

1. 输出永远是小写
2. 不会失败，不会返回 error
3. 等价但轻量于：`fmt.Sprintf("%t", b)`

### `strconv.FormatInt`

将一个有符号整数按指定进制转换为字符串

```
func FormatInt(i int64, base int) string
```

**参数**

- `i int64`：要转换的整数
- `base int`：进制 `2<=base<=36`

**返回值**

- string：转换后的字符串表示

**注意事项**

1. 支持复数，会自动加`-`
2. `base > 10` 时使用 `a-z`，且 base 非法会导致 panic

### `strconv.FormatUint`

将一个无符号整数按指定进制转换为字符串

```
func FormatUint(i uint64, base int) string
```

**参数**

- `i uint64`：要转换的整数，无符号整数
- `base int`：进制 `2<=base<=36`

**返回值**

- string：转换后的字符串表示

**注意事项**

1. 永远不会输出符号
2. 进制规则和 FormatInt 完全一致

### `strconv.FormatFloat`

将浮点数按指定 格式 / 精度 / 位宽 转换为字符串

```
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
```

**参数**

- `f float64`：要格式化的浮点值
- `fmt byte`：输出格式：

    - `'f'` 定点
    - `'e' / 'E'` 科学计数法
    - `'g' / 'G'` 自动选择
- `prec int`：精度，依赖于 fmt
- `bitSize int`：32 或 64

**返回值**

- string：转换后的浮点字符串

**注意事项**

1. prec 的语义不是统一的
    - fmt 为 `f`，prec 表示小数点后位数
    - fmt 为 `e/E`，prec 表示小数点后位数
    - fmt 为 `g/G`，prec 表示有效数字位数
    - fmt 为 `-1`，prec 表示最短且可逆表示
2. bitSize 会影响舍入行为
3. 自动处理 `Nan`、`+Inf`、`-Inf`
4. 不建议直接使用 `fmt.Sprintf("%f", ...)` 替代

### `strconv.FormatComplex`

将复数按指定格式转换为字符串，格式为：实部加虚部(real+imagi)

```
func FormatComplex(c complex128, fmt byte, prec, bitSize int) string
```

**参数**

- `c complex128`：复数值
- `fmt / prec / bitSize`：完全复用 FormatFloat 规则

**返回值**

- string：转换后的复数字符串表示

**注意事项**

1. 输出始终带括号
2. 虚部始终带 i
3. 内部是对实部/虚部分别调用浮点格式化

### `strconv.Itoa`

**最常用**，int 类型的便捷十进制字符串转换。

```
func Itoa(i int) string
```

**参数**

- `i int`：平台相关位宽的整数

**返回值**

- string：整数十进制表示字符串

**注意事项**

1. 等价于 `strconv.FromatInt(int64(i), 10)`

## 解析字符串

把字符串转为特定类型，可能会显式报错（语法错误/溢出），是配置、CLI、HTTP 参数、CSV 解析的主力

- 基本类型解析

    - `ParseBool(str string) (bool, error)`
    - `ParseInt(s string, base int, bitSize int) (int64, error)`
    - `ParseUint(s string, base int, bitSize int) (uint64, error)`
    - `ParseFloat(s string, bitSize int) (float64, error)`
    - `ParseComplex(s string, bitSize int) (complex128, error)`

- 便捷入口

    - `Atoi(s string) (int, error)`（等价于十进制 + int 位宽的便捷版）

### `strconv.ParseBool`

把字符串解析成布尔值

```
func ParseBool(str string) (bool, error)
```

**参数**

- `str string`：输入字符串（大小写不敏感）

**返回值**

- `bool`：解析结果
- `error`：非法输入时报错

**注意事项**

1. 只接受以下值（忽略大小写）
    - true：`1`、``t、`T`、`true`、`TRUE`
    - false：`0`、``f、`F`、`false`、`FALSE`
2. 不接受空字符串

### `strconv.ParseInt`

把字符串解析为有符号整数，支持进制和位宽控制

```
func ParseInt(s string, base int, bitSize int) (i int64, err error)
```

**参数**

- `s string`：输入字符串，可以包含 `+` / `-` 正负号
- `base int`：进制，如果输入 0 时自动识别，正常为 2-36
- `bitSize int`：目标位宽（8/16/32/64）

**返回值**

- `int64`：解析结果（按照 bitSize 截断验证）
- `error`：语法错误或溢出

**注意事项**

1. `base = 0` 是最为推荐的默认选择
2. bitSize 决定溢出检查
3. 返回类型始终为 int64 需要手动转换

### `strconv.ParseUint`

把字符串解析为无符号整数，正常进制和位宽

```
func ParseUint(s string, base int, bitSize int) (uint64, error)
```

**参数**

- `s string`：输入字符串，不能是负数
- `base int`：同 ParseInt
- `bitSize int`：目标位宽（8/16/32/64）

**返回值**

- `uint64`：解析结果（按照 bitSize 截断验证）
- `error`：语法错误或溢出

**注意事项**

1. 负号直接报错
2. base/bitSize 规则和 ParseInt 一致

### `strconv.Atoi`

ParseInt 的简化版，最常用，把字符串解析为十进制的 int

```
func Atoi(s string) (int, error)
```

**参数**

- `s string`：十进制整数字符串

**返回值**

- `int`：解析结果（平台相关位宽）
- `error`：非法或溢出时报错

**注意事项**

1. 等价于 `ParseInt(s, 10, strconv.IntSize)`
2. 不支持其它进制

### `strconv.ParseFloat`

把字符串解析为浮点数，支持科学计数法、小数、特殊值

```
func ParseFloat(s string, bitSize int) (float64, error)
```

**参数**

- `s string`：浮点数字符串（如 "3.14", "1e-9"）
- `bitSize int`：32 或 64

**返回值**

- `float64`：解析结果（按 bitSize 规则）
- `error`：语法或溢出错误

**注意事项**

1. 支持 `NaN`、`Inf`、`+Inf`、`-Inf`
2. bitSize 会影响舍入与溢出判断
3. 返回类型始终是 float64

### `strconv.ParseComplex`

把字符串解析为复数，格式为实部加虚部(real+imagi)

```
func ParseComplex(s string, bitSize int) (complex128, error)
```

**参数**

- `s string`：复数字符串（必须带括号）
- `bitSize int`：64（complex64）或 128（complex128）

**返回值**

- `complex128`：解析结果（按 bitSize 规则）
- `error`：语法或溢出错误

**注意事项**

1. 实部/虚部使用 ParseFloat 规则
2. 格式严格，不适合宽松输入