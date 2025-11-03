package main

import (
	"fmt"
	"time"
)

func LoadLocationDemo() {
	// 将当前时间转换为不同地区的时间
	now := time.Now() // 带本地时区
	fmt.Println("Local Time：", now)

	shanghai, _ := time.LoadLocation("Asia/Shanghai")
	newYork, _ := time.LoadLocation("America/New_York")
	london, _ := time.LoadLocation("Europe/London")
	// .In(location) 用于 转换同一时刻 在不同地区的本地表示
	fmt.Println("ShangHai：", now.In(shanghai))
	fmt.Println("NewYork：", now.In(newYork))
	fmt.Println("London：", now.In(london))
}

func CreateTimeByCustomLocDemo() {
	// 使用特定时期创建时间
	loc, _ := time.LoadLocation("America/New_York")
	t := time.Date(2025, 1, 1, 10, 0, 0, 0, loc)
	fmt.Println(t) // 带上 Location 才能明确这是什么地方的 10 点
}

func ParseVsParseInLocationDemo() {
	/*
		Parse 默认 UTC 解析（不推荐）
		ParseInLocation 使用指定的地点解析（推荐）
	*/
	// Parse
	t, _ := time.Parse("2006-01-02 15:04:05", "2025-01-01 10:00:00")
	fmt.Println(t) // 会变成 10 点的 UTC
	// ParseInLocation
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ = time.ParseInLocation("2006-01-02 15:04:05",
		"2025-01-01 10:00:00", loc)
	fmt.Println(t) // 解释为上海时间
}

func TimeToTimestampDemo() {
	// 时间点（timestamp）不变，Location 只影响显示解释
	loc, _ := time.LoadLocation("Asia/Shanghai")

	t := time.Date(2025, 1, 1, 10, 0, 0, 0, loc)
	ts := t.Unix()

	fmt.Println("Unix timestamp：", ts)
	fmt.Println("recover china:", time.Unix(ts, 0).In(loc))
	fmt.Println("recover utc:", time.Unix(ts, 0).UTC())

}

func main() {
	LoadLocationDemo()
	CreateTimeByCustomLocDemo()
	ParseVsParseInLocationDemo()
	TimeToTimestampDemo()
}
