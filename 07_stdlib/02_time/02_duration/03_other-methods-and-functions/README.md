## `time.Duration` 计算

因为本质上是 int 类型，所以支持一些数学计算

常见操作与方法如下：

| 方法/操作              | 用途                     | 示例                                       |
|--------------------|------------------------|------------------------------------------|
| 比较（`<`, `>`, `==`） | 判断时间长短                 | `d1 < d2`                                |
| 加/减 Duration       | 时间偏移计算                 | `t.Add(5*time.Minute)`                   |
| 相反数（取负）            | 用于回退、逆向偏移              | `-d`                                     |
| 绝对值                | 获取 Duration 的正值        | `abs := d; if abs < 0 { abs = -abs }`    |
| 四舍五入到最近单位          | `Round()` `Truncate()` | `d.Round(time.Second)`                   |
| 转换字符串              | `d.String()`           | `fmt.Println(90*time.Second)` -> `1m30s` |

### 比较

```
d1 := 3 * time.Second
d2 := 5 * time.Second

fmt.Println(d1 < d2)  // true
fmt.Println(d1 > d2)  // false
fmt.Println(d1 == 3*time.Second) // true
```

> 非常直观常用，不需要额外方法

### 取负（反向时间间隔）

```
d := 30 * time.Second
fmt.Println(-d) // -30s
```

常与 `t.Add()` 配合使用

```
now := time.Now()
past := now.Add(-30 * time.Minute)
```

### 四舍五入和向下取整

> Round 和 Truncate，不常用

`func (d Duration) Round(m Duration) Duration`

Round 函数返回将 d 按照传入的 m 进行四舍五入到最接近的 m 的倍数的结果。对于中间值，四舍五入的原则是远离零。
如果 m <= 0，则 Round 函数返回 d 不变。

`func (d Duration) Truncate(m Duration) Duration`

Truncate 函数返回将 d 向下取整四舍五入到 m 的倍数的结果。如果 m <= 0，则 Truncate 函数返回 d 不变。

```
d := 1250 * time.Millisecond // 1.25s
fmt.Println(d.Round(time.Second))    // 1s  四舍五入
fmt.Println(d.Truncate(time.Second)) // 1s  向下取整
```

### 相除

可以使用整数的除法来提取分钟+剩余秒等信息

```
d := 125 * time.Second

minutes := d / time.Minute   // 2
seconds := d % time.Minute / time.Second // 5

fmt.Println(minutes, seconds) // 2 5
```

### 求绝对值

```
func (d Duration) Abs() Duration
```

Abs 返回 d 的绝对值。作为一种特殊情况，使其幅度减少 1 纳秒

### Since / Until

```
func Since(t Time ) Duration
```

返回的是自 t 以来经过的时间。是 `time.Now().Sub(t)` 的简写

```
func Until(t Time ) Duration
```

Until 返回到 t 的持续时间。它是 `t.Sub(time.Now())` 的简写