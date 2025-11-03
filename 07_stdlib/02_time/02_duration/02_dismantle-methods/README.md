## Duration 拆解

因为 `time.Duration` 本质是 int64 纳秒。所以所有拆解方法都是基于纳秒 -> 目标单位进行转换。

主要方法如下：

| 方法                                       | 含义     |
|------------------------------------------|--------|
| `func (d Duration) Nanoseconds() int64`  | 纳秒     |
| `func (d Duration) Milliseconds() int64` | 毫秒     | 
| `func (d Duration) Microseconds() int64` | 微秒     |       
| `func (d Duration) Seconds() float64`    | 秒，可小数  | 
| `func (d Duration) Minutes() float64`    | 分钟，可小数 | 
| `func (d Duration) Hours() float64`      | 小时，可小数 | 
