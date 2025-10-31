package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// 时间转时间戳方法
	fmt.Println(t.Unix())      // 秒级
	fmt.Println(t.UnixMilli()) // 毫秒级
	fmt.Println(t.UnixMicro()) // 微秒级
	fmt.Println(t.UnixNano())  // 纳秒级

	// 时间戳转时间

	t2 := time.Unix(1737200000, 0) // 秒+纳秒
	fmt.Println("Unix: ", t2)

	ms := int64(1737200000000) // 毫秒级时间戳
	t3 := time.UnixMilli(ms)
	fmt.Println("UnixMilli: ", t3)

	us := int64(1737200000000000) // 微秒级时间戳
	t4 := time.UnixMicro(us)
	fmt.Println("UnixMicro:", t4)

}
