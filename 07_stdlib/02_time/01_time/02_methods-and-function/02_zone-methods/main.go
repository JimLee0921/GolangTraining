package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// 获取 t 所属时区
	fmt.Println(t.Location())
	// 获取时区名称和相对 UTC 的秒级偏移量
	fmt.Println(t.Zone())

	// 修改时区
	t = t.UTC()
	fmt.Println(t.Location())
	fmt.Println(t.Zone())

	t = t.UTC().Local()
	fmt.Println(t.Location())
	fmt.Println(t.Zone())

	loc, _ := time.LoadLocation("Africa/Cairo")
	t = t.In(loc)
	fmt.Println(t.Location())
	fmt.Println(t.Zone())
}
