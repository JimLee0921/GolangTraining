package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Date
			根据年月日+时分秒+纳秒+时区构造
			month 可以传 int 也可以使用 time.January 这种枚举
	*/
	t1 := time.Date(2025, 10, 18, 15, 30, 45, 123456789, time.Local)
	fmt.Println("Date: ", t1)

	/*
		Unix
			秒+纳秒（用于精准）
	*/
	t2 := time.Unix(1737200000, 0)
	fmt.Println("Unix: ", t2)

	/*
		UnixMilli
			毫秒级时间戳
	*/
	ms := int64(1737200000000) // 2025-01-18 的某毫秒值
	t3 := time.UnixMilli(ms)
	fmt.Println("UnixMilli: ", t3)

	/*
		UnixMicro
			微秒级时间戳
	*/
	us := int64(1737200000000000)
	t4 := time.UnixMicro(us)
	fmt.Println("UnixMicro:", t4)

}
