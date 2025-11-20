# 基本类型转为字符串

`Format` 系列是将基本类型转为字符串。

## 方法总览

| 函数                                   | 输入类型         | 输出                   | 说明      |
|--------------------------------------|--------------|----------------------|---------|
| `FormatBool(b)`                      | `bool`       | `"true"` / `"false"` | 布尔转字符串  |
| `FormatInt(i, base)`                 | `int64`      | 对应进制的字符串表示           | 有符号整数   |
| `FormatUint(u, base)`                | `uint64`     | 对应进制的字符串表示           | 无符号整数   |
| `FormatFloat(f, fmt, prec, bitSize)` | `float32/64` | 浮点数字符串               | 可控精度和格式 |

### FormatBool

```
func FormatBool(b bool) string
```

FormatBool 根据 b 的值返回 true 或 false

```
strconv.FormatBool(true) // "true"
strconv.FormatBool(false) // "false"
```

### FormatInt

```
func FormatInt(i int64 , base int ) string
```

* `base` 进制，取值 `2 ~ 36`
* 常用：2（二进制），8（八进制），10（十进制），16（十六进制）

```
strconv.FormatInt(255, 10) // "255"
strconv.FormatInt(255, 2) // "11111111"
strconv.FormatInt(255, 16) // "ff"
```

### FormatUint

```
func FormatUint(i uint64 , base int ) string
```

```
strconv.FormatUint(255, 16) // "ff"
```

---

### FormatFloat

```
func FormatFloat(f float64 , fmt byte , prec, bitSize int) string
```

| 参数        | 意义                                    |
|-----------|---------------------------------------|
| `f`       | 浮点数                                   |
| `fmt`     | 格式（`'f'`, `'e'`, `'E'`, `'g'`, `'G'`） |
| `prec`    | 小数精度位数（有效位数）                          |
| `bitSize` | `32` 或 `64` 对应 `float32` / `float64`  |

**常用格式符**

| 格式符   | 输出风格          | 指数类型      | 大小写风格 | 用途特点                 |
|-------|---------------|-----------|-------|----------------------|
| **b** | 二进制科学计数法      | **二进制指数** | 小写    | 用于调试，精确还原 float      |
| **e** | 十进制科学计数法      | 十进制指数     | `e`   | 常规科学计数法，默认指数输出小写     |
| **E** | 十进制科学计数法      | 十进制指数     | `E`   | 同 `e` 但指数标记为大写       |
| **f** | 普通小数          | **无指数**   | -     | 人类可读最常用输出（如金额、百分比）   |
| **g** | 自适应 `f` 或 `e` | 自适应       | 小写    | 当数字太大或太小时自动选择 `e`    |
| **G** | 自适应 `f` 或 `E` | 自适应       | 大写    | 与 `g` 类似但指数为 `E`     |
| **x** | 十六进制科学计数法     | 二进制指数     | `p`   | IEEE754 浮点精确格式用于底层分析 |
| **X** | 十六进制科学计数法     | 二进制指数     | `P`   | 同 `x`，字母为大写版本        |

| 格式符   | 输出示例格式              | 示例数 f=123.456   |
|-------|---------------------|-----------------|
| **b** | `-ddddp±ddd`        | `8670418p-16`   |
| **e** | `-d.dddde±dd`       | `1.2346e+02`    |
| **E** | `-d.ddddE±dd`       | `1.2346E+02`    |
| **f** | `-ddd.dddd`         | `123.4560`      |
| **g** | 大 or 小 → `e`，否则 `f` | `123.456`       |
| **G** | 与 g 类似但指数标记为 `E`    | `123.456`       |
| **x** | `-0xd.ddddp±ddd`    | `0x1.e2406p+06` |
| **X** | 与 x 类似但字母大写         | `0X1.E2406P+06` |

```
strconv.FormatFloat(3.1415926, 'f', 2, 64) // "3.14"
strconv.FormatFloat(3.1415926, 'e', 4, 64) // "3.1416e+00"
strconv.FormatFloat(3.1415926, 'g', 6, 64) // "3.14159"
```

## `Itoa`

等价于 FormatInt，小 int 的快捷方式

| 函数                    | 描述              |
|-----------------------|-----------------|
| `strconv.Itoa(i int)` | `int -> 十进制字符串` |

```
strconv.Itoa(123) // "123"
```

本质上等价于：

```
strconv.FormatInt(int64(123), 10)
```



