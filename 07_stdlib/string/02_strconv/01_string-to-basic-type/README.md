# 字符串转为基本类型

## 常用 Parse 方法

| 函数           | 作用      | 示例输入                         | 返回类型      | 举例                        |
|--------------|---------|------------------------------|-----------|---------------------------|
| `ParseBool`  | 解析布尔值   | "true" / "false" / "1" / "0" | `bool`    | `"true" -> true`          |
| `ParseInt`   | 解析有符号整数 | "123" / "-45"                | `int64`   | `"42" -> int64(42)`       |
| `ParseUint`  | 解析无符号整数 | "123"                        | `uint64`  | `"42" -> uint64(42)`      |
| `ParseFloat` | 解析浮点数   | "3.14" / "1e9"               | `float64` | `"3.14" -> float64(3.14)` |

---

### ParseBool

```
func ParseBool(str string) (bool, error)
```

支持如下表达：

| 字符串                                       | 解析结果    |
|-------------------------------------------|---------|
| `"1"`, `"t"`, `"T"`, `"true"`, `"TRUE"`   | `true`  |
| `"0"`, `"f"`, `"F"`, `"false"`, `"FALSE"` | `false` |

示例：

```
v, err := strconv.ParseBool("true")
```

### ParseInt

```
func ParseInt(s string, base int, bitSize int) (i int64, err error)
```

* `s`: 字符串（如 `"123"`, `"-FF"`）
* `base`: 进制（2~36），`0` 表示自动识别 (`0x` -> 16进制)
* `bitSize`: 转成 `int8 / int16 / int32 / int64`

```
n, err := strconv.ParseInt("101", 2, 64) // 二进制 -> 5
n, err := strconv.ParseInt("FF", 16, 64) // 十六进制 -> 255
n, err := strconv.ParseInt("-42", 10, 32) // -42
```

### ParseUint

```
func ParseUint(s string, base int, bitSize int) (uint64, error)
```

同上，但不允许负数：

```
u, err := strconv.ParseUint("FF", 16, 64) // 255
```

### ParseFloat

```
func ParseFloat(s string, bitSize int) (float64, error)
```

* `bitSize = 32` -> 返回 `float32`
* `bitSize = 64` -> 返回 `float64`

```
f, err := strconv.ParseFloat("3.14", 64) // float64(3.14)
f, err = strconv.ParseFloat("1e6", 64) // 科学计数法也可以
```

---

### `Atoi` 和 `Itoa`

对应十进制整数的常用快速转换：

| 函数                | 作用         | 基于                        |
|-------------------|------------|---------------------------|
| `strconv.Atoi(s)` | 字符串 -> int | `ParseInt(s, 10, 0)`      |
| `strconv.Itoa(i)` | int -> 字符串 | `FormatInt(int64(i), 10)` |

```
i, err := strconv.Atoi("123")
s := strconv.Itoa(123) // "123"
```

## bitSize 快速记忆

| bitSize | 返回值应转换成                  |
|---------|--------------------------|
| 8       | int8 / uint8             |
| 16      | int16 / uint16           |
| 32      | int32 / uint32 / float32 |
| 64      | int64 / uint64 / float64 |

