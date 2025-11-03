## 生成 `time.Duration`

`time.Duration` 主要有 3 种生成方式

| 方式                  | 示例                            | 使用场景            | 是否常用 |
|---------------------|-------------------------------|-----------------|:----:|
| 数字 × 时间单位常量         | `5 * time.Second`             | 代码中直接定义时间间隔     |  常用  |
| 时间差 Sub()           | `t2.Sub(t1)`                  | 求两个时间点之间的差值     |  常用  |
| ParseDuration 字符串解析 | `time.ParseDuration("1h30m")` | 配置文件、环境变量、命令行参数 | 不常用  |

### 数字 × 时间单位常量

Go 内置了时间单位常量：

```
time.Nanosecond
time.Microsecond
time.Millisecond
time.Second
time.Minute
time.Hour
```

可以直接：

```
d1 := 10 * time.Second        // 10 秒
d2 := 500 * time.Millisecond  // 0.5 秒
d3 := 2 * time.Hour           // 2 小时

fmt.Println(d1) // 10s
fmt.Println(d2) // 500ms
fmt.Println(d3) // 2h0m0s
```

### 时间差计算

可以使用 `time.Time.Sub()` 方法计算时间差，返回 Duration 对象

```
t1 := time.Now()
t2 := t1.Add(75 * time.Minute)

d := t2.Sub(t1)
fmt.Println(d) // 1h15m0s
```

> Duration = Time2 - Time1

### 字符串解析

可以使用 `time.ParseDuration()` 处理：

- 配置文件 （config.yaml / JSON）
- 环境变量
- 命令行参数
- 用户输入

字符串不正确会解析失败，不能硬编码，所以要解析

```
d, err := time.ParseDuration("1h30m")
if err != nil {
    panic(err)
}
fmt.Println(d) // 1h30m0s
```

**支持格式示例**

| 字符串       | 解析后   |
|-----------|-------|
| `"30s"`   | 30秒   |
| `"500ms"` | 500毫秒 |
| `"1m30s"` | 1分30秒 |
| `"2h"`    | 2小时   |
| `"-5s"`   | -5秒   |

