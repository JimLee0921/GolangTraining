package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建两个 Time 对象用于比较判断
	t1 := time.Now()
	t2 := time.Unix(1761880539, 0)
	fmt.Println(t1.Before(t2)) // false
	fmt.Println(t1.After(t2))  // true
	fmt.Println(t1.Equal(t2))  //false

	t3 := time.Unix(0, 0)
	fmt.Println(t3)          // 1970-01-01 08:00:00 +0800 +08
	fmt.Println(t3.IsZero()) // false
}
