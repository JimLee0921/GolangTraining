# `time.Time`

`time.Time` 是 Go 用来表示某一个具体时刻的结构体，
包含一个时间点（纳秒级别）和 一个时区（Location）。

使用时间的程序通常应该将时间存储和传递为值，而不是指针。
也就是说，时间变量和结构体字段应该是 `time.Time` 类型，而不是 `*time.Time` 类型。

多个 goroutine 可以同时使用 Time 值，
但`Time.GobDecode`、`Time.UnmarshalBinary`、`Time.UnmarshalJSON`等一些t方法不是并发安全的。



> time.Time 永远代表一个真实发生的时刻，即使格式看起来不同

