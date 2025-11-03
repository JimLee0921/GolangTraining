package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// 1. Add 加减固定时间长度

	t1 := t.Add(2 * time.Hour)     // 加两小时
	t2 := t.Add(-30 * time.Minute) // 减三十分钟
	t3 := t.Add(90 * time.Second)  // 加90秒
	fmt.Println("======== Add ========")

	fmt.Println(t.Format(time.DateTime))
	fmt.Println(t1.Format(time.DateTime))
	fmt.Println(t2.Format(time.DateTime))
	fmt.Println(t3.Format(time.DateTime))

	// 2. AddDate 年/月/日 按照日历逻辑加减
	t4 := t.AddDate(0, 1, 0)  // 加一个月
	t5 := t.AddDate(-1, 0, 0) // 减一年
	t6 := t.AddDate(0, 0, 7)  // 加七天
	fmt.Println("======== AddDate ========")
	fmt.Println(t4.Format(time.DateTime))
	fmt.Println(t5.Format(time.DateTime))
	fmt.Println(t6.Format(time.DateTime))

	// 3. Sub 求两个时间差，得到 duration 对象（本质是纳秒整数）
	d := t2.Sub(t1)
	fmt.Println("======== Sub ========")
	fmt.Println(d, d.Minutes(), d.Seconds())

	// 4. Since / Until 方便与现在比较
	start := time.Now()
	time.Sleep(120 * time.Millisecond)
	elapsed := time.Since(start)
	fmt.Println("======== Since / Until ========")
	fmt.Println(elapsed)

	future := time.Now().Add(10 * time.Second)
	remain := time.Until(future)
	fmt.Println(remain) // 10s
}
