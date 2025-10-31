package main

import (
	"fmt"
	"time"
)

func main() {
	// 可以通过时区等转换来得到新的 Time 对象

	t := time.Now()
	fmt.Println(t)
	//转为 UTC 时区
	fmt.Println(t.UTC())
	// 转回本地时区
	fmt.Println(t.UTC().Local())
	// 转为指定时区
	loc, _ := time.LoadLocation("Africa/Cairo")
	fmt.Println(t.In(loc))
}

/*
2025-10-31 10:44:44.5381968 +0800 +08 m=+0.000000001
2025-10-31 02:44:44.5381968 +0000 UTC
2025-10-31 10:44:44.5381968 +0800 +08
2025-10-31 04:44:44.5381968 +0200 EET
*/
